package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	tokenEncodeString = []byte(os.Getenv("TOKEN_HASH"))
)

type Claims struct {
	UserID      int       `json:"user_id"`
	Login       string    `json:"login"`
	CreatedDate time.Time `json:"created_date"`
	jwt.StandardClaims
}

// NewJwtToken создание нового jwt токена
func NewJwtToken(userID int, login string, now time.Time) (string, error) {
	// Create the Claims
	claims := Claims{
		userID,
		login,
		now,
		jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(tokenEncodeString)
}

// ParseToken parse token
func ParseToken(unparsedToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(unparsedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenEncodeString, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid

}
