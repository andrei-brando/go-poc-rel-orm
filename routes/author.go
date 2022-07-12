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

		err := repo.Find(context, &author, where.Eq("author_id", authorID))
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		return ctx.JSON(http.StatusOK, author)
	}
}

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

		var book entity.Book

		err := repo.Preload(context, &book, "author")
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		fmt.Printf("PRIMEIRO - %v", book)

		err = repo.Find(context, &book, where.Eq("book_id", bookID))
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		fmt.Printf("SEGUNDO - %v", book)

		return ctx.JSON(http.StatusOK, book)
	}
}
