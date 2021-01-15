package v1

import (
	"github.com/deoooo/gin_demo/models"
	"github.com/deoooo/gin_demo/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckPassword(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	code := e.SUCCESS
	data := models.CheckPassword(username, password)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
