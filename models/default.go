package models

import "time"

type GORMModel struct {
	CreatedAt time.Time `gorm:"type:datetime,autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:datetime,autoUpdateTime" json:"updated_at,omitempty"`
}
