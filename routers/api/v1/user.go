package v1

import (
	"github.com/deoooo/gin_demo/models"
	"github.com/deoooo/gin_demo/pkg/e"
	"github.com/deoooo/gin_demo/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 登陆
// @Description 通过帐号密码登陆，获取token
// @Tags Users
// @Accept  json
// @Produce  json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {string} json "{"code":200,"data":{"token"：""},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":{},"msg":"invalid params"}"
// @Router /api/v1/login [get]
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	code := e.SUCCESS
	data := make(map[string]interface{})

	ok, u := models.CheckPassword(username, password)
	if ok {
		token, _ := util.GenerateToken(u.ID)
		data["token"] = token
	} else {
		code = e.AuthFail
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func Hello(c *gin.Context) {
	code := e.SUCCESS
	userId, exit := c.Get("userId")
	data := make(map[string]interface{})
	if exit {
		data["userId"] = userId
	} else {
		code = e.AuthFail
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
