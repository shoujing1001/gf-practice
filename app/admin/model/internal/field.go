// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// Field is the golang structure for table cd_field.
type Field struct {
	Id               int    `orm:"id,primary"         json:"id"`               //
	MenuId           int    `orm:"menu_id"            json:"menuId"`           // 模块ID
	Title            string `orm:"title"              json:"title"`            // 字段名称
	Field            string `orm:"field"              json:"field"`            //
	Type             int    `orm:"type"               json:"type"`             // 表单类型1输入框 2下拉框 3单选框 4多选框 5上传图片 6编辑器 7时间
	ListShow         int    `orm:"list_show"          json:"listShow"`         // 列表显示
	SearchType       int    `orm:"search_type"        json:"searchType"`       // 1精确匹配 2模糊搜索
	PostStatus       int    `orm:"post_status"        json:"postStatus"`       // 是否前台录入
	CreateTableField int    `orm:"create_table_field" json:"createTableField"` //
	Validate         string `orm:"validate"           json:"validate"`         // 验证方式
	Rule             string `orm:"rule"               json:"rule"`             // 验证规则
	Sortid           int    `orm:"sortid"             json:"sortid"`           // 排序号
	Sql              string `orm:"sql"                json:"sql"`              // 字段配置数据源sql
	DefaultValue     string `orm:"default_value"      json:"defaultValue"`     //
	Datatype         string `orm:"datatype"           json:"datatype"`         // 字段数据类型
	Length           string `orm:"length"             json:"length"`           // 字段长度
	Indexdata        string `orm:"indexdata"          json:"indexdata"`        // 索引
	ShowCondition    string `orm:"show_condition"     json:"showCondition"`    //
	ItemConfig       string `orm:"item_config"        json:"itemConfig"`       //
	Width            string `orm:"width"              json:"width"`            // 单元格宽度
	DatetimeConfig   string `orm:"datetime_config"    json:"datetimeConfig"`   // 其他配置
	OtherConfig      string `orm:"other_config"       json:"otherConfig"`      //
	BelongTable      string `orm:"belong_table"       json:"belongTable"`      // 虚拟字段所属表 用户多表关联
}
