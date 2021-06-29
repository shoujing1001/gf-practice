package sys

import (
	"gf-practice/app/admin/dao"
	"gf-practice/app/admin/service/sys"
	"gf-practice/common/response"

	"github.com/gogf/gf/net/ghttp"
)

type adminMenuApi struct{}

var AdminMenuApi = new(adminMenuApi)

func (i *adminMenuApi) GetMenuList(r *ghttp.Request) {
	field := "menu_id,pid,title,controller_name,status,icon,sortid,router_path,component_path"
	var tempResult []*dao.RoleMenusTree
	// 结构体数组转换方式
	err := dao.Menu.M.Fields(field).Where("app_id = ?", 1).Order("sortid asc").Scan(&tempResult)

	res := sys.AdminuserService.GetMenusTree(tempResult, 0)

	if err != nil {
		response.JsonExit(r, errorCode, err.Error())
	}

	response.JsonExit(r, successCode, "操作成功", res)
}
