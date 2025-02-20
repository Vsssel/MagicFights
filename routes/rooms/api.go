package rooms

import (
	"crypto/md5"

	"github.com/gin-gonic/gin"
)

func createRoom(ctx *gin.Context) {
	hash := md5.New()

	hash.Write([]byte(ctx.A))
}