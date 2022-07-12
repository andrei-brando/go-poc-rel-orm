// 0001_create_authors.go

package migrations

import (
	"context"
	"poc-orm/entity"

	"github.com/go-rel/rel"
)

// MigrateCreateAuthors definition
func MigrateCreateAuthors(schema *rel.Schema) {
	schema.CreateTable("authors", func(t *rel.Table) {
		t.ID("author_id")
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

// RollbackCreateAuthors definition
func RollbackCreateAuthors(schema *rel.Schema) {
	schema.DropTable("authors")
}
