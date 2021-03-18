package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-studio/transport/http"
)

func NewHttpServer() *http.Server {
	srv := http.NewServer("")

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	srv.Handle("/", engine)

	return srv
}
