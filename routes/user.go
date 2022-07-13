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

func HandleInsertUser(repo rel.Repository) func(echo.Context) error {
	return func(ctx echo.Context) error {
		context := context.Background()

		user := entity.User{
			Name:   "Andrei",
			Active: true,
		}

		err := repo.Insert(context, &user)
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		return ctx.String(http.StatusOK, "Deu certo")
	}
}

func HandleGetUserByID(repo rel.Repository, userID uint32) func(echo.Context) error {
	return func(ctx echo.Context) error {
		context := context.Background()

		var user entity.User

		err := repo.Preload(context, &user, "contact_infos")
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		err = repo.Find(context, &user, where.Eq("id", userID))
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		return ctx.JSON(http.StatusOK, user)
	}
}
