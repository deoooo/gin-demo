package api

import (
	"github.com/deoooo/gin_demo/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 健康检查
// @Description 健康检查,返回服务当前的状态
// @Tags 基础接口
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  "ok",
	})
}
