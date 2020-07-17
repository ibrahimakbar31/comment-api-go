package v1

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/ibrahimakbar31/comment-api-go/controller/validation"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
	"golang.org/x/crypto/bcrypt"
)

//LoginCredential struct
type LoginCredential struct {
	LoginID  string `json:"login_id" valid:"stringlength(5|200)~LOGIN_ID_VALUE_INVALID" groups:"member"`
	Password string `json:"password" valid:"stringlength(1|200)~PASSWORD_VALUE_INVALID" groups:"member"`
}

//AuthLoginResponse struct
type AuthLoginResponse struct {
	Message string           `default:"success" json:"message" groups:"login"`
	Token   middleware.Token `json:"token" groups:"login"`
}

//SubmitLogin function
func SubmitLogin(c *gin.Context, app *middleware.App) (interface{}, string, error) {
	var memberToken middleware.MemberToken
	var err error
	var response AuthLoginResponse
	var loginCredential LoginCredential
	group := "login"
	c.ShouldBindJSON(&loginCredential)
	db := app.DB1
	loginType := "username"
	_, err = govalidator.ValidateStruct(loginCredential)
	if err != nil {
		errs := err.(govalidator.Errors).Errors()
		return response, group, errs[0]
	}
	checkIsEmail := validation.ValidEmailFormat(loginCredential.LoginID)
	if checkIsEmail == true {
		loginType = "email"
	}

	if loginType == "username" {
		checkIsUsername := validation.ValidUsernameFormat(loginCredential.LoginID)
		if checkIsUsername == false {
			return response, group, errors.New("LOGIN_ID_VALUE_INVALID")
		}
	}

	err = db.Where(loginType+" = ?", loginCredential.LoginID).Preload("Language").Preload("OrganizationMembers").Preload("OrganizationMembers.Organization").First(&memberToken.Member).Error
	if err != nil {
		return response, group, errors.New("LOGIN_INVALID")
	}

	checkPwd := ComparePassword(memberToken.Member.Password, []byte(loginCredential.Password))
	if checkPwd == false {
		return response, group, errors.New("LOGIN_INVALID")
	}

	memberToken, err = middleware.GenerateToken(memberToken.Member)
	if err != nil {
		return response, group, errors.New("LOGIN_TOKEN_ERROR")
	}

	response.Token = memberToken.Token
	return response, group, err
}

//HashAndSaltPassword function
func HashAndSaltPassword(rawPwd string) (string, error) {
	pwd := []byte(rawPwd)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
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
