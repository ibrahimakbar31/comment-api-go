package db

import (
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	uuid "github.com/satori/go.uuid"
)

//CreateComment function
func (db *DB1) CreateComment(comment *model.Comment) error {
	err := db.Create(&comment).Error
	return err
}

//GetCommentsByOrganizationID function
func (db *DB1) GetCommentsByOrganizationID(organizationID uuid.UUID, pagination model.Pagination) ([]model.Comment, error) {
	var comments []model.Comment
	var err error
	comment := model.Comment{
		OrganizationID: uuid.NullUUID{
			UUID:  organizationID,
			Valid: true,
		},
	}
	query := db.Model(&comments)
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
	query = query.Where(comment).Preload("Member").Order("created_at desc").Find(&comments)

	return comments, err
}

//DeleteCommentsByOrganizationID function
func (db *DB1) DeleteCommentsByOrganizationID(organizationID uuid.UUID) (int64, error) {
	var deleteComment int64
	var err error
	comment := model.Comment{
		OrganizationID: uuid.NullUUID{
			UUID:  organizationID,
			Valid: true,
		},
	}
	deleteComment = db.Where(comment).Delete(&model.Comment{}).RowsAffected
	return deleteComment, err
}
