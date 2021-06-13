package admin

import (
	"gf-practice/app/admin/controller/sys"
	jwtauth "gf-practice/common/service/middleware/jwtAuth"
	"gf-practice/common/service/middleware/router"
)

func init() {
	// 无需验证登录的接口
	g1 := router.New("admin", "sys")
	g1.POST("/login", "", sys.AdminUserApi.Login)

	g2 := router.New("admin", "sys", jwtauth.Auth)
	g2.POST("/getUserInfo", "", sys.AdminUserApi.GetUserInfo)
}
