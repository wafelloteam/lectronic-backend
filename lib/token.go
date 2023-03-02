package lib

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var Secret = []byte(os.Getenv("JWT_KEY"))

type claims struct {
	UserID string
	Role   string
	jwt.StandardClaims
}

func NewToken(uuid string, role string) *claims {
	return &claims{
		UserID: uuid,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
}

func (c *claims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

func CheckToken(userToken string) (*claims, error) {
	token, err := jwt.ParseWithClaims(userToken, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*claims)
	return claims, nil

}
