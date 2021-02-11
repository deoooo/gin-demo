package sign

import (
	"encoding/json"
	"github.com/deoooo/gin_demo/pkg/e"
	"github.com/deoooo/gin_demo/pkg/util"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ValidSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := ioutil.ReadAll(c.Request.Body)
		m := make(map[string]interface{})
		err := json.Unmarshal(data, &m)
		ok := util.ValidSign(m)

		if err != nil || !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": e.INVALID_PARAMS,
				"msg":  "invalid sign",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
