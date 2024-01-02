package record_server

import (
	"APICallerStats/app/grpc/auth_help"
	pd "APICallerStats/app/grpc/record_server/recordserver"
	"APICallerStats/app/serve/controller"
	"APICallerStats/common/record_dao_interface"
	"APICallerStats/model"
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"net"
	"time"
)

type GRPCRecordServer struct {
	pd.UnimplementedRecordServerServer
	Dao record_dao_interface.ACSRecordDao
	Env *model.Env
}

func (s *GRPCRecordServer) Record(ctx context.Context, in *pd.RecordRequest) (*pd.RecordResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if !auth_help.CheckAuth(md, s.Env) {
		return &pd.RecordResponse{
			Code: 100,
			Msg:  "Access Denied!",
		}, nil
	}

	p, ok := peer.FromContext(ctx)
	if !ok {
		return &pd.RecordResponse{
			Code: 100,
			Msg:  "Failed to get peer from context",
		}, nil
	}
	ip := in.Ip
	if ip == "" {
		tempIp, _, err := net.SplitHostPort(p.Addr.String())
		if err != nil {
			return &pd.RecordResponse{
				Code: 100,
				Msg:  err.Error(),
			}, nil
		}
		ip = tempIp
	}
	controller.AddACSRecord(
		&model.BaseRecordModel{
			GroupId: int(in.GetGroupId()),
			Name:    in.GetName(),
			Url:     in.GetUrl(),
			Method:  in.GetMethod(),
			IP:      ip,
			AddTime: time.Now(),
		},
	)
	return &pd.RecordResponse{
		Code: 200,
		Msg:  "Success!",
	}, nil
}
