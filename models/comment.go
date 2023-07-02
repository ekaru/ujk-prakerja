package models

type Comment struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	ArticleId uint   `gorm:"not null" json:"id_article"`
	Username  string `gorm:"size:30; not null" json:"username"`
	Content   string `gorm:"type:longtext; not null" json:"content"`
	GORMModel
}
