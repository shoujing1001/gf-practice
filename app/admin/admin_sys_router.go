package admin

import (
	"gf-practice/app/admin/controller/sys"
	"gf-practice/common/service/middleware/cors"
	jwtauth "gf-practice/common/service/middleware/jwtAuth"
	"gf-practice/common/service/middleware/router"
)

func init() {
	// 无需验证登录的接口
	g1 := router.New("admin", "sys/user", cors.MiddlewareCORS)
	g1.POST("/login", "", sys.AdminUserApi.Login)

	g2 := router.New("admin", "sys/", cors.MiddlewareCORS, jwtauth.Auth)
	g2.POST("user/getUserInfo", "", sys.AdminUserApi.GetUserInfo)
	g2.POST("menu/getMenuList", "", sys.AdminMenuApi.GetMenuList)
}
