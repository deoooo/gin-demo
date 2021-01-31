package api

import (
	"github.com/deoooo/gin_demo/pkg/e"
	"github.com/deoooo/gin_demo/service/qrcode_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GenerateQrcodeDto struct {
	Url string `json:"url"`
}

func GenerateQrcode(c *gin.Context) {
	var generateQrcodeDto GenerateQrcodeDto
	err := c.BindJSON(&generateQrcodeDto)
	data := make(map[string]interface{})
	code := e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		path, err := qrcode_service.GenerateQrcode(generateQrcodeDto.Url)
		if err != nil {
			code = e.ERROR
		} else {
			data["path"] = path
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
