package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

func TestErrorInputLogin(t *testing.T) {
	app := GetAppTest(t)
	CheckErrorInputLogin(t, app, `{}`, "LOGIN_ID_VALUE_INVALID")
	CheckErrorInputLogin(t, app, `{"login_id":"ibr"}`, "LOGIN_ID_VALUE_INVALID")
	CheckErrorInputLogin(t, app, `{"login_id":"ibrahi%$#","password":"test123"}`, "LOGIN_ID_VALUE_INVALID")
	CheckErrorInputLogin(t, app, `{"login_id":"ibrahim1","password":"test123"}`, "LOGIN_INVALID")
	CheckErrorInputLogin(t, app, `{"login_id":"ibrahim1@gmail.com","password":"12345"}`, "LOGIN_INVALID")
}

func CheckErrorInputLogin(t *testing.T, app *middleware.App, input string, errorNameExpected string) {
	var jsonStr = []byte(input)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(jsonStr))
	app.Router.ServeHTTP(w, req)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, w.Code, 400)
	var apiError model.APIError
	MarshallingTestErrorResponse(t, w.Body.String(), &apiError)
	if apiError.Name != errorNameExpected {
		t.Errorf("invalid api error name. Got: " + apiError.Name)
	}
}
