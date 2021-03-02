package sse

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// source from https://github.com/lmas/gin-sse/blob/master/sse_handler.go

type SSEHandler struct {
	// Create a map of clients, the keys of the map are the channels over
	// which we can push messages to attached clients. (The values are just
	// booleans and are meaningless.)
	// SSE成功链接队列 key logicID
	clients map[string]chan string

	// Channel into which new clients can be pushed
	// http握手成功等待加入SSE队列
	newClients chan SseClient

	// Channel into which disconnected clients should be pushed
	// http链接关闭 等待从SSE成功队列移除
	defunctClients chan SseClient

	// Channel into which messages are pushed to be broadcast out
	// 广播消息队列
	messages chan SseMessage
}

type SseClient struct {
	LogicID  string
	UserChan chan string
}

// LogicIDs为空，为广播
type SseMessage struct {
	LogicIDs []string
	Msg      string
}

// Make a new SSEHandler instance.
func NewSSEHandler() *SSEHandler {
	b := &SSEHandler{
		clients:        make(map[string]chan string),
		newClients:     make(chan SseClient),
		defunctClients: make(chan SseClient),
		messages:       make(chan SseMessage, 100), // buffer 100 msgs and don't block sends
	}
	return b
}

// Start handling new and disconnected clients, as well as sending messages to
// all connected clients.
// 并发安全，性能瓶颈
// 修改该函数时，特别小心，确定知道自己在做什么同时熟悉channel
func (b *SSEHandler) HandleEvents() {
	go func() {
		for {
			select {
			case s := <-b.newClients:
				b.clients[s.LogicID] = s.UserChan
			case s := <-b.defunctClients:
				userChan, ok := b.clients[s.LogicID]
				delete(b.clients, s.LogicID)
				if ok {
					close(userChan)
				}
			case sseMsg := <-b.messages:
				if len(sseMsg.LogicIDs) == 0 {
					for _, s := range b.clients {
						s <- sseMsg.Msg
					}
				} else {
					for _, id := range sseMsg.LogicIDs {
						if userChan, ok := b.clients[id]; ok {
							userChan <- sseMsg.Msg
						}
					}
				}
			}
		}
	}()
}

// 广播 Send out a simple string to all clients.
func (b *SSEHandler) SendString(msg SseMessage) {
	b.messages <- msg
}

func (b *SSEHandler) AppendClient(cli SseClient) {
	b.newClients <- cli
}

func (b *SSEHandler) RemoveClient(cli SseClient) {
	b.defunctClients <- cli
}

// Subscribe a new client and start sending out messages to it.
// 示例 请不要直接使用
func (b *SSEHandler) Subscribe(c *gin.Context) {
	w := c.Writer
	f, ok := w.(http.Flusher)
	if !ok {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Streaming unsupported"))
		return
	}
	logicID := ""
	id, ok := c.Get("logic_id")
	switch id.(type) {
	case string:
		logicID = id.(string)
	default:
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("illegal logic_id"))
		return
	}

	if !ok {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("need logic_id"))
		return
	}
	// Create a new channel, over which we can send this client messages.
	messageChan := make(chan string)
	// Add this client to the map of those that should receive updates
	b.AppendClient(SseClient{
		LogicID:  string(logicID),
		UserChan: messageChan,
	})

	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		// Remove this client from the map of attached clients
		b.RemoveClient(SseClient{
			LogicID:  string(logicID),
			UserChan: messageChan,
		})
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		msg, open := <-messageChan
		if !open {
			// If our messageChan was closed, this means that
			// the client has disconnected.
			break
		}

		fmt.Fprintf(w, "data: Message: %s\n\n", msg)
		// Flush the response. This is only possible if the repsonse
		// supports streaming.
		f.Flush()
	}

	c.AbortWithStatus(http.StatusOK)
}
