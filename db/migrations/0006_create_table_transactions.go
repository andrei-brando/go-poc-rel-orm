package migrations

import (
	"github.com/go-rel/rel"
)

func MigrateCreateTableTransactions(schema *rel.Schema) {
	schema.CreateTable("transactions", func(t *rel.Table) {
		t.ID("id")
		t.String("item")
		t.String("status")
		t.Int("buyer_id", rel.Unsigned(true))
		t.ForeignKey("buyer_id", "new_users", "id")
	})

}

func RollbackCreateTableTransactions(schema *rel.Schema) {
	schema.DropTable("transactions")
}
