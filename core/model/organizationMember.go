package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//OrganizationMember struct. Define organization member struct
type OrganizationMember struct {
	ID             uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" groups:""`
	OrganizationID uuid.NullUUID `gorm:"type:uuid" json:"" groups:""`
	Organization   Organization  `gorm:"foreignkey:ID;association_foreignkey:OrganizationID" json:"organization" groups:"member"`
	MemberID       uuid.NullUUID `gorm:"type:uuid" json:"" groups:""`
	Member         Member        `gorm:"foreignkey:ID;association_foreignkey:MemberID" json:"member" groups:""`
	CreatedAt      time.Time     `gorm:"" json:"created_at" groups:"member"`
	UpdatedAt      time.Time     `gorm:"" json:"updated_at" groups:""`
	DeletedAt      *time.Time    `gorm:"" json:"deleted_at" groups:""`
}
