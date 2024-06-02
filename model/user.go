package model

import (
	"github.com/google/uuid"
	"loanManagement/constant"
)

type LoggedInUser struct {
	ID   uuid.UUID
	Type constant.UserType
}
