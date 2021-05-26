package controller

import (
	"fmt"
	"gf-practice/app/api/model"
	"gf-practice/common/response"

	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
var User = new(userApi)

type userApi struct{}

// @summary 用户注册接口
// @tags    用户服务
// @produce json
// @param   entity  body model.UserApiSignUpReq true "注册请求"
// @router  /user/signup [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignUpReq
		serviceReq *model.UserServiceSignUpReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	// if err := gconv.Struct(apiReq, &serviceReq); err != nil {
	// 	response.JsonExit(r, 1, err.Error())
	// }
	fmt.Println(apiReq)
	fmt.Println(serviceReq)
	// if err := service.User.SignUp(serviceReq); err != nil {
	// 	response.JsonExit(r, 1, err.Error())
	// } else {
	response.JsonExit(r, 0, "ok")
	// }
}
