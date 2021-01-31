package qrcode_service

import (
	"github.com/boombuler/barcode/qr"
	"github.com/deoooo/gin_demo/pkg/qrcode"
)

func GenerateQrcode(url string) (string, error) {
	qrc := qrcode.NewQrCode(url, 300, 300, qr.M, qr.Auto)
	path := qrcode.GetQrCodeFullPath()
	_, src, err := qrc.Encode(path)
	return src, err
}
