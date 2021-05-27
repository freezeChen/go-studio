package proxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	registry2 "go-studio/core/registry"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strings"
)

// Proxy will proxy rpc requests as http POST requests. It is a server.Proxy
type Proxy struct {
	Endpoint string
	first    bool
}

func NewProxy(reg registry2.Registrar) *httputil.ReverseProxy {

	director := func(req *http.Request) {
		reqPath := req.URL.Path

		if reqPath == "" {
			return
		}
		pathArray := strings.Split(reqPath, "/")
		serviceName := pathArray[1]
		if serviceName == "favicon.ico" {
			return
		}
		svcs, err := reg.GetService(serviceName)
		if err != nil {
			fmt.Println("query service instance error", err.Error())
			return
		}
		if len(svcs) == 0 {
			fmt.Println("no such service instance", serviceName)
			return
		}
		destPath := strings.Join(pathArray[2:], "/")

		tgt := svcs[rand.Int()%len(svcs)]
		//设置代理服务地址信息
		req.URL.Scheme = "http"
		req.URL.Host = tgt.Endpoints["HTTP"]
		req.URL.Path = "/" + destPath
	}
	gin.New().GET("", func(context *gin.Context) {
		context.RemoteIP()
	})

	return &httputil.ReverseProxy{
		Director: director,
		ErrorHandler: func(writer http.ResponseWriter, request *http.Request, err error) {

			fmt.Println(err)

		},
	}
}
