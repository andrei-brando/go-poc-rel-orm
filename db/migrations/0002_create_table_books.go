// 0002_create_table_books.go

package migrations

import (
	"github.com/go-rel/rel"
)

func MigrateCreateTableBooks(schema *rel.Schema) {
	schema.CreateTable("books", func(t *rel.Table) {
		t.ID("book_id")
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.String("title")
		t.String("category")
		t.Float("price")
		t.Bool("discount")
		t.Int("stock")
		t.Int("author_id", rel.Unsigned(true))
		t.ForeignKey("author_id", "authors", "author_id")
		t.String("publisher")
	})

}

func RollbackCreateTableBooks(schema *rel.Schema) {
	schema.DropTable("books")
}
