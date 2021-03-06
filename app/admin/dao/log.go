// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gf-practice/app/admin/dao/internal"
)

// logDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type logDao struct {
	internal.LogDao
}

var (
	// Log is globally public accessible object for table cd_log operations.
	Log = logDao{
		internal.Log,
	}
)

// Fill with you ideas below.
