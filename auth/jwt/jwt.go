package jwt

import (
	x "github.com/CrocdileChan/common/errhandle"
	lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Jwt struct {
	key []byte
}

type jwtCustomClaims struct {
	lib.StandardClaims

	uid int
}

func NewJwt(key string) *Jwt {
	return &Jwt{
		key: []byte(key),
	}
}

func (j *Jwt) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if c.Request.Method == http.MethodOptions {
			c.JSON(http.StatusOK, nil)
		} else {
			claims, err := j.ParseToken(auth)
			if err != nil || claims == nil {
				x.UnauthCheck(c, err, "")
			}
			c.Request.Header.Set("user_id", strconv.Itoa(claims.(jwtCustomClaims).uid))
			c.Next()
		}

	}
}

func (j *Jwt) CreateToken(uid int) (string, error) {
	claims := jwtCustomClaims{
		StandardClaims: lib.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
		},
		uid: uid,
	}
	token := lib.NewWithClaims(lib.SigningMethodHS256, claims)
	return token.SignedString(j.key)
}

func (j *Jwt) ParseToken(auth string) (claims lib.Claims, err error) {
	var token *lib.Token
	token, err = lib.Parse(auth, func(*lib.Token) (interface{}, error) {
		return j.key, nil
	})
	claims = token.Claims
	return
}

/*
func (j *Jwt) ParseToken(auth string) (map[string]interface{}, error) {
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
*/
