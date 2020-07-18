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

func CheckErrorToken(t *testing.T, app *middleware.App, url string, method string, inputToken string, errorNameExpected string) {
	var jsonStr = []byte(`{}`)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", inputToken)
	app.Router.ServeHTTP(w, req)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, w.Code, 401)
	var apiError model.APIError
	MarshallingTestErrorResponse(t, w.Body.String(), &apiError)
	if apiError.Name != errorNameExpected {
		t.Errorf("invalid api error name. Got: " + apiError.Name)
	}
}

func CheckErrorOrganizationCode(t *testing.T, app *middleware.App, url string, method string, inputToken string, errorNameExpected string) {
	var jsonStr = []byte(`{}`)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+inputToken)
	app.Router.ServeHTTP(w, req)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, w.Code, 401)
	var apiError model.APIError
	MarshallingTestErrorResponse(t, w.Body.String(), &apiError)
	if apiError.Name != errorNameExpected {
		t.Errorf("invalid api error name. Got: " + apiError.Name)
	}
}

func CheckErrorInputCommentAndMember(t *testing.T, app *middleware.App, url string, method string, inputToken string, inputRequest string, errorNameExpected string) {
	var jsonStr = []byte(inputRequest)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+inputToken)
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

func CheckSuccessCommentAndMember(t *testing.T, app *middleware.App, url string, method string, inputToken string, inputRequest string) {
	var jsonStr = []byte(inputRequest)
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+inputToken)
	app.Router.ServeHTTP(w, req)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, w.Code, 200)
	if method == "POST" {
		var commentResponse v1.CommentCreateResponse
		byteData := []byte(w.Body.String())
		err := json.Unmarshal(byteData, &commentResponse)
		if err != nil {
			t.Errorf(err.Error())
		}
		app.DB1.Unscoped().Delete(&commentResponse.Comment)
	} else if method == "DELETE" {
		app.DB1.Exec("UPDATE comments SET deleted_at = NULL WHERE deleted_at is NOT NULL")
	}
}
