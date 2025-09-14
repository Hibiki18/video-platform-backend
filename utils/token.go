package utils

import (
	"fmt"
	"strings"
	"time"
	"video-platform-backend/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() []byte {
    return []byte(config.AppConfig.JWT.Secret)
}

func GenerateJWT(username string) (string, error) {
	token_lifespan := config.AppConfig.JWT.ExpireDuration
	if token_lifespan <= 0 {
        token_lifespan = 72 // 默认 72 小时
    }
		
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

// 从请求头中获取token
func ExtractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ParseToken(c *gin.Context) (error) {
	tokenString := ExtractToken(c)
	fmt.Println(tokenString)

	parser := jwt.NewParser(jwt.WithLeeway(30 * time.Second))

	_, err := parser.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return getJWTSecret(), nil
	})
	if err != nil {
		return err
	}

	return nil
}

// 从jwt中解析出username
func ExtractUsername(c *gin.Context) (string, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return getJWTSecret(), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	// 如果jwt有效
	if ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	}
	return "", fmt.Errorf("invalid token")
}