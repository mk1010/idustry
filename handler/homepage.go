package handler

import (
	"context"

	"github.com/gin-gonic/gin"
)

type homePage struct {
}

func (h *homePage) Register(e *gin.Engine) {
	group := e.Group("/")
	group.Use()
	group.GET("/") // todo handler
}

func homePageWorker(ctx *context.Context) {

}

func lengthOfLongestSubstring(s string) int {
    if s=="" {
        return 0
    }
    max,left:=0,-1
	preMap:=make(map[rune]int)
    for i,value:=range s {
       if pos,ok:=preMap[value];ok{
			if pos>left{
				left=pos
			}
	   }
		if max<i-left{
				max=i-left
		}
	   
	   preMap[value]=i
	}
    return max
}
