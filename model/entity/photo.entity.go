package entity

import "time"

type Photo struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Image      string    `json:"image"`
	CategoryId uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"-" gorm:"index,column:deleted_at"`
}
