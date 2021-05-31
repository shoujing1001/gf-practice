// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"gf-practice/app/api/model/internal"
)

// User is the golang structure for table prac_user.
type User internal.User

// Fill with you ideas below.
// 请求参数，用于前后端交互参数格式约定
type UserApiSignUpReq struct {
	UserName  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Nickname  string
	AvatarUrl string `d:"http://qiniuoss.nauzone.cn/common/avatar.png"`
}

type UserApiSignInReq struct {
	UserName string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
}

// 服务输入参数
type UserServiceSignUpReq struct {
	UserName  string
	Password  string
	Nickname  string
	AvatarUrl string
}

// 接口输出参数
type UserApiSignInRes struct {
	UserInfo *User
	Token    string
}
