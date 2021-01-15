package jwt

import (
	"github.com/deoooo/gin_demo/pkg/e"
	"github.com/deoooo/gin_demo/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			c.Set("userId", claims.UserId)
			if err != nil {
				code = e.AuthFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.NeedAuth
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
