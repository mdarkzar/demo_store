package restapi

import (
	"net/http"
	"store/internal/entity/global"
	"store/internal/transaction"

	"github.com/gin-gonic/gin"
)

func (e *RestAPI) Auth(c *gin.Context) {
	var form struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.ShouldBind(&form); err != nil {
		e.errorResponse(c, global.ErrParamsIncorrect)
		return
	}

	e.ReturnResult(c, func(ts transaction.Session) (gin.H, error) {
		jwtToken, err := e.Usecase.User.Auth(ts, form.Login, form.Password)
		if err != nil {
			return nil, err
		}

		e.setTokenInCookie(c, global.AuthToken, jwtToken)

		return gin.H{"token": jwtToken}, nil
	})

}

func (e *RestAPI) Profile(c *gin.Context) {
	userData := c.MustGet(global.UserObjectKey)

	e.returnResult(c, gin.H{"user": userData})
}

func (e *RestAPI) Logout(c *gin.Context) {
	e.setTokenInCookie(c, global.AuthToken, "")

	c.JSON(http.StatusOK, nil)
}
