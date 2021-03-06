package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-version"
	"github.com/ibrahimakbar31/comment-api-go/controller/validation"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/core/model/db"
	"github.com/ibrahimakbar31/comment-api-go/service/postgres"
	"github.com/liip/sheriff"
	"github.com/spf13/viper"
)

//App struct
type App struct {
	Router *gin.Engine
	DB1    *db.DB1
}

//Output struct
type Output struct {
	Ctx        *gin.Context
	Err        error
	Group      string
	StructData interface{}
	App        *App
}

//Handler function to handle JSON response output
func (output Output) Handler() {
	if output.Err != nil {
		db := output.App.DB1
		var errData model.APIError
		err := db.Where("name = ?", output.Err.Error()).First(&errData).Error
		if err != nil {
			ResponseDefaultError(output.Ctx, output.Err)
		} else {
			marshalData, _ := GenerateMarshal([]string{"error"}, errData)
			output.Ctx.JSON(errData.Code, marshalData)
		}
	} else {
		var marshalData, _ = GenerateMarshal([]string{output.Group}, output.StructData)
		output.Ctx.JSON(http.StatusOK, marshalData)
	}
}

//ResponseHandler function to handle response data from controller
func (app *App) ResponseHandler(f func(*gin.Context, *App) (interface{}, string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		interfaceData, group, err := f(c, app)
		responseData := Output{
			Ctx:        c,
			Err:        err,
			Group:      group,
			StructData: interfaceData,
			App:        app,
		}
		responseData.Handler()
	}
}

//ResponseDefaultError function
func ResponseDefaultError(c *gin.Context, err error) {
	var errData model.APIError
	errData.ID = 0
	errData.Code = 400
	errData.Name = "UNKNOWN_ERROR"
	errData.Message = err.Error()
	var marshalData, _ = GenerateMarshal([]string{"error"}, errData)
	c.JSON(errData.Code, marshalData)
}

//GenerateMarshal function to generate marshal interface
func GenerateMarshal(groups []string, structModel interface{}) (interface{}, error) {
	ver, _ := version.NewVersion(viper.GetString("Version"))
	o := &sheriff.Options{
		Groups:     groups,
		ApiVersion: ver,
	}
	outputInterface, err := sheriff.Marshal(o, structModel)
	if err != nil {
		return nil, err
	}
	return outputInterface, nil
}

//ValidateUser to validate user route
func (app *App) ValidateUser(validationName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch validationName {
		case "auth":
			var token Token
			reqToken := c.GetHeader("Authorization")
			if reqToken == "" {
				app.GenerateError(c, errors.New("TOKEN_INVALID"))
				c.Abort()
				return
			}
			splitToken := strings.Split(reqToken, "Bearer")
			if len(splitToken) != 2 {
				app.GenerateError(c, errors.New("TOKEN_INVALID"))
				c.Abort()
				return
			}
			reqToken = strings.TrimSpace(splitToken[1])
			token.Value = reqToken

			memberToken, err := token.Validate(app)
			if err != nil {
				app.GenerateError(c, err)
				c.Abort()
				return
			}
			c.Set("memberToken", memberToken)
			c.Next()
		case "memberInOrganization":
			orgCode := strings.ToLower(c.Param("orgCode"))
			memberToken, ok := c.MustGet("memberToken").(MemberToken)
			if !ok {
				app.GenerateError(c, errors.New("TOKEN_DATA_INVALID"))
				c.Abort()
				return
			}
			db := app.DB1
			organization, err := db.GetOrganizationByCode(orgCode)
			if err != nil {
				app.GenerateError(c, errors.New("ORGANIZATION_CODE_INVALID"))
				c.Abort()
				return
			}
			ok = validation.ValidMemberInOrganization(organization.Code, memberToken.Member)
			if !ok {
				app.GenerateError(c, errors.New("ORGANIZATION_UNAUTHORIZED"))
				c.Abort()
				return
			}

			c.Set("organization", organization)
			c.Next()
		default:
			c.Next()
		}
	}
}

//ValidateInput function
func (app *App) ValidateInput(validationName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch validationName {
		case "getPaginationQuery":
			var err error
			var page int64
			var perPage int64
			strPage := govalidator.Trim(c.DefaultQuery("page", "0"), "")
			strPerPage := govalidator.Trim(c.DefaultQuery("per_page", "0"), "")
			if strPage == "" {
				strPage = "0"
			}
			page, err = strconv.ParseInt(strPage, 10, 64)
			if err != nil {
				app.GenerateError(c, errors.New("PAGE_VALUE_INVALID"))
				c.Abort()
				return
			}
			if strPerPage == "" {
				strPerPage = "0"
			}
			perPage, err = strconv.ParseInt(strPerPage, 10, 64)
			if err != nil {
				app.GenerateError(c, errors.New("PER_PAGE_VALUE_INVALID"))
				c.Abort()
				return
			}
			pagination := model.Pagination{
				Page:    page,
				PerPage: perPage,
			}

			if pagination.Page > 0 {
				if pagination.PerPage == 0 {
					app.GenerateError(c, errors.New("PER_PAGE_MUST_SET"))
					c.Abort()
					return
				}
			}
			if pagination.PerPage > 0 {
				if pagination.Page == 0 {
					pagination.Page = 1
				}
			}
			c.Set("pagination", pagination)
			c.Next()
		default:
			c.Next()
		}
	}
}

//GenerateError to generate error output
func (app *App) GenerateError(c *gin.Context, err error) {
	responseData := Output{
		Ctx:        c,
		Err:        err,
		Group:      "error",
		StructData: err,
		App:        app,
	}
	responseData.Handler()
}

// GetDB function to get a database connection
func (app *App) GetDB() error {
	var err error
	app.DB1, err = postgres.ConnectDB1()
	if err != nil {
		return err
		//return errors.New("cannot connect DB")
	}
	return nil
}
