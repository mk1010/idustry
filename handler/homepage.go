package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type homePage struct{}

func (h *homePage) Register(e *gin.Engine) {
	group := e.Group("/")
	group.Use()
	group.GET("/", Subscribe) // todo handler
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

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		time.Sleep(time.Second)
		_, err := fmt.Fprintf(w, "data: Message: %v\n\n", time.Now())
		if err != nil {
			log.Println(err)
			break
		}
		// Flush the response. This is only possible if the repsonse
		// supports streaming.
		f.Flush()
	}
	log.Println("mk", c.Request.Proto)
	c.AbortWithStatus(http.StatusOK)
}

var messageChan chan string

func handleSSE() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Get handshake from client")

		// prepare the header
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// instantiate the channel
		messageChan = make(chan string)

		// close the channel after exit the function
		defer func() {
			close(messageChan)
			messageChan = nil
			log.Printf("client connection is closed")
		}()

		// prepare the flusher
		flusher, _ := w.(http.Flusher)

		// trap the request under loop forever
		for {
			select {

			// message will received here and printed
			case message := <-messageChan:
				fmt.Fprintf(w, "%s\n", message)
				flusher.Flush()

			// connection is closed then defer will be executed
			case <-r.Context().Done():
				return

			}
		}
	}
}

func sendMessage(message string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if messageChan != nil {
			log.Printf("print message to client")

			// send the message through the available channel
			messageChan <- message
		}
	}
}
