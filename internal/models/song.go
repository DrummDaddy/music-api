package models

// Song представляет структуру песни
type Song struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" gorm:"not null" `
	Artist string `json:"artist" gorm:"not null"`
	Album  string `json:"album"`
}
