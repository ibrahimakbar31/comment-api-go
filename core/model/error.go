package model

//APIError struct. Define error API struct
type APIError struct {
	ID      int64  `gorm:"primary_key;" json:"id" groups:""`
	Code    int    `gorm:"" json:"code" groups:"error"`
	Name    string `gorm:"size:255;unique_index" json:"name" groups:"error"`
	Message string `gorm:"size:255" json:"message" groups:"error"`
}
