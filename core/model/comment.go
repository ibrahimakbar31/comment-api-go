package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Comment struct
type Comment struct {
	ID             uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" groups:"comment,orgComments"`
	Value          string        `gorm:"type:text;size:1000;" json:"value" groups:"comment,orgComments"`
	MemberID       uuid.NullUUID `gorm:"type:uuid;" json:"" groups:""`
	Member         Member        `gorm:"foreignkey:ID;association_foreignkey:MemberID" json:"member" groups:"orgComments"`
	OrganizationID uuid.NullUUID `gorm:"type:uuid;index:fk_comment_organization" json:"" groups:""`
	Organization   Organization  `gorm:"foreignkey:ID;association_foreignkey:OrganizationID;" json:"organization" groups:""`
	CreatedAt      time.Time     `gorm:"" json:"created_at" groups:"comment,orgComments"`
	UpdatedAt      time.Time     `gorm:"" json:"updated_at" groups:"comment,orgComments"`
	DeletedAt      *time.Time    `gorm:"" json:"deleted_at" groups:""`
}

//CommentCreate struct
type CommentCreate struct {
	Value string `valid:"stringlength(3|1000)~COMMENT_VALUE_INVALID" json:"value" groups:"comment"`
}

//CommentsPagination struct
type CommentsPagination struct {
	Comments   []Comment  `json:"comments" groups:"orgComments"`
	Pagination Pagination `json:"pagination" groups:"orgComments"`
}
