package v1

import (
	"fmt"
	"github.com/deoooo/gin_demo/models"
	"github.com/deoooo/gin_demo/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CallDto struct {
	MobileNumber string `json:"mobileNumber"`
}

func Call(c *gin.Context) {
	var callDto CallDto
	err := c.BindJSON(&callDto)
	code := e.SUCCESS
	if err != nil {
		fmt.Printf("bing json err %v \n", err)
		code = e.INVALID_PARAMS
	}
	record, err := models.CreateCallRecord(callDto.MobileNumber)
	if err != nil {
		fmt.Printf("create record fail err:%v \n", err)
		code = e.ERROR
	} else {
		fmt.Printf("create record success id:%d \n", record.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
