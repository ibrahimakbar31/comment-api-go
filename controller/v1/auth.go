package v1

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/ibrahimakbar31/comment-api-go/controller/validation"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
	"golang.org/x/crypto/bcrypt"
)

//LoginCredential struct
type LoginCredential struct {
	LoginID  string `json:"login_id" valid:"stringlength(5|200)~LOGIN_ID_VALUE_INVALID" groups:"member"`
	Password string `json:"password" valid:"stringlength(1|200)~PASSWORD_VALUE_INVALID" groups:"member"`
}

//LoginSubmit function
func LoginSubmit(loginCredential LoginCredential, app *middleware.App) (middleware.MemberToken, error) {
	var memberToken middleware.MemberToken
	var err error
	db := app.DB1
	loginType := "username"
	_, err = govalidator.ValidateStruct(loginCredential)
	if err != nil {
		return memberToken, err
	}
	checkIsEmail := validation.EmailFormat(loginCredential.LoginID)
	if checkIsEmail == true {
		loginType = "email"
	}

	if loginType == "username" {
		checkIsUsername := validation.UsernameFormat(loginCredential.LoginID)
		if checkIsUsername == false {
			return memberToken, errors.New("LOGIN_ID_VALUE_INVALID")
		}
	}

	err = db.Where(loginType+" = ?", loginCredential.LoginID).Preload("Language").Preload("OrganizationMembers").Preload("OrganizationMembers.Organization").First(&memberToken.Member).Error
	if err != nil {
		return memberToken, errors.New("LOGIN_INVALID")
	}

	checkPwd := ComparePassword(memberToken.Member.Password, []byte(loginCredential.Password))
	if checkPwd == false {
		return memberToken, errors.New("LOGIN_INVALID")
	}

	memberToken, err = middleware.TokenGenerate(memberToken.Member)
	if err != nil {
		return memberToken, errors.New("LOGIN_TOKEN_ERROR")
	}

	return memberToken, err
}

//HashAndSaltPassword function
func HashAndSaltPassword(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash), err
}

//ComparePassword function
func ComparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
