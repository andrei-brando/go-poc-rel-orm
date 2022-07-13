package migrations

import (
	"github.com/go-rel/rel"
)

func MigrateCreateTableContactInfos(schema *rel.Schema) {
	schema.CreateTable("contact_infos", func(t *rel.Table) {
		t.ID("id")
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.String("email")
		t.String("phone")
		t.Int("user_id", rel.Unsigned(true))
		t.ForeignKey("user_id", "users", "id")
	})

}

func RollbackCreateTableContactInfos(schema *rel.Schema) {
	schema.DropTable("contact_infos")
}
