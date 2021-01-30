package api

import (
	"fmt"
	"github.com/deoooo/gin_demo/pkg/e"
	"github.com/deoooo/gin_demo/service/excel_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Export(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	path, err := excel_service.Export()
	if err != nil {
		fmt.Printf("export err:%v \n", err)
		code = e.ERROR
	} else {
		data["path"] = path
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
