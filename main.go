package main

import (
	"poc-orm/db/adapter"
	"poc-orm/routes"

	"github.com/go-rel/rel"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", routes.HandleHelthCheck)

	adapter, err := adapter.GetRelAdapter()
	if err != nil {
		panic(err)
	}

	repo := rel.New(adapter)

	e.GET("/insert-author", routes.HandleInsertAuthor(repo))
	e.GET("/get-author", routes.HandleGetAuthorByID(repo, 2))
	e.GET("/insert-books", routes.HandleInsertBook(repo))
	e.GET("/get-book", routes.HandleGetBook(repo, 1))

	e.Logger.Fatal(e.Start(":1323"))
}
