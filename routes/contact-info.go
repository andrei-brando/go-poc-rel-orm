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

func HandleInsertContactInfo(repo rel.Repository) func(echo.Context) error {
	return func(ctx echo.Context) error {
		context := context.Background()

		contactInfo := entity.ContactInfo{
			Email:  "teste@teste.com",
			Phone:  "(11) 11111-1111",
			UserID: 1,
		}

		err := repo.Insert(context, &contactInfo)
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		return ctx.String(http.StatusOK, "Deu certo")
	}
}

func HandleGetContactInfoByID(repo rel.Repository, contactInfoID uint32) func(echo.Context) error {
	return func(ctx echo.Context) error {
		context := context.Background()

		var contactInfo entity.ContactInfo

		err := repo.Preload(context, &contactInfo, "user")
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		err = repo.Find(context, &contactInfo, where.Eq("id", contactInfoID))
		if err != nil {
			fmt.Printf("Deu erro: %v", err.Error())
			return nil
		}

		return ctx.JSON(http.StatusOK, contactInfo)
	}
}
