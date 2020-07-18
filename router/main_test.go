package router

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
	"github.com/ibrahimakbar31/comment-api-go/service/postgres"
	"github.com/spf13/viper"
)

func TestRouter(t *testing.T) {
	app := GetAppTest(t)
	CheckErrorInputLogin(t, app, `{}`, "LOGIN_ID_VALUE_INVALID")
	CheckErrorInputLogin(t, app, `{"login_id":"ibr"}`, "LOGIN_ID_VALUE_INVALID")
	CheckErrorInputLogin(t, app, `{"login_id":"ibrahi%$#","password":"test123"}`, "LOGIN_ID_VALUE_INVALID")
	CheckErrorInputLogin(t, app, `{"login_id":"ibrahim1","password":"test123"}`, "LOGIN_INVALID")
	CheckErrorInputLogin(t, app, `{"login_id":"ibrahim1@gmail.com","password":"12345"}`, "LOGIN_INVALID")
	tokenValue := CheckSuccessLogin(t, app, `{"login_id":"ibrahim1@test.com","password":"12345"}`)

	CheckErrorToken(t, app, "/v1/orgs/xendit1/comments/", "POST", "", "TOKEN_INVALID")
	CheckErrorToken(t, app, "/v1/orgs/xendit1/comments/", "POST", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNZW1iZXJJRCI6IjAxZDRhMWExLWRkZTUtNDg1Mi1hMGRkLWZhODVlMzU2NTdlNCIsImV4cCI6MTU5NTEwNTAwNH0.shXvCCTmZMnXaveF6AX0OEwBvz7ONZIAb08KVGXu-QU", "TOKEN_INVALID")
	CheckErrorToken(t, app, "/v1/orgs/xendit1/comments/", "GET", "", "TOKEN_INVALID")
	CheckErrorToken(t, app, "/v1/orgs/xendit1/comments/", "GET", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNZW1iZXJJRCI6IjAxZDRhMWExLWRkZTUtNDg1Mi1hMGRkLWZhODVlMzU2NTdlNCIsImV4cCI6MTU5NTEwNTAwNH0.shXvCCTmZMnXaveF6AX0OEwBvz7ONZIAb08KVGXu-QU", "TOKEN_INVALID")
	CheckErrorToken(t, app, "/v1/orgs/xendit1/comments/", "DELETE", "", "TOKEN_INVALID")
	CheckErrorToken(t, app, "/v1/orgs/xendit1/comments/", "DELETE", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNZW1iZXJJRCI6IjAxZDRhMWExLWRkZTUtNDg1Mi1hMGRkLWZhODVlMzU2NTdlNCIsImV4cCI6MTU5NTEwNTAwNH0.shXvCCTmZMnXaveF6AX0OEwBvz7ONZIAb08KVGXu-QU", "TOKEN_INVALID")
	CheckErrorToken(t, app, "/v1/orgs/xendit1/members/", "GET", "", "TOKEN_INVALID")
	CheckErrorToken(t, app, "/v1/orgs/xendit1/members/", "GET", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNZW1iZXJJRCI6IjAxZDRhMWExLWRkZTUtNDg1Mi1hMGRkLWZhODVlMzU2NTdlNCIsImV4cCI6MTU5NTEwNTAwNH0.shXvCCTmZMnXaveF6AX0OEwBvz7ONZIAb08KVGXu-QU", "TOKEN_INVALID")

	CheckErrorOrganizationCode(t, app, "/v1/orgs/xendit2/comments/", "POST", tokenValue, "ORGANIZATION_UNAUTHORIZED")
	CheckErrorOrganizationCode(t, app, "/v1/orgs/xendit12/comments/", "POST", tokenValue, "ORGANIZATION_CODE_INVALID")
	CheckErrorOrganizationCode(t, app, "/v1/orgs/xendit2/comments/", "GET", tokenValue, "ORGANIZATION_UNAUTHORIZED")
	CheckErrorOrganizationCode(t, app, "/v1/orgs/xendit12/comments/", "GET", tokenValue, "ORGANIZATION_CODE_INVALID")
	CheckErrorOrganizationCode(t, app, "/v1/orgs/xendit2/comments/", "DELETE", tokenValue, "ORGANIZATION_UNAUTHORIZED")
	CheckErrorOrganizationCode(t, app, "/v1/orgs/xendit12/comments/", "DELETE", tokenValue, "ORGANIZATION_CODE_INVALID")
	CheckErrorOrganizationCode(t, app, "/v1/orgs/xendit2/members/", "GET", tokenValue, "ORGANIZATION_UNAUTHORIZED")
	CheckErrorOrganizationCode(t, app, "/v1/orgs/xendit12/members/", "GET", tokenValue, "ORGANIZATION_CODE_INVALID")

	CheckErrorInputCommentAndMember(t, app, "/v1/orgs/xendit1/comments/", "POST", tokenValue, `{}`, "COMMENT_VALUE_INVALID")
	CheckErrorInputCommentAndMember(t, app, "/v1/orgs/xendit1/comments/?page=1&per_page=", "GET", tokenValue, `{}`, "PER_PAGE_MUST_SET")
	CheckErrorInputCommentAndMember(t, app, "/v1/orgs/xendit1/members/?page=1&per_page=", "GET", tokenValue, `{}`, "PER_PAGE_MUST_SET")

	CheckSuccessCommentAndMember(t, app, "/v1/orgs/xendit1/comments/", "POST", tokenValue, `{"value":"123abc"}`)
	CheckSuccessCommentAndMember(t, app, "/v1/orgs/xendit1/comments/", "GET", tokenValue, `{}`)
	CheckSuccessCommentAndMember(t, app, "/v1/orgs/xendit1/comments/", "DELETE", tokenValue, `{}`)
	CheckSuccessCommentAndMember(t, app, "/v1/orgs/xendit1/members/", "GET", tokenValue, `{}`)
}

func GetAppTest(t *testing.T) *middleware.App {
	gin.SetMode(gin.TestMode)
	var app middleware.App
	viper.Set("Env", "Test")
	GetConfigFileTest(t)
	err := app.GetDB()
	if err != nil {
		t.Errorf(err.Error())
	}
	println("checking migrating table....")
	postgres.DB1Migration(app.DB1)
	println("checking migration done.")
	InitAllRoutes(&app)
	return &app
}

func GetConfigFileTest(t *testing.T) {
	getFilename := os.Getenv("GOCONFIGFILENAME")
	if getFilename == "" {
		getFilename = "config"
	}
	viper.SetConfigName(getFilename)
	viper.SetConfigType("json")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	if err != nil {
		t.Errorf("CANNOT_READ_CONFIG_FILE")
	}
}
