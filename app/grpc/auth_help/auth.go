package auth_help

import (
	"APICallerStats/libs/token_lib"
	"APICallerStats/model"
	"google.golang.org/grpc/metadata"
	"strings"
)

func CheckAuth(md metadata.MD, env *model.Env) bool {
	authHeader := md["authorization"]
	if len(authHeader) <= 0 {
		return false
	}
	tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")
	err := token_lib.CheckAuthToken("grpc", tokenStr, env.SecretKey, env.JWTExpireTime)
	return err == nil
}
