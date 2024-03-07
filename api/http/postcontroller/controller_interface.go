package postcontroller

import "github.com/gin-gonic/gin"

type PostController interface {
	AddPostC(c *gin.Context)
}
