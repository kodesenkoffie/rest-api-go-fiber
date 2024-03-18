package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	Id uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`

	Title string `json:"title" gorm:"not null;column:title"`
	Post  string `json:"post" gorm:"not null;column:post"`

	CreatedAt time.Time      `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deletedAt;index"`
}
