package call_service

import "github.com/deoooo/gin_demo/models"

func Call(mobileNumber string) {
	models.CreateCallRecord(mobileNumber)
}
