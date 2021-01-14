module github.com/deoooo/gin_demo

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace github.com/deoooo/gin_demo/conf => ./conf

replace github.com/deoooo/gin_demo/models => ./models

replace github.com/deoooo/gin_demo/routers => ./routers

replace github.com/deoooo/gin_demo/pkg/setting => ./pkg/setting
