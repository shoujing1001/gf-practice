// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// AdminUser is the golang structure for table cd_admin_user.
type AdminUser struct {
	UserId     int    `orm:"user_id,primary" json:"userId"`     // 编号
	Name       string `orm:"name"            json:"name"`       // 真实姓名
	User       string `orm:"user"            json:"user"`       // 用户名
	Pwd        string `orm:"pwd"             json:"pwd"`        // 密码
	RoleId     int    `orm:"role_id"         json:"roleId"`     // 所属分组
	Note       string `orm:"note"            json:"note"`       // 备注
	Status     int    `orm:"status"          json:"status"`     // 状态
	CreateTime int    `orm:"create_time"     json:"createTime"` // 创建时间
}
