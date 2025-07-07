package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("microMemorandum")

type Claims struct {
	Id uint `json:"id"`
	jwt.RegisteredClaims
}

// GenerateToken 签发token
func GenerateToken(id uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{ // 新版本，即将丢弃
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "teacher_hand",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenCliams, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	fmt.Println("tokenCliams", *tokenCliams)

	if tokenCliams != nil {
		if claims, ok := tokenCliams.Claims.(*Claims); ok && tokenCliams.Valid {
			fmt.Println("claims", *claims)
			return claims, nil
		}
	}

	return nil, err
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

// GenerateEmailToken 签发邮箱token
func GenerateEmailToken(id, operation uint, email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := EmailClaims{
		UserID:        id,
		Email:         email,
		Password:      password,
		OperationType: operation,
		StandardClaims: jwt.StandardClaims{ // 老版本，即将丢弃
			ExpiresAt: expireTime.Unix(),
			Issuer:    "teacher_hand",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseEmailToken 验证邮箱token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenEmailCliams, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenEmailCliams != nil {
		if emailClaims, ok := tokenEmailCliams.Claims.(*EmailClaims); ok && tokenEmailCliams.Valid {
			return emailClaims, nil
		}
	}

	return nil, err
}
