package entity

import "time"

type Book struct {
	ID        int `db:"book_id,primary"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Category  string
	Price     int
	Discount  bool
	Stock     int
	Author    Author `ref:"author_id" fk:"author_id"`
	AuthorID  uint32
	Publisher string
}
