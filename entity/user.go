package entity

import "time"

type User struct {
	ID           uint32 `db:"id,primary"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string
	Active       bool
	ContactInfos []ContactInfo `ref:"id" fk:"user_id" autoload:"true"`
}
