package models

type Tag struct {
	Id        uint   `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}
