// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysParam = "sys_param"

// SysParam mapped from table <sys_param>
type SysParam struct {
	SysParamID   string     `gorm:"column:sys_param_id;primaryKey" json:"sys_param_id"`                       // 主键ID
	ParamName    *string    `gorm:"column:param_name" json:"param_name"`                                      // 参数名称
	ParamValue   *string    `gorm:"column:param_value" json:"param_value"`                                    // 参数值
	Notes        *string    `gorm:"column:notes" json:"notes"`                                                // 备注
	HospitalCode *string    `gorm:"column:hospital_code" json:"hospital_code"`                                // 院区编码
	SysProjectID *int32     `gorm:"column:sys_project_id;default:-1" json:"sys_project_id"`                   // 合作项目
	CreateTime   *time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间（数据库自动写入）
	ModifyTime   *time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP" json:"modify_time"` // 记录修改时间（数据库自动写入）
	Comments     *string    `gorm:"column:comments" json:"comments"`                                          // 备注说明
}

// TableName SysParam's table name
func (*SysParam) TableName() string {
	return TableNameSysParam
}

func (*SysParam) PrimaryKey() []string {
	return []string{"sys_param_id"}
}