package util

import (
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

func GenerateToken(uid int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "unison",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func CheckClaims(token string) (*Claims, int) {
	code := e.SUCCESS
	claims, err := ParseToken(token)
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		default:
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
	}

	return claims, code
}

func ParseUidFromContext(c *gin.Context) (int64, int) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		return 0, e.ERROR_AUTH_CHECK_TOKEN_FAIL
	}
	claims, code := CheckClaims(token)
	if code != e.SUCCESS {
		return 0, code
	}

	if claims.Uid == 0 {
		return 0, e.ERROR_AUTH_CHECK_TOKEN_FAIL
	}

	return claims.Uid, e.SUCCESS
}
