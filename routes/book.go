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

func HandleInsertBook(repo rel.Repository) func(echo.Context) error {
	return func(ctx echo.Context) error {
		context := context.Background()

		book := entity.Book{
			ID:       1,
			Title:    "Livro 1",
			AuthorID: 1,
		}

		err := repo.Insert(context, &book)
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		return ctx.String(http.StatusOK, "Deu certo")
	}
}

func HandleGetBook(repo rel.Repository, bookID uint32) func(echo.Context) error {
	return func(ctx echo.Context) error {
		context := context.Background()

		// var book entity.Book

		var author entity.Author

		repo.Find(context, &author, where.Eq("author_id", bookID))

		repo.Preload(context, &author, "books")

		// err := repo.Find(context, &book, where.Eq("book_id", bookID))
		// if err != nil {
		// 	fmt.Printf("Deu erro: %v", err.Error())
		// 	return nil
		// }

		// fmt.Printf("PRIMEIRO - %v", book)

		// err := repo.Find(context, &book, rel.Select("*", "author.*").JoinAssoc("author"), where.Eq("book_id", bookID))
		// if err != nil {
		// 	fmt.Printf("Deu erro: %v", err.Error())
		// 	return nil
		// }

		// var response []interface{}
		// response = append(response, author, book)

		return ctx.JSON(http.StatusOK, author)
	}
}
