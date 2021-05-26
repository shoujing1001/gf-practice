package api

import (
	"fmt"
	"gf-practice/app/api/controller"
	"gf-practice/app/api/controller/hello"
	"gf-practice/common/service/middleware/router"
)

func init() {
	g1 := router.New("api", "/hello")
	fmt.Println("api router 初始化")
	g1.GET("/world", "api:hello:world", hello.World)
	g2 := router.New("api", "/user")
	g2.POST("/signup", "", controller.User.SignUp)
}
