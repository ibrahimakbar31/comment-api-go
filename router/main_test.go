package router

/*func TestHandleNoRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var app middleware.App
	InitAllRoutes(&app)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testnoroute", nil)
	app.Router.ServeHTTP(w, req)
	if w.Code == 404 {
		t.Logf("test pass, expected got 404 error")
	} else {
		t.Errorf("handle no route failed")
	}
}*/
