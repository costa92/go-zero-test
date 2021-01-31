package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"image/internal/config"
	"image/internal/handler"
	"image/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/open-api.yaml", "the config file")

// 返回的结构体，json格式的body
type Message struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

// 定义错误处理函数
func errorHandler(err error) (int, interface{}) {
	return http.StatusConflict, Message{
		Code: -1,
		Desc: err.Error(),
	}
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	httpx.SetErrorHandler(errorHandler)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
