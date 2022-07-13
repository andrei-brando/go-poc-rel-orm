package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-rel/rel"
	"github.com/labstack/echo/v4"
)

// User schema.
type NewUser struct {
	ID        int
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time

	// has many transactions.
	// with custom reference and foreign field declaration.
	// ref: id refers to User.ID field.
	// fk: buyer_id refers to Transaction.BuyerID
	Transactions []Transaction `ref:"id" fk:"buyer_id"`

	// has one address.
	// doesn't contains primary key of other struct.
	// REL can guess the reference and foreign field if it's not specified.
	// autosave tag tells rel to automatically save the association when the parent is inserted/updated/deleted.
	// autoload tag tells rel to automatically load the association when record is queried.
	// alternatively you can use auto to enable both autoload and autosave.
	// Address Address `autosave:"true" autoload:"true"`
}

// Transaction schema.
type Transaction struct {
	ID     int
	Item   string
	Status string

	// belongs to user.
	// contains primary key of other struct.
	Buyer   NewUser `ref:"buyer_id" fk:"id"`
	BuyerID int
}

type Address struct {
	ID   int
	City string

	// belongs to user.
	User   *NewUser
	UserID *int
}

func HandleGet(repo rel.Repository) func(echo.Context) error {
	return func(ctx echo.Context) error {
		context := context.Background()

		errors := make([]string, 0)

		var transaction Transaction

		err := repo.Find(context, &transaction, rel.Select("*", "buyer.*").JoinAssoc("buyer"))
		if err != nil {
			errors = append(errors, err.Error())
		}

		if len(errors) > 0 {
			for i, erro := range errors {
				fmt.Printf("Erro %d: %v", i, erro)
			}
		}

		return ctx.JSON(http.StatusOK, transaction)
	}
}
