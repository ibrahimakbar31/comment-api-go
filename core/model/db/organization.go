package db

import "github.com/ibrahimakbar31/comment-api-go/core/model"

//GetOrganizationByCode function
func (db *DB1) GetOrganizationByCode(organizationCode string) (model.Organization, error) {
	var organization model.Organization
	var err error

	err = db.Where("code = ?", organizationCode).First(&organization).Error
	if err != nil {
		return organization, err
	}

	return organization, err
}
