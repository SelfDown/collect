// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAuthUser = "auth_user"

// AuthUser mapped from table <auth_user>
type AuthUser struct {
	ID          int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Password    string     `gorm:"column:password;not null" json:"password"`
	LastLogin   *time.Time `gorm:"column:last_login" json:"last_login"`
	IsSuperuser bool       `gorm:"column:is_superuser;not null" json:"is_superuser"`
	Username    string     `gorm:"column:username;not null" json:"username"`
	FirstName   string     `gorm:"column:first_name;not null" json:"first_name"`
	LastName    string     `gorm:"column:last_name;not null" json:"last_name"`
	Email       string     `gorm:"column:email;not null" json:"email"`
	IsStaff     bool       `gorm:"column:is_staff;not null" json:"is_staff"`
	IsActive    bool       `gorm:"column:is_active;not null" json:"is_active"`
	DateJoined  time.Time  `gorm:"column:date_joined;not null" json:"date_joined"`
}

// TableName AuthUser's table name
func (*AuthUser) TableName() string {
	return TableNameAuthUser
}

func (*AuthUser) PrimaryKey() []string {
	return []string{"id"}
}