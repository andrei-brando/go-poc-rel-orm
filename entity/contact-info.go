package entity

import "time"

type ContactInfo struct {
	ID        uint32 `db:"id,primary"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Phone     string
	UserID    uint32
	User      User `ref:"user_id" fk:"id"`
}
