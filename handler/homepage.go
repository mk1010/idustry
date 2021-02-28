package handler

import (
	"context"

	"github.com/gin-gonic/gin"
)

type homePage struct{}

func (h *homePage) Register(e *gin.Engine) {
	group := e.Group("/")
	group.Use()
	group.GET("/") // todo handler
}

func homePageWorker(ctx *context.Context) {
}
