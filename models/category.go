package models

type Category struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:15; not null" json:"name" validate:"required"`
	Description string `gorm:"type:text" json:"description"`
	GORMModel
}

type CategoryResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
