package app

import (
	"crypto/subtle"
	"text/template"

	"github.com/dgraph-io/badger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var db *badger.DB

func Main() {
	var err error
	db, err = badger.Open(badger.DefaultOptions("pages.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := echo.New()
	app.Logger.SetLevel(log.DEBUG)
	app.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("pontaoski")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("friedenundliebesiegenimmer")) == 1 {
			return true, nil
		}
		return false, nil
	}))
	app.Renderer = &Template{
		templates: template.Must(template.ParseGlob("app/templates/*.html")),
	}
	app.GET("/view/:id", View)
	app.POST("/save/:id", Save)
	app.GET("/edit/:id", Edit)
	app.GET("/", List)
	app.GET("/pages", List)
	app.Logger.Fatal(app.Start(":8000"))
}
