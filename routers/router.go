package routers

import (
	"github.com/deoooo/gin_demo/middleware/jwt"
	"github.com/deoooo/gin_demo/pkg/upload"
	"github.com/deoooo/gin_demo/routers/api"
	v1 "github.com/deoooo/gin_demo/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.GET("/swagger/*any",
		gin.BasicAuth(gin.Accounts{"username": "password"}),
		ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", api.Health)

	r.POST("/upload", api.UploadImage)
	r.POST("/excel/export", api.Export)
	r.POST("/qrcode", api.GenerateQrcode)
	r.POST("/poster", api.GeneratePoster)
	apiV1 := r.Group("/api/v1")
	{
		// 登陆接口
		apiV1.GET("/login", v1.Login)
	}

	v1User := apiV1.Group("/user")
	{
		v1User.GET("/hello", jwt.JWT(), v1.Hello)
	}

	return r
}
