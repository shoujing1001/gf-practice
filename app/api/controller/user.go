package controller

import (
	"fmt"
	"gf-practice/app/api/model"
	"gf-practice/app/api/service"
	"gf-practice/common/response"
	"gf-practice/common/token"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
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
func (i *userApi) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignUpReq
		serviceReq *model.UserServiceSignUpReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	fmt.Println(apiReq)
	fmt.Println(serviceReq)
	if err := service.User.SignUp(serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// @summary 用户登录接口
// @tags    用户服务
// @produce json
// @param   entity  body model.UserApiSignInReq true "注册请求"
// @router  /user/signup [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (i *userApi) SignIn(r *ghttp.Request) {
	var apiReq *model.UserApiSignInReq
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	fmt.Println(apiReq)
	userInfo, err := service.User.SignIn(apiReq)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if userInfo == nil {
		response.JsonExit(r, 1, "账密错误")
	}
	fmt.Println(userInfo)

	jwtToken, err := token.New(gconv.String(userInfo.UserId)).CreateToken()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	var signRes = new(model.UserApiSignInRes)
	signRes.UserInfo = userInfo
	signRes.Token = jwtToken
	response.JsonExit(r, 0, "操作成功", signRes)
}
