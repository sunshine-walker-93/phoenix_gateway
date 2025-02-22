package util

import (
	"github.com/golang-jwt/jwt"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/config"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func init() {
	jwtSecret = []byte(config.GetGlobalConfig().AppSetting.JwtSecret)
}

// GenerateToken 生成token
func GenerateToken(name, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.GetGlobalConfig().AppSetting.TokenExpireHour) * time.Hour)

	claims := Claims{
		name,
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析token
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
