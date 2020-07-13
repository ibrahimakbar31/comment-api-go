package db

import (
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	uuid "github.com/satori/go.uuid"
)

//GetMemberByID function
func (db *DB1) GetMemberByID(memberID uuid.UUID) (model.Member, error) {
	var member model.Member
	var err error

	err = db.Where("id = ?", memberID).Preload("OrganizationMembers").Preload("OrganizationMembers.Organization").First(&member).Error
	if err != nil {
		return member, err
	}

	return member, err
}

//GetMembersByOrganizationID function
func (db *DB1) GetMembersByOrganizationID(organizationID uuid.UUID, pagination model.Pagination) ([]model.Member, error) {
	var err error
	var members []model.Member

	query := db.Model(&members).Joins("left join organization_members on organization_members.member_id = members.id")
	if pagination.PerPage > 0 || pagination.Page > 0 {
		if pagination.Page == 0 {
			pagination.Page = 1
		}
		limit := pagination.PerPage
		var offset int64
		if pagination.Page > 1 {
			offset = (pagination.Page - 1) * limit
		}
		query = query.Limit(limit).Offset(offset)
	}
	query = query.Where("organization_members.organization_id = ?", organizationID.String()).Where("organization_members.deleted_at is NULL").Preload("Language").Order("members.follower_count desc").Find(&members)

	return members, err
}
