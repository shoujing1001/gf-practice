// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// File is the golang structure for table cd_file.
type File struct {
	Id         int    `orm:"id,primary"  json:"id"`         //
	Filepath   string `orm:"filepath"    json:"filepath"`   // 图片路径
	Hash       string `orm:"hash"        json:"hash"`       // 文件hash值
	CreateTime int    `orm:"create_time" json:"createTime"` // 创建时间
}