package migrations

import (
	"github.com/go-rel/rel"
)

func MigrateCreateTableAddresses(schema *rel.Schema) {
	schema.CreateTable("addresses", func(t *rel.Table) {
		t.ID("id")
		t.String("city")
		t.Int("user_id", rel.Unsigned(true))
		t.ForeignKey("user_id", "new_users", "id")
	})

}

func RollbackCreateTableAddresses(schema *rel.Schema) {
	schema.DropTable("addresses")
}
