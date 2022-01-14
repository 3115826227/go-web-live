package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

var (
	randSource = rand.NewSource(time.Now().UnixNano())
)

//生成八位数字
func GenerateSerialNumber() string {
	return fmt.Sprintf("1%07v", rand.New(randSource).Int31n(10000000))
}

/*
	根据用户id和创建时间生成jwt Token
*/
func GenerateToken(userID string, createTime time.Time, tokenSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     userID,
		"create_time": createTime,
	})

	return token.SignedString([]byte(tokenSecret))
}
