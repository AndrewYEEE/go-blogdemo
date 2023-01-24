package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()
	//以下是優雅的結束Server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
