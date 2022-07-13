package migrations

import (
	"context"
	"poc-orm/entity"

	"github.com/go-rel/rel"
)

func MigrateCreateTableUsers(schema *rel.Schema) {
	schema.CreateTable("users", func(t *rel.Table) {
		t.ID("id")
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.String("name")
		t.Bool("active")
	})

	schema.Do(func(ctx context.Context, repo rel.Repository) error {
		// add seeds
		return repo.Insert(ctx, &entity.Author{Name: "Autor 1"})
	})
}

func RollbackCreateTableUsers(schema *rel.Schema) {
	schema.DropTable("users")
}
