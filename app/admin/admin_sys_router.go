package admin

import (
	"gf-practice/app/admin/controller/sys"
	"gf-practice/common/service/middleware/router"
)

func init() {
	g1 := router.New("admin", "sys")
	g1.POST("/login", "", sys.AdminUserApi.Login)
}
