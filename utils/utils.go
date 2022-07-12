package utils

import (
	"context"

	"github.com/labstack/echo/v4"
)

func GetContext(ctx echo.Context) context.Context {
	ct := ctx.Get("domain")
	return ct.(context.Context)
}
