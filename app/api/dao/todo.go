// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gf-practice/app/api/dao/internal"
)

// todoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type todoDao struct {
	internal.TodoDao
}

var (
	// Todo is globally public accessible object for table prac_todo operations.
	Todo = todoDao{
		internal.Todo,
	}
)

// Fill with you ideas below.