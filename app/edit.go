package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Save(c echo.Context) error {
	id := c.Param("id")
	body := c.FormValue("body")
	p := &Page{Title: id, Body: body}
	p.save()
	return c.Redirect(http.StatusFound, "/view/"+id)
}

func Edit(c echo.Context) error {
	id := c.Param("id")
	p, err := loadPage(id)
	if err != nil {
		p = &Page{Title: id}
	}
	return c.Render(http.StatusOK, "edit", p)
}
