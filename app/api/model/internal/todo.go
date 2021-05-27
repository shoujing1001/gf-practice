// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// Todo is the golang structure for table prac_todo.
type Todo struct {
	TodoId     int    `orm:"todoId,primary" json:"todoId"`     // 待做主键
	UserId     int    `orm:"userId"         json:"userId"`     // 关联用户主键
	Title      string `orm:"title"          json:"title"`      // 事项名称
	Content    string `orm:"content"        json:"content"`    // todo事项内容
	AddTime    int    `orm:"addTime"        json:"addTime"`    // 添加时间
	Status     string `orm:"status"         json:"status"`     // 事项状态
	UpdateTime int    `orm:"updateTime"     json:"updateTime"` // 更新时间
}