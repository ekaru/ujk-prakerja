package models

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"size:20; unique; not null" json:"username" validate:"required"`
	Password string `gorm:"size:100; not null" json:"password" validate:"required,min=4"`
	Email    string `gorm:"size:30; unique; not null" json:"email" validate:"required,email"`
	GORMModel
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type UserResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
