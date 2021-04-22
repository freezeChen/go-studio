package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-studio/transport/http"
)

func NewHttpServer() *http.Server {
	srv := http.NewServer(":8081")
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.GET("test", func(ctx *gin.Context) {
		ctx.JSON(200, "hello")
	})

	srv.Handle(engine)

	return srv
}
