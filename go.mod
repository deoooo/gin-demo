module github.com/deoooo/gin_demo

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace github.com/deoooo/gin_demo/conf => ./conf

replace github.com/deoooo/gin_demo/models => ./models

replace github.com/deoooo/gin_demo/routers => ./routers

replace github.com/deoooo/gin_demo/pkg/setting => ./pkg/setting
