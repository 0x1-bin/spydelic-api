package handlers

import "github.com/gin-gonic/gin"

func Root(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "hello world"})
}
