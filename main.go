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
	e.GET("/get-author", routes.HandleGetAuthorByID(repo, 1))
	e.GET("/insert-books", routes.HandleInsertBook(repo))
	e.GET("/get-book", routes.HandleGetBook(repo, 1))

	/*
		e.GET("/insert-user", routes.HandleInsertUser(repo))
		e.GET("/get-user", routes.HandleGetUserByID(repo, 1))

		e.GET("/insert-info", routes.HandleInsertContactInfo(repo))
		e.GET("/get-info", routes.HandleGetContactInfoByID(repo, 1))
	*/

	// e.GET("/get-info", routes.HandleGet(repo))

	e.Logger.Fatal(e.Start(":1323"))
}
