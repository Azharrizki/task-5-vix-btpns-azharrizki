package models

type Photo struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Caption  string `gorm:"type:text" json:"caption"`
	PhotoUrl string `gorm:"type:varchar(255)" json:"photo_url"`
	UserId   int64  `gorm:"type:bigint" json:"user_id"`
}
