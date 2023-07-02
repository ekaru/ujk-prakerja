package models

type Article struct {
	Id         uint     `gorm:"primaryKey" json:"id"`
	Title      string   `gorm:"type:longtext;not null" json:"title" validate:"required,min=5"`
	Content    string   `gorm:"type:longtext;not null" json:"content" validate:"required,min=10"`
	UserID     uint     `gorm:"not null" json:"user_id"`
	User       User     `gorm:"foreignKey:UserID" json:"user" validate:"-"`
	CategoryID uint     `gorm:"not null" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category" validate:"-"`
	GORMModel
}

type ArticleResponse struct {
	Id         uint             `json:"id"`
	Title      string           `json:"title"`
	Content    string           `json:"content"`
	UserID     uint             `json:"-"`
	User       UserResponse     `gorm:"foreignKey:UserID" json:"user"`
	CategoryID uint             `json:"-"`
	Category   CategoryResponse `gorm:"foreignKey:CategoryID" json:"category"`
	GORMModel
}
