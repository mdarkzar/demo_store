package restapi

import (
	"store/internal/entity/global"
	"store/internal/entity/user"
	"store/internal/transaction"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (e *RestAPI) CreateProduct(c *gin.Context) {
	var form struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	if err := c.ShouldBind(&form); err != nil {
		e.errorResponse(c, global.ErrParamsIncorrect)
		return
	}

	e.ReturnResult(c, func(ts transaction.Session) (gin.H, error) {
		userData := c.MustGet(global.UserObjectKey).(user.User)

		productID, err := e.Usecase.Product.Create(ts, userData.ID, form.Name, form.Price)
		if err != nil {
			return nil, err
		}

		return gin.H{"product_id": productID}, nil
	})
}

func (e *RestAPI) RemoveProduct(c *gin.Context) {
	var form struct {
		ProductID int `json:"product_id"`
	}

	if err := c.ShouldBind(&form); err != nil {
		e.errorResponse(c, global.ErrParamsIncorrect)
		return
	}

	e.ReturnSuccessNull(c, func(ts transaction.Session) error {
		userData := c.MustGet(global.UserObjectKey).(user.User)

		err := e.Usecase.Product.Remove(ts, userData.ID, form.ProductID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (e *RestAPI) FindProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.errorResponse(c, global.ErrParamsIncorrect)
		return
	}

	e.ReturnResultWithoutCommit(c, func(ts transaction.Session) (gin.H, error) {
		productData, err := e.Usecase.Product.FindByID(ts, productID)
		if err != nil {
			return nil, err
		}

		return gin.H{"product": productData}, nil
	})
}

func (e *RestAPI) LoadAllProduct(c *gin.Context) {
	e.ReturnResultWithoutCommit(c, func(ts transaction.Session) (gin.H, error) {
		productData, err := e.Usecase.Product.LoadAll(ts)
		if err != nil {
			return nil, err
		}

		return gin.H{"productList": productData}, nil
	})
}
