// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDepartment = "sys_department"

// Department 系统部门表
type Department struct {
	ID        int64          `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true;comment:主键ID|Primary key" json:"id"`                     // 主键ID|Primary key
	ParentID  *int64         `gorm:"column:parent_id;type:bigint unsigned;comment:父部门ID|Parent department ID" json:"parent_id"`                           // 父部门ID|Parent department ID
	Name      string         `gorm:"column:name;type:varchar(50);not null;comment:部门名称|Department name" json:"name"`                                      // 部门名称|Department name
	Code      string         `gorm:"column:code;type:varchar(50);not null;comment:部门编码|Department code" json:"code"`                                      // 部门编码|Department code
	Leader    *string        `gorm:"column:leader;type:varchar(32);comment:部门负责人|Department leader" json:"leader"`                                        // 部门负责人|Department leader
	Phone     *string        `gorm:"column:phone;type:varchar(11);comment:联系电话|Contact number" json:"phone"`                                              // 联系电话|Contact number
	Email     *string        `gorm:"column:email;type:varchar(64);comment:邮箱|Email" json:"email"`                                                         // 邮箱|Email
	Sort      int64          `gorm:"column:sort;type:int unsigned;not null;comment:排序|Sort" json:"sort"`                                                  // 排序|Sort
	Status    *int64         `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:状态 1:启用 2:禁用|Status 1:Enable 2:Disable" json:"status"` // 状态 1:启用 2:禁用|Status 1:Enable 2:Disable
	Remark    *string        `gorm:"column:remark;type:varchar(255);comment:备注|Remark" json:"remark"`                                                     // 备注|Remark
	CreatedAt *time.Time     `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间|Created Time" json:"created_at"`      // 创建时间|Created Time
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间|Updated Time" json:"updated_at"`      // 更新时间|Updated Time
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间|Deleted Time" json:"deleted_at"`                                         // 删除时间|Deleted Time
	Parent    *Department    `gorm:"foreignKey:ParentID;references:ID" json:"parent"`
	Children  []*Department  `gorm:"foreignKey:ParentID;references:ID" json:"children"`
}

// TableName Department's table name
func (*Department) TableName() string {
	return TableNameDepartment
}
