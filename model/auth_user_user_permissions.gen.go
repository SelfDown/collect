// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAuthUserUserPermissions = "auth_user_user_permissions"

// AuthUserUserPermissions mapped from table <auth_user_user_permissions>
type AuthUserUserPermissions struct {
	ID           int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID       int32 `gorm:"column:user_id;not null" json:"user_id"`
	PermissionID int32 `gorm:"column:permission_id;not null" json:"permission_id"`
}

// TableName AuthUserUserPermissions's table name
func (*AuthUserUserPermissions) TableName() string {
	return TableNameAuthUserUserPermissions
}

func (*AuthUserUserPermissions) PrimaryKey() []string {
	return []string{"id"}
}