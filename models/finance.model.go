package models

import (
	"time"

	"gorm.io/gorm"
)

type Finance struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Desc      string         `json:"desc" form:"desc" validate:"gte=6,lte=32" gorm:"not null"`
	Amount    int            `json:"amount" form:"amount" validate:"required" gorm:"not null"`
	TrxDate   time.Time      `json:"trxDate" form:"trxDate" validate:"required" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
