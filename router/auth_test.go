package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	v1 "github.com/ibrahimakbar31/comment-api-go/controller/v1"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

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

func MarshallingTestErrorResponse(t *testing.T, jsonData string, apiError *model.APIError) {
	byteData := []byte(jsonData)
	err := json.Unmarshal(byteData, apiError)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func CheckSuccessLogin(t *testing.T, app *middleware.App, input string) string {
	var jsonStr = []byte(input)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(jsonStr))
	app.Router.ServeHTTP(w, req)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, w.Code, 200)
	var authLoginResponse v1.AuthLoginResponse
	byteData := []byte(w.Body.String())
	err = json.Unmarshal(byteData, &authLoginResponse)
	if err != nil {
		t.Errorf(err.Error())
	}
	if authLoginResponse.Token.Value == "" {
		t.Errorf("token empty. Expected filled with token.")
	}
	_, err = authLoginResponse.Token.Validate(app)
	if err != nil {
		t.Errorf("token invalid. Expected valid token.")
	}

	return authLoginResponse.Token.Value
}
