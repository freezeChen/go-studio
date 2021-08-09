package interfaces

import (
	 "github.com/freezeChen/go-studio/core/transport/http"
	"github.com/gin-gonic/gin"
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
