package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	Id uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`

	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique"`

	CreatedAt time.Time      `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deletedAt;index"`
}
