package interfaces

import (
	"github.com/gin-gonic/gin"
	http2 "github.com/freezeChen/go-studio/core/transport/http"
)

func NewHttpServer() *http2.Server {
	srv := http2.NewServer(":8081")
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.GET("test", func(ctx *gin.Context) {
		ctx.JSON(200, "hello")
	})

	srv.Handle(engine)

	return srv
}
