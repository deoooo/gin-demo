package api

import (
	"fmt"
	"github.com/boombuler/barcode/qr"
	"github.com/deoooo/gin_demo/pkg/e"
	"github.com/deoooo/gin_demo/pkg/qrcode"
	"github.com/deoooo/gin_demo/service/poster_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	URL = "https://blog.deoooo.me"
)

func GeneratePoster(c *gin.Context) {
	qrCode := qrcode.NewQrCode(URL, 300, 300, qr.M, qr.Auto)
	posterName := poster_service.GetPosterFlag() + "-" + qrcode.GetQrCodeFileName(qrCode.URL) + qrCode.GetQrCodeExt()
	poster := poster_service.NewPoster(posterName, qrCode)
	posterBg := poster_service.NewPosterBg("bg.jpg", poster,
		&poster_service.Rect{X0: 0, Y0: 0, X1: 550, Y1: 700},
		&poster_service.Pt{X: 125, Y: 298},
	)

	_, filePath, err := posterBg.Generate()
	code := e.SUCCESS
	data := make(map[string]interface{})
	if err != nil {
		fmt.Printf("GeneratePoster error:%v", err)
		code = e.ERROR
	} else {
		data["path"] = filePath
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
