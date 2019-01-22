package jwt

import (
	x "github.com/CrocdileChan/common/errhandle"
	lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

type Jwt struct {
	token []byte
}

func NewJwt(token string) *Jwt {
	return &Jwt{
		token: []byte(token),
	}
}

func (j *Jwt) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, nil)
		} else {
			claims, err := j.parseHmac(auth)
			if err != nil || claims == nil {
				x.UnauthCheck(c, err, "")
			}
			c.Request.Header.Set("user_id", claims["user_id"].(string))
			c.Next()
		}

	}
}

func (j *Jwt) parseHmac(auth string) (map[string]interface{}, error) {
	token, err := lib.Parse(auth, func(token *lib.Token) (interface{}, error) {
		if _, ok := token.Method.(*lib.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte("1234567890"), nil
	})
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, errors.New("ParseHmac error")
	}
	if claims, ok := token.Claims.(lib.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Can't parse token")
	}
}
