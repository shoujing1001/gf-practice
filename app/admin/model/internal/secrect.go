// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// Secrect is the golang structure for table cd_secrect.
type Secrect struct {
	SecrectId int    `orm:"secrect_id,primary" json:"secrectId"` // 编号
	Name      string `orm:"name"               json:"name"`      //
	Data      string `orm:"data"               json:"data"`      //
}