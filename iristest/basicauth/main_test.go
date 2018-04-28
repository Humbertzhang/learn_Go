package main

import (
	"testing"
	"github.com/kataras/iris/httptest"
)

func TestBasicAuth(t *testing.T) {
	app := App()

	e := httptest.New(t, app)

	e.GET("/").Expect().Status(httptest.StatusUnauthorized)
	e.GET("/admin").WithBasicAuth("username", "password").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin username:password")
	e.GET("/admin/profile").WithBasicAuth("username", "password").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin/profile username:password")
	e.GET("/admin/settings").WithBasicAuth("username", "password").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin/settings username:password")

	e.GET("/admin/settings").WithBasicAuth("invalidusername", "invalidpassword").
		Expect().Status(httptest.StatusUnauthorized)
}