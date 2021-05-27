package service

import (
	"errors"
	"fmt"
	"gf-practice/app/api/dao"
	"gf-practice/app/api/model"
)

// 用户service对象
var User = userService{}

type userService struct{}

// 用户注册
func (s *userService) SignUp(r *model.UserServiceSignUpReq) error {
	// 昵称为非必需参数，默认使用账号名称
	if r.Nickname == "" {
		r.Nickname = r.UserName
	}

	// 账号唯一性数据检查
	if !s.CheckPassport(r.UserName) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", r.UserName))
	}
	fmt.Println(r)
	if _, err := dao.User.Save(r); err != nil {
		return err
	}
	return nil
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckPassport(passport string) bool {
	if i, err := dao.User.FindCount("userName", passport); err != nil {
		return false
	} else {
		return i == 0
	}
}
