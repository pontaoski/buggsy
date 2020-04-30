package app

import (
	"net/http"

	"github.com/dgraph-io/badger"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	var names []string
	db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			names = append(names, string(it.Item().Key()))
		}
		return nil
	})
	return c.Render(http.StatusOK, "list", names)
}

func View(c echo.Context) error {
	id := c.Param("id")
	p, err := loadPage(id)
	if err != nil {
		return c.Redirect(http.StatusFound, "/edit/"+id)
	}
	return c.Render(http.StatusOK, "view", p)
}
