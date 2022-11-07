package models

type Todo struct {
	ID     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Text   string `json:"title"`
	Status string `gorm:"type:varchar(255)" json:"status"`
}
