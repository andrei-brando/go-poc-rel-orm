package routes

import (
	"context"
	"fmt"
	"net/http"
	"poc-orm/entity"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/labstack/echo/v4"
)

func HandleHelthCheck(c echo.Context) error {
	fmt.Println("teste")
	return c.String(http.StatusOK, "Hello, World!")
}

func HandleInsertAuthor(repo rel.Repository) func(echo.Context) error {
	return func(ctx echo.Context) error {
		context := context.Background()

		author := entity.Author{
			Name:   "Autor 2",
			Active: false,
		}

		err := repo.Insert(context, &author)
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		return ctx.String(http.StatusOK, "Deu certo")
	}
}

func HandleGetAuthorByID(repo rel.Repository, authorID uint32) func(echo.Context) error {
	return func(ctx echo.Context) error {
		context := context.Background()

		var author entity.Author

		err := repo.Find(context, &author, rel.Select("*", "books.*").JoinAssoc("books"), where.Eq("author_id", authorID))
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		return ctx.JSON(http.StatusOK, author)
	}
}
