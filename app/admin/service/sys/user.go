package sys

import (
	"fmt"
	"gf-practice/app/admin/dao"
	"gf-practice/app/admin/model"
	"gf-practice/common/utils"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

var AdminuserService = new(adminUserService)

type adminUserService struct{}

// 树级菜单
type menusTreee struct {
	Children []menusTreee
	model.Menu
	Name string
}

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

func (s *adminUserService) GetUserInfoService(uid int) (g.Map, error) {
	if uid == 0 {
		return nil, gerror.New("用户ID不存在")
	}
	userInfo, err := dao.AdminUser.M.Fields("role_id,user").Where("user_id = ?", uid).FindOne()

	// 将结果转化为普通的map，以便后续追加属性
	userInfoMap := userInfo.Map()

	fmt.Println(userInfoMap)

	if err != nil {
		return nil, err
	}

	if userInfo == nil {
		return nil, gerror.New("用户不存在")
	}

	roleInfo, err := dao.Role.M.Where("role_id = ?", userInfoMap["role_id"]).FindOne()

	userInfoMap["site_title"] = g.Cfg("gfsadmin").GetString("gfayabase.siteTitle")
	userInfoMap["headimg"] = g.Cfg("gfsadmin").GetString("gfayabase.logo")
	userInfoMap["show_notice"] = g.Cfg("gfsadmin").GetBool("gfayabase.showNotice")

	menusTree, err := getGfsMenus(roleInfo)
	var component []*dao.RoleComponents
	// 此条件有问题，不应当返回只有路径而无组件的菜单给其path
	dao.Menu.M.Where("app_id", 1).Where("component_path <>", "").Scan(&component)

	resData := g.Map{
		"data":      userInfoMap,
		"menu":      menusTree,
		"component": component,
	}

	return resData, nil
}

// 返回当前应用登录用户的菜单列表
func getGfsMenus(roleInfo gdb.Record) ([]*dao.RoleMenusTree, error) {
	roleAccess := gstr.Split(roleInfo["access"].String(), ",")
	fmt.Println(roleAccess)
	roleMenus, err := getGfsRoleMenus(roleAccess)
	if err != nil {
		fmt.Println("获取角色菜单数据错误", err)
		return nil, err
	}
	menusTree := getMenusTree(roleMenus, 0)
	return menusTree, nil
}

// 获取当前角色有权限的菜单
func getGfsRoleMenus(roleAccess []string) ([]*dao.RoleMenusTree, error) {
	field := "menu_id,pid,title,controller_name,status,icon,sortid,router_path,component_path"
	var tempResult []*dao.RoleMenusTree
	// 结构体数组转换方式
	err := dao.Menu.M.Fields(field).Where("status = ? AND app_id = ?", 1, 1).Order("sortid asc").Scan(&tempResult)
	// err := dao.Menu.M.Fields(field).Where("status = ? AND app_id = ?", 1, 1).Order("sortid asc").Structs(&tempResult)

	if err != nil {
		return nil, err
	}

	if err != nil {
		fmt.Println("获取菜单数据错误", err)
		return nil, err
	}
	var result []*dao.RoleMenusTree
	for _, v := range tempResult {
		if utils.InStrArray(roleAccess, gconv.String(v.MenuId)) {
			result = append(result, v)
		}
	}
	fmt.Println(result)
	return result, nil
}

// 递归实现(返回树状菜单数据)
func getMenusTree(allMenus []*dao.RoleMenusTree, pid int) []*dao.RoleMenusTree {
	var menusTree []*dao.RoleMenusTree
	for _, v := range allMenus {
		if pid == v.Pid {
			fmt.Println("当前遍历之菜单：", v.MenuId, pid)
			menusItem := v
			// fmt.Println("加入该角色权限菜单：", menusItem)
			children := getMenusTree(allMenus, v.MenuId)
			if children != nil {
				menusItem.Children = children
			} else {
				// 初始化之后，让此字段返回为[]而非null
				menusItem.Children = []*dao.RoleMenusTree{}
			}
			menusTree = append(menusTree, menusItem)
		}
	}
	return menusTree
}
