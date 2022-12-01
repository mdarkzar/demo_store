package product

import "time"

type Product struct {
	ID          int       `json:"id" db:"product_id"`
	Name        string    `json:"name" db:"name"`
	Price       float64   `json:"price" db:"price"`
	CreatorID   int       `json:"creator_id" db:"creator_id"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
}
