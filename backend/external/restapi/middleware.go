package restapi

import (
	"store/internal/entity/global"
	"store/internal/entity/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func (e *RestAPI) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token    string
			inCookie = false
		)

		// токен проверяется как в header (для http клиентов, мобильных приложений)
		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			// и для браузеров в куки, в связи с особенностями хранения
			if cookieToken, exists := getTokenByCookie(c); exists {
				token = cookieToken
			} else {
				e.errorResponse(c, global.ErrNeedAuth)
				return
			}
		}

		if !inCookie && strings.HasPrefix(bearerToken, "Bearer") {
			token = strings.Replace(bearerToken, "Bearer ", "", -1)
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			e.errorResponse(c, global.ErrNeedAuth)
			return
		}

		userData, err := e.Usecase.User.FindUser(claims.UserID)
		if err != nil {
			e.errorResponse(c, err)
			return
		}

		c.Set(global.UserObjectKey, userData)
		c.Next()
	}
}

func getTokenByCookie(c *gin.Context) (token string, exists bool) {
	cookie, err := c.Request.Cookie(global.AuthToken)
	if err != nil {
		return
	}
	token = cookie.Value
	exists = token != ""
	return
}
