package main

import (
	"fmt"
	_ "github.com/deoooo/gin_demo/docs"
	"github.com/deoooo/gin_demo/models"
	"github.com/deoooo/gin_demo/pkg/setting"
	"github.com/deoooo/gin_demo/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	setting.Setup()
	models.Setup()

	gin.SetMode(setting.ServerSetting.RunMode)
	r := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	s := &http.Server{
		Addr:           endPoint,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	s.ListenAndServe()
}
