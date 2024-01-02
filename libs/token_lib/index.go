package token_lib

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

// CreateAccessToken 生成JWT token
func CreateAccessToken(user, secret string, expiry int, source string) (string, int64) {
	exp := time.Now().Unix() + int64(expiry)
	claims := jwt.MapClaims{
		//"exp":    exp,
		"created_at":  fmt.Sprintf("%d", time.Now().Unix()),
		"user":        user,
		"exp_time_at": expiry,
		"source":      source,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(secret))
	return t, exp
}

func CheckAuthToken(source, oldToken, secret string, expire int) error {
	const BearerSchema string = "Bearer "
	if len(oldToken) > 7 && oldToken[0:7] == BearerSchema {
		oldToken = oldToken[7:]
	}
	jwtToken, err := jwt.Parse(
		oldToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				//nil secret key
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		},
	)
	if err != nil {
		return err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return err
	}
	oldSource := fmt.Sprintf("%v", claims["source"])
	iat, _ := strconv.ParseInt(fmt.Sprintf("%v", claims["created_at"]), 10, 64)
	exp, _ := strconv.Atoi(fmt.Sprintf("%v", claims["exp_time_at"]))
	if oldSource == "" || oldSource != source {
		return errors.New("token is invalid! ")
	}
	// 永不过期
	if expire == 0 && exp == 0 {
		return nil
	}
	// 判断配置的token 是否过期 修改过期时间后，作废旧的token
	if iat+int64(exp) > iat+int64(expire) {
		return errors.New("token is expired! ")
	}
	// 判断是否过期
	if iat+int64(exp) < time.Now().Unix() {
		return errors.New("token is expired! ")
	}

	return nil
}
