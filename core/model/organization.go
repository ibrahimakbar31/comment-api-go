package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Organization struct. Define organization struct
type Organization struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" groups:"member"`
	Code      string     `gorm:"size:50;unique_index" json:"code" groups:"member"`
	Name      string     `gorm:"size:100" json:"name" groups:"member"`
	Comments  []Comment  `gorm:"foreignkey:CommentID;association_foreignkey:ID" json:"comments" groups:""`
	CreatedAt time.Time  `gorm:"" json:"created_at" groups:""`
	UpdatedAt time.Time  `gorm:"" json:"updated_at" groups:""`
	DeletedAt *time.Time `gorm:"" json:"deleted_at" groups:""`
}
