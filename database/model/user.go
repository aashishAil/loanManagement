package model

import "loanManagement/constant"

type User struct {
	Base
	Name              string            `json:"name" gorm:"column:name"`
	Email             string            `json:"email" gorm:"column:email"`
	EncryptedPassword string            `json:"encryptedPassword" gorm:"column:encrypted_password"`
	Type              constant.UserType `json:"type" gorm:"column:type"`
}

func (User) TableName() string {
	return "user"
}
