package models

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null;size:255" json:"password"`
}