package postgres

import (
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/core/model/db"
	uuid "github.com/satori/go.uuid"
)

//DB1Migration function
func DB1Migration(db *db.DB1) {
	addData := false
	if db.HasTable(&model.Language{}) == false {
		addData = true
	}

	db.AutoMigrate(&model.Language{}, &model.APIError{}, &model.APIErrorMessage{}, &model.Organization{}, &model.Member{}, &model.OrganizationMember{}, &model.Comment{})

	db.Model(&model.APIErrorMessage{}).AddForeignKey("language_id", "languages(id)", "CASCADE", "CASCADE")
	db.Model(&model.APIErrorMessage{}).AddForeignKey("api_error_id", "api_errors(id)", "CASCADE", "CASCADE")

	db.Model(&model.Member{}).AddForeignKey("language_id", "languages(id)", "CASCADE", "CASCADE")

	db.Model(&model.Comment{}).AddForeignKey("member_id", "members(id)", "CASCADE", "CASCADE")
	db.Model(&model.Comment{}).AddForeignKey("organization_id", "organizations(id)", "CASCADE", "CASCADE")

	db.Model(&model.OrganizationMember{}).AddForeignKey("organization_id", "organizations(id)", "CASCADE", "CASCADE")
	db.Model(&model.OrganizationMember{}).AddForeignKey("member_id", "members(id)", "CASCADE", "CASCADE")

	if addData == true {
		db.Create(&model.Language{
			ID:   "EN",
			Name: "English",
		})
		db.Create(&model.Language{
			ID:   "ID",
			Name: "Indonesia",
		})

		db.Create(&model.APIError{
			Code:    404,
			Name:    "PAGE_NOT_FOUND",
			Message: "Page Not Found",
		})

		db.Create(&model.APIError{
			Code:    401,
			Name:    "UNAUTHORIZED",
			Message: "Unauthorized",
		})

		db.Create(&model.APIError{
			Code:    400,
			Name:    "LOGIN_ID_VALUE_INVALID",
			Message: "Invalid Login ID. Login ID must have length between 1 and 200. Only accept underscore with alphanumeric combination or email format",
		})

		db.Create(&model.APIError{
			Code:    400,
			Name:    "PASSWORD_VALUE_INVALID",
			Message: "Password invalid. Please provide valid password",
		})

		db.Create(&model.APIError{
			Code:    400,
			Name:    "LOGIN_INVALID",
			Message: "Invalid Login ID / Password combination",
		})

		db.Create(&model.APIError{
			Code:    401,
			Name:    "TOKEN_INVALID",
			Message: "Invalid Token. Please provide bearer token on authorization header",
		})

		db.Create(&model.APIError{
			Code:    401,
			Name:    "ORGANIZATION_UNAUTHORIZED",
			Message: "Unauthorized organization",
		})

		db.Create(&model.APIError{
			Code:    401,
			Name:    "ORGANIZATION_CODE_INVALID",
			Message: "Invalid Organization Code",
		})

		db.Create(&model.APIError{
			Code:    400,
			Name:    "COMMENT_VALUE_INVALID",
			Message: "Comment Value must have string value between 3 - 1000",
		})

		db.Create(&model.APIError{
			Code:    400,
			Name:    "PER_PAGE_MUST_SET",
			Message: "Per Page value must set",
		})

		db.Create(&model.APIError{
			Code:    400,
			Name:    "PER_PAGE_VALUE_INVALID",
			Message: "Invalid per Page value",
		})

		db.Create(&model.APIError{
			Code:    400,
			Name:    "PAGE_VALUE_INVALID",
			Message: "Invalid page value",
		})

		db.Create(&model.Organization{
			Code: "xendit1",
			Name: "Xendit One",
		})

		db.Create(&model.Organization{
			Code: "xendit2",
			Name: "Xendit Two",
		})

		db.Create(&model.Member{
			Email:          "ibrahim1@test.com",
			Username:       "ibrahim1",
			Name:           "Ibrahim One",
			Avatar:         "ibrahim1.jpg",
			LanguageID:     "EN",
			Password:       "$2a$04$XYyDwLh4C39VLqs0bB2SOuLwii4sqKD2yCSI4QW11bMyW5PPJ6wzG",
			FollowerCount:  50,
			FollowingCount: 24,
		})

		db.Create(&model.Member{
			Email:          "ibrahim2@test.com",
			Username:       "ibrahim2",
			Name:           "Ibrahim Two",
			Avatar:         "ibrahim2.jpg",
			LanguageID:     "EN",
			Password:       "$2a$04$XYyDwLh4C39VLqs0bB2SOuLwii4sqKD2yCSI4QW11bMyW5PPJ6wzG",
			FollowerCount:  15,
			FollowingCount: 31,
		})

		db.Create(&model.Member{
			Email:          "ibrahim3@test.com",
			Username:       "ibrahim3",
			Name:           "Ibrahim Three",
			Avatar:         "ibrahim3.jpg",
			LanguageID:     "EN",
			Password:       "$2a$04$XYyDwLh4C39VLqs0bB2SOuLwii4sqKD2yCSI4QW11bMyW5PPJ6wzG",
			FollowerCount:  21,
			FollowingCount: 42,
		})

		db.Create(&model.Member{
			Email:          "ibrahim4@test.com",
			Username:       "ibrahim4",
			Name:           "Ibrahim Four",
			Avatar:         "ibrahim4.jpg",
			LanguageID:     "EN",
			Password:       "$2a$04$XYyDwLh4C39VLqs0bB2SOuLwii4sqKD2yCSI4QW11bMyW5PPJ6wzG",
			FollowerCount:  73,
			FollowingCount: 64,
		})

		db.Create(&model.Member{
			Email:          "ibrahim5@test.com",
			Username:       "ibrahim5",
			Name:           "Ibrahim Five",
			Avatar:         "ibrahim5.jpg",
			LanguageID:     "EN",
			Password:       "$2a$04$XYyDwLh4C39VLqs0bB2SOuLwii4sqKD2yCSI4QW11bMyW5PPJ6wzG",
			FollowerCount:  61,
			FollowingCount: 11,
		})

		var ibrahim1 model.Member
		var ibrahim2 model.Member
		var ibrahim3 model.Member
		var ibrahim4 model.Member
		var ibrahim5 model.Member

		var xendit1 model.Organization
		var xendit2 model.Organization

		db.Where("username = ?", "ibrahim1").First(&ibrahim1)
		db.Where("username = ?", "ibrahim2").First(&ibrahim2)
		db.Where("username = ?", "ibrahim3").First(&ibrahim3)
		db.Where("username = ?", "ibrahim4").First(&ibrahim4)
		db.Where("username = ?", "ibrahim5").First(&ibrahim5)

		db.Where("code = ?", "xendit1").First(&xendit1)
		db.Where("code = ?", "xendit2").First(&xendit2)

		db.Create(&model.OrganizationMember{
			OrganizationID: uuid.NullUUID{
				UUID:  xendit1.ID,
				Valid: true},
			MemberID: uuid.NullUUID{
				UUID:  ibrahim1.ID,
				Valid: true},
		})

		db.Create(&model.OrganizationMember{
			OrganizationID: uuid.NullUUID{
				UUID:  xendit1.ID,
				Valid: true},
			MemberID: uuid.NullUUID{
				UUID:  ibrahim2.ID,
				Valid: true},
		})

		db.Create(&model.OrganizationMember{
			OrganizationID: uuid.NullUUID{
				UUID:  xendit2.ID,
				Valid: true},
			MemberID: uuid.NullUUID{
				UUID:  ibrahim3.ID,
				Valid: true},
		})

		db.Create(&model.OrganizationMember{
			OrganizationID: uuid.NullUUID{
				UUID:  xendit2.ID,
				Valid: true},
			MemberID: uuid.NullUUID{
				UUID:  ibrahim4.ID,
				Valid: true},
		})

		db.Create(&model.OrganizationMember{
			OrganizationID: uuid.NullUUID{
				UUID:  xendit1.ID,
				Valid: true},
			MemberID: uuid.NullUUID{
				UUID:  ibrahim5.ID,
				Valid: true},
		})

		db.Create(&model.OrganizationMember{
			OrganizationID: uuid.NullUUID{
				UUID:  xendit2.ID,
				Valid: true},
			MemberID: uuid.NullUUID{
				UUID:  ibrahim5.ID,
				Valid: true},
		})
	}
}
