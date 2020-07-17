package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ibrahimakbar31/comment-api-go/core/model"
	"github.com/ibrahimakbar31/comment-api-go/middleware"
)

func TestNoInputLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var app middleware.App
	InitAllRoutes(&app)
	w := httptest.NewRecorder()
	//handler := http.HandlerFunc(HealthCheckHandler)
	var jsonStr = []byte(`{"login_id":"test"}`)
	req, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(jsonStr))
	app.Router.ServeHTTP(w, req)
	bodyStr := w.Body.String()
	//app.Router.
	var apiError model.APIError
	json.Unmarshal([]byte(bodyStr), &apiError)
	fmt.Printf("%+v\n", apiError)
	if w.Code == 400 {
		t.Logf("test pass, expected got 400 error")
	} else {
		t.Fatalf("wrong status code. expected 400")
	}
}
