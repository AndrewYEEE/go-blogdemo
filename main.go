package main

import (
	"fmt"
	"log"
	"net/http"

	"go-gin-demo/models"
	"go-gin-demo/pkg/setting"
	"go-gin-demo/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("init....")
	setting.Setup()
	models.Setup()
}

func main() {
	gin.SetMode(setting.ENV.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ENV.ReadTimeout
	writeTimeout := setting.ENV.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ENV.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
