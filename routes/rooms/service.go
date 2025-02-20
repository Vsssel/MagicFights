package rooms

import (
	"crypto/md5"
)

func CreateRoom() {
	hash := md5.New()

	hash.Write([]byte(ctx.A))
}