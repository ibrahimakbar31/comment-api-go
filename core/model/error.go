package model

//APIError struct. Define error API struct
type APIError struct {
	ID       int64             `gorm:"primary_key;" json:"id" groups:""`
	Code     int               `gorm:"" json:"code" groups:"error"`
	Name     string            `gorm:"size:255;unique_index" json:"name" groups:"error"`
	Message  string            `gorm:"size:255" json:"message" groups:"error"`
	Messages []APIErrorMessage `gorm:"foreignkey:ErrorID;association_foreignkey:ID" json:"messages" groups:""`
}

//APIErrorMessage struct. Define error message with language
type APIErrorMessage struct {
	ID         int64    `gorm:"primary_key;" json:"id" groups:"error"`
	APIErrorID int64    `gorm:"" json:"" groups:""`
	APIError   APIError `gorm:"foreignkey:ID;association_foreignkey:APIErrorID" json:"api_error" groups:"error"`
	LanguageID string   `gorm:"size:5" json:"" groups:""`
	Language   Language `gorm:"foreignkey:ID;association_foreignkey:LanguageID" json:"language" groups:"error"`
	Text       string   `gorm:"size:255" json:"text" groups:"error"`
}
