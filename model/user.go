package model

import (
	"github.com/google/uuid"
	"loanManagement/constant"
)

type LoggedInUser struct {
	ID   uuid.UUID         `json:"id"`
	Type constant.UserType `json:"type"`
}
