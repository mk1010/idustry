package handler

import "github.com/gin-gonic/gin"

//向指定位置（index）插入中间件
func AddMiddleware(e *gin.Engine, index int, fn gin.HandlerFunc) {
	pre := e.Handlers
	cur := make(gin.HandlersChain, len(pre)+1)
	copy(cur, pre[0:index])
	copy(cur[index+1:], pre[index:])
	cur[index] = fn
	e.Handlers = cur
}

//set
func SetMiddleware(e *gin.Engine, index int, fn gin.HandlerFunc) {
	e.Handlers[index] = fn
}
