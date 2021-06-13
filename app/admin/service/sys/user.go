package sys

import (
	"fmt"
	"gf-practice/app/admin/dao"
	"gf-practice/app/admin/model"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

var AdminuserService = new(adminUserService)

type adminUserService struct{}

func (s *adminUserService) LoginService(user *model.AdminUserApiLoginReq) (gdb.Record, error) {
	where := g.Map{
		"a.user": user.User,
		"a.pwd":  gmd5.MustEncryptString(user.Pwd + g.Cfg().GetString("server.PassSecret")),
	}
	// 使用dao生成的常用函数，返回是model定义的结构体，一些字段会丢失
	// userInfo, err := dao.AdminUser.As("a").InnerJoin("cd_role", "b", "a.role_id in(b.role_id)").Fields("a.user_id,a.name,a.user as username,a.status,a.role_id as user_role_ids,b.role_id,b.name as role_name,b.status as role_status").Where(where).FindOne()

	// 设置prefix可以省略前缀
	userInfo, err := dao.AdminUser.M.As("a").InnerJoin("role", "b", "a.role_id in(b.role_id)").Fields("a.user_id,a.name,a.user as username,a.status,a.role_id as user_role_ids,b.role_id,b.name as role_name,b.status as role_status").Where(where).FindOne()
	// userInfo, err := g.DB("admin").Model("admin_user").As("a").InnerJoin("role", "b", "a.role_id in(b.role_id)").Fields("a.user_id,a.name,a.user as username,a.status,a.role_id as user_role_ids,b.role_id,b.name as role_name,b.status as role_status").Where(where).FindOne()

	// 原生查询
	// userInfo, err := g.DB("admin").GetOne("SELECT `a`.`user_id`,`a`.`name`,a.user as username,`a`.`status`,a.role_id as user_role_ids,`b`.`role_id`,b.name as role_name,b.status as role_status FROM `cd_admin_user` `a` INNER JOIN `cd_role` `b` ON a.role_id in(b.role_id) WHERE  `a`.`user` = 'admin'  AND `a`.`pwd` = '6a5888d05ceb8033ebf0a3dfbf2b416e'")
	// fmt.Println(userInfo)
	if err != nil {
		return nil, err
	}

	if userInfo == nil {
		return nil, gerror.New("请检查用户名或者密码")
	}

	// 如何判断userInfo (gdb.Record)中的 status 属性存在
	// 思考方法一：使用gvar (通用变量)的接口进行类型转换后判断，
	// 参考https://goframe.org/pages/viewpage.action?pageId=1114353
	// fmt.Println(status.Int())
	if userInfo["status"].Int() == 0 || userInfo["role_status"].Int() == 0 {
		fmt.Println("用户被禁用")
		return nil, gerror.New("该账户或角色已被禁用")
	}

	return userInfo, err
}
