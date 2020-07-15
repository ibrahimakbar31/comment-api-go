package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

//Member struct. Define member struct
type Member struct {
	ID                  uuid.UUID            `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id" groups:"member,membersOrganization"`
	Email               string               `gorm:"size:200;unique_index;" json:"email" groups:"member,orgComments,membersOrganization"`
	Username            string               `gorm:"size:200;unique_index" json:"username" groups:"member,orgComments,membersOrganization"`
	Name                string               `gorm:"size:100" json:"name" groups:"member,orgComments,membersOrganization"`
	Avatar              string               `gorm:"size:255" json:"avatar" groups:"member,membersOrganization,orgComments"`
	LanguageID          string               `gorm:"size:5" json:"" groups:""`
	Language            Language             `gorm:"foreignkey:ID;association_foreignkey:LanguageID" json:"language" groups:"member,membersOrganization"`
	Password            string               `gorm:"size:255" json:"" groups:""`
	FollowerCount       int64                `gorm:"index:fk_member_follower" json:"follower_count" groups:"membersOrganization"`
	FollowingCount      int64                `gorm:"" json:"following_count" groups:"membersOrganization"`
	OrganizationMembers []OrganizationMember `gorm:"foreignkey:MemberID;association_foreignkey:ID" json:"organization_members" groups:"member"`
	CreatedAt           time.Time            `gorm:"" json:"created_at" groups:"member,membersOrganization"`
	UpdatedAt           time.Time            `gorm:"" json:"updated_at" groups:"member,membersOrganization"`
	DeletedAt           *time.Time           `gorm:"" json:"deleted_at" groups:"member"`
}

//AfterFind member hook function
func (member *Member) AfterFind() (err error) {
	avatarURL := viper.GetString(viper.GetString("Env") + ".AvatarURL")
	fullPath := avatarURL + "/" + member.Avatar
	member.Avatar = fullPath
	return
}

//MembersPagination struct
type MembersPagination struct {
	Members    []Member   `json:"members" groups:"membersOrganization"`
	Pagination Pagination `json:"pagination" groups:"membersOrganization"`
}
