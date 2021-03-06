// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gf-practice/app/admin/dao/internal"
)

// menuDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type menuDao struct {
	internal.MenuDao
}

var (
	// Menu is globally public accessible object for table cd_menu operations.
	Menu = menuDao{
		internal.Menu,
	}
)

// Fill with you ideas below.

// type RoleMenuItem struct {
// 	MenuId        int    `json:"menu_id"`                 //
// 	Pid           int    `json:"pid"`                     // 父级id
// 	Name          string `c:"ControllerName" json:"name"` // 模块名称
// 	Title         string `json:"title"`                   // 模块标题
// 	Status        int    `json:"status"`                  // 0隐藏 1显示
// 	Sortid        int    `json:"sort_id"`                 // 排序号
// 	ComponentPath string `json:"component_path"`          // 组件路径
// 	Icon          string `json:"icon"`                    // icon字体图标
// }

type RoleMenusTree struct {
	MenuId   int              `json:"menu_id"`                 //
	Pid      int              `json:"pid"`                     // 父级id
	Name     string           `c:"ControllerName" json:"name"` // 模块名称
	Title    string           `json:"title"`                   // 模块标题
	Status   int              `json:"status"`                  // 0隐藏 1显示
	Sortid   int              `json:"sort_id"`                 // 排序号
	Icon     string           `json:"icon"`                    // icon字体图标
	Children []*RoleMenusTree `json:"children" d:"[]"`
	Path     string           `c:"router_path" json:"path"`
	CPath    string           `c:"component_path" json:"cPath"` // 组件路径
}

type RoleComponents struct {
	ComponentPath string `json:"component_path"`          // 组件文件位置
	Name          string `c:"ControllerName" json:"name"` // 模块名称
	Title         string `json:"title"`                   // 模块标题
	// Path          string `c:"ComponentPath" json:"path"`  // 组件路由路径
}
