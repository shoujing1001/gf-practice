package sys

import (
	"fmt"
	"gf-practice/app/admin/model"
	"gf-practice/common/response"

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
}
