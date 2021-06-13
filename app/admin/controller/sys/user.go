package sys

import (
	"fmt"
	"gf-practice/app/admin/model"
	"gf-practice/app/admin/service/sys"
	"gf-practice/common/response"
	"gf-practice/common/token"

	"github.com/gogf/gf/net/ghttp"
)

var AdminUserApi = new(adminUserApi)

type adminUserApi struct{}

func (i *adminUserApi) Login(r *ghttp.Request) {
	var loginReq *model.AdminUserApiLoginReq
	if err := r.Parse(&loginReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	fmt.Println(loginReq)
	userInfo, err := sys.AdminuserService.LoginService(loginReq)
	fmt.Println(userInfo)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	jwtToken, err := token.New(userInfo["user_id"].String()).CreateToken()

	response.JsonExit(r, 0, "操作成功", jwtToken)
}

func (i *adminUserApi) GetUserInfo(r *ghttp.Request) {
	adminId := r.GetParam("uid")
	fmt.Println(adminId)
	response.JsonExit(r, 0, "操作成功")
}
