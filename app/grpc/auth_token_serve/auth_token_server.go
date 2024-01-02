package auth_token_serve

import (
	pd "APICallerStats/app/grpc/auth_token_serve/auth_token"
	"APICallerStats/libs/token_lib"
	"APICallerStats/model"
	"context"
)

type GRPCAuthTokenServer struct {
	pd.UnimplementedAuthTokenServerServer
	Env *model.Env
}

func (s *GRPCAuthTokenServer) GetToken(ctx context.Context, in *pd.GetTokenRequest) (*pd.GetTokenResponse, error) {
	user, pwd := in.GetUser(), in.GetPassword()
	if s.Env.GRPCUser != user || s.Env.GRPCPassword != pwd {
		return &pd.GetTokenResponse{Code: 100, Msg: "验证失败！"}, nil
	}
	// 生成token
	token, exp := token_lib.CreateAccessToken(user, s.Env.SecretKey, s.Env.JWTExpireTime, "grpc")
	return &pd.GetTokenResponse{
		Code: 200, Msg: "Success！", Data: &pd.TokenInfoModel{Token: token, Exp: int32(exp)},
	}, nil
}
