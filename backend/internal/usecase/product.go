package usecase

import (
	"fmt"
	"store/bimport"
	"store/internal/entity/global"
	"store/internal/entity/notification"
	"store/internal/entity/product"
	"store/internal/rimport"
	"store/internal/transaction"

	"github.com/sirupsen/logrus"
)

type ProductUsecase struct {
	log *logrus.Logger

	rimport.RepositoryImports
	*bimport.BridgeImports
}

func NewProductUsecase(
	log *logrus.Logger,
	ri rimport.RepositoryImports,
	bi *bimport.BridgeImports,
) *ProductUsecase {
	return &ProductUsecase{
		log:               log,
		RepositoryImports: ri,
		BridgeImports:     bi,
	}
}

// Create создание товара
func (u *ProductUsecase) Create(ts transaction.Session, userID int, name string, price float64, stID int) (id int, err error) {
	lf := logrus.Fields{
		"userID": userID,
		"name":   name,
		"price":  price,
		"st_id":  stID,
	}

	userData, err := u.Repository.User.FindByID(ts, userID)
	if err != nil {
		u.log.WithFields(lf).Errorln(
			fmt.Sprintf("не удалось найти пользователя; ошибка: %v", err),
		)
		return 0, global.ErrInternalError
	}

	id, err = u.Repository.Product.Create(ts, userID, name, price, stID)
	if err != nil {
		u.log.WithFields(lf).Errorln(
			fmt.Sprintf("не удалось создать продукт; ошибка: %v", err),
		)
		err = global.ErrInternalError
		return
	}

	u.Bridge.Notification.SendAll(ts, notification.CreateProductTitle, notification.CreateProductBody(userData.Login, name, price))

	u.log.WithFields(lf).Infoln("создан продукт №", id)

	return

}

// Remove удаление товара
func (u *ProductUsecase) Remove(ts transaction.Session, userID int, productID int) error {
	lf := logrus.Fields{
		"userID":    userID,
		"productID": productID,
	}

	productData, err := u.Repository.Product.FindByID(ts, productID)
	if err != nil {
		if err == global.ErrNoData {
			return fmt.Errorf("продукт не найден")
		}

		u.log.WithFields(lf).Errorln(
			fmt.Sprintf("не удалось найти продукт; ошибка: %v", err),
		)
		return global.ErrInternalError
	}

	userData, err := u.Repository.User.FindByID(ts, userID)
	if err != nil {
		u.log.WithFields(lf).Errorln(
			fmt.Sprintf("не удалось найти пользователя; ошибка: %v", err),
		)
		return global.ErrInternalError
	}

	if err := u.Repository.Product.Remove(ts, productID); err != nil {
		u.log.WithFields(lf).Errorln(
			fmt.Sprintf("не удалось удалить продукт; ошибка: %v", err),
		)
		return global.ErrInternalError
	}

	u.Bridge.Notification.SendAll(ts, notification.DeleteProductTitle, notification.DeleteProductBody(productID, userData.Login, productData.Name, productData.Price))

	u.log.WithFields(lf).Infoln("удален продукт №", productID)

	return nil
}

// FindByID найти по id
func (u *ProductUsecase) FindByID(ts transaction.Session, productID int) (product.Product, error) {
	lf := logrus.Fields{
		"productID": productID,
	}

	p, err := u.Repository.Product.FindByID(ts, productID)
	switch err {
	case global.ErrNoData:
		return p, global.ErrNoData
	case nil:
		return p, nil
	default:
		u.log.WithFields(lf).Errorln(
			fmt.Sprintf("не удалось найти продукт; ошибка: %v", err),
		)
		return p, global.ErrInternalError
	}
}

// LoadAll загрузить все товары
func (u *ProductUsecase) LoadAll(ts transaction.Session) ([]product.Product, error) {
	data, err := u.Repository.Product.LoadAll(ts)
	switch err {
	case global.ErrNoData:
		return nil, global.ErrNoData
	case nil:
		return data, nil
	default:
		u.log.Errorln(
			fmt.Sprintf("не удалось найти продукты; ошибка: %v", err),
		)
		return data, global.ErrInternalError
	}

}
