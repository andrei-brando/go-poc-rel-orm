package entity

import "time"

type Author struct {
	ID        uint32 `db:"author_id,primary"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Active    bool
	// Book      []Book `ref:"author_id" fk:"author_id" `
}
