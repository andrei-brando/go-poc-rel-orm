package migrations

import (
	"github.com/go-rel/rel"
)

func MigrateCreateTableNewUsers(schema *rel.Schema) {
	schema.CreateTable("new_users", func(t *rel.Table) {
		t.ID("id")
		t.String("name")
		t.Int("age")
		t.DateTime("created_at")
		t.DateTime("updated_at")
	})

}

func RollbackCreateTableNewUsers(schema *rel.Schema) {
	schema.DropTable("new_users")
}
