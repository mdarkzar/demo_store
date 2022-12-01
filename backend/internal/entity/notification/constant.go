package notification

import "fmt"

const (
	CreateProductTitle = "Создан новый продукт"
	DeleteProductTitle = "Удален продукт"
)

var CreateProductBody = func(login, name string, price float64) string {
	return fmt.Sprintf("Создал: %s\nНазвание %s\nСтоимость: %.2f", login, name, price)
}

var DeleteProductBody = func(productID int, login, name string, price float64) string {
	return fmt.Sprintf("Удален продукт №%d\nУдалил: %s\nНазвание %s\nСтоимость: %.2f", productID, login, name, price)
}
