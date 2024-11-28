package models

//Artist представляет структуру исполнителя
type Artist struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null;unique"`
	Songs []Song `gorm:"foreignKey:ArtistID"`
}

// Song представляет структуру песни
type Song struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Title    string   `json:"title" gorm:"not null;index" `
	ArtistID uint     `json:"artist_id"`
	Artist   Artist   `json"artist" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Album    string   `json:"album"`
	Lyrics   []string `json: "lyrics"`
}
