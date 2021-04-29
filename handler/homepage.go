package handler

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mk1010/idustry/handler/sse"
)

type homePage struct{}

var (
	homePageSseHandler      *sse.SSEHandler
	homePageSseInstanceOnce sync.Once
)

func HomePageSseInstance() *sse.SSEHandler {
	homePageSseInstanceOnce.Do(func() {
		homePageSseHandler = sse.NewSSEHandler()
		homePageSseHandler.HandleEvents()
	})
	return homePageSseHandler
}

func (h *homePage) Register(e *gin.Engine) {
	group := e.Group("/")
	group.Use()
	group.GET("/device/:LogicID", Subscribe) // todo handler
	// group.GET("/", handlers ...gin.HandlerFunc)
}

func homePageWorker() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(ctx.Request)
		if pusher := ctx.Writer.Pusher(); pusher != nil {
			// if err:=pusher.Push(target string, opts *http.PushOptions)
		} else {
			log.Printf("不支持push")
		}

		/*for {
			time.Sleep(1 * time.Second)
			select {
			case <-time.NewTicker(1 * time.Second).C:
				ctx.SSEvent("sse", fmt.Sprintf("%v", time.Now()))
			case <-ctx.Request.Context().Done():
				return
			}
		}*/
		ctx.JSON(200, gin.H{
			"hello": "mk",
		})
	}
}

func Subscribe(c *gin.Context) {
	w := c.Writer
	f, ok := w.(http.Flusher)
	if !ok {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Streaming unsupported"))
		return
	}
	logicID := c.Param("LogicID")
	// Create a new channel, over which we can send this client messages.
	messageChan := make(chan string)
	// Add this client to the map of those that should receive updates
	HomePageSseInstance().AppendClient(sse.SseClient{
		LogicID:  string(logicID),
		UserChan: messageChan,
	})

	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		// Remove this client from the map of attached clients
		HomePageSseInstance().RemoveClient(sse.SseClient{
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
