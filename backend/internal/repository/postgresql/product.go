package postgresql

import (
	"store/internal/entity/product"
	"store/internal/repository"
	"store/internal/transaction"
	"store/tools/gensql"
)

type productRepo struct{}

func NewProduct() repository.Product {
	return &productRepo{}
}

func (r *productRepo) Create(ts transaction.Session, userID int, name string, price float64) (productID int, err error) {
	return gensql.Get[int](SqlxTx(ts), `INSERT INTO product (name, price, creator_id) VALUES ($1, $2, $3) returning product_id`, name, price, userID)
}

func (r *productRepo) Remove(ts transaction.Session, productID int) error {
	_, err := SqlxTx(ts).Exec(`UPDATE product SET deleted_date = now() where product_id = $1`, productID)
	return err
}

func (r *productRepo) FindByID(ts transaction.Session, productID int) (product.Product, error) {
	sqlQuery := `
	select p.product_id, p.name, p.price, p.creator_id, p.created_date
	from product p
	where p.deleted_date is null
	and p.product_id = $1
	order by p.created_date
	`

	return gensql.Get[product.Product](SqlxTx(ts), sqlQuery, productID)
}

func (r *productRepo) LoadAll(ts transaction.Session) ([]product.Product, error) {
	sqlQuery := `
	select p.product_id, p.name, p.price, p.creator_id, p.created_date
	from product p
	where p.deleted_date is null
	order by p.created_date
	`

	return gensql.Select[product.Product](SqlxTx(ts), sqlQuery)
}
