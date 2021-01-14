package main

import (
	"fmt"
	"github.com/deoooo/gin_demo/pkg/setting"
	"github.com/deoooo/gin_demo/routers"
	"net/http"
)

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf("127.0.0.1:%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	e := s.ListenAndServe()
	if e != nil {
		fmt.Printf("Listen Error %v", e)
	} else {
		fmt.Printf("Listen on port:%d", setting.HTTPPort)
	}
}
