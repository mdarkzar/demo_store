package restapi

import (
	"store/internal/entity/global"
	"store/internal/entity/user"
	"store/internal/transaction"

	"github.com/gin-gonic/gin"
)

func (e *RestAPI) LoadMessages(c *gin.Context) {
	e.ReturnResult(c, func(ts transaction.Session) (gin.H, error) {
		userData := c.MustGet(global.UserObjectKey).(user.User)

		messages, err := e.Usecase.Notification.LoadUserMessages(ts, userData.ID)
		if err != nil {
			return nil, err
		}

		return gin.H{"messageList": messages}, nil
	})
}
