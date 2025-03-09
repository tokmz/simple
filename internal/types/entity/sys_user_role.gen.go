// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameUserRole = "sys_user_role"

// UserRole 用户-角色关系表
type UserRole struct {
	ID        int64      `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:主键ID|Primary key" json:"id"`                // 主键ID|Primary key
	UserID    int64      `gorm:"column:user_id;type:bigint unsigned;not null;comment:用户ID|User ID" json:"user_id"`                               // 用户ID|User ID
	RoleID    int64      `gorm:"column:role_id;type:bigint unsigned;not null;comment:角色ID|Role ID" json:"role_id"`                               // 角色ID|Role ID
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间|Created Time" json:"created_at"` // 创建时间|Created Time
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间|Updated Time" json:"updated_at"` // 更新时间|Updated Time
}

// TableName UserRole's table name
func (*UserRole) TableName() string {
	return TableNameUserRole
}
