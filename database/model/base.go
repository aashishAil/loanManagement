package model

import (
	"time"

	"github.com/google/uuid"
)

type BaseWithUpdatedAt struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;default:(now() at time zone 'utc')"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;default:(now() at time zone 'utc')"`
}

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;default:(now() at time zone 'utc')"`
}
