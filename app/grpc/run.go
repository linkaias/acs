package grpc

import (
	"APICallerStats/app/grpc/auth_token_serve"
	pd3 "APICallerStats/app/grpc/auth_token_serve/auth_token"
	"APICallerStats/app/grpc/record_group_server"
	pd2 "APICallerStats/app/grpc/record_group_server/recordgroupserver"
	"APICallerStats/app/grpc/record_server"
	pd "APICallerStats/app/grpc/record_server/recordserver"
	"APICallerStats/common/record_dao_interface"
	"APICallerStats/common/record_group_dao_interface"
	"APICallerStats/model"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func RunGRPCServer(
	env *model.Env, recordDao record_dao_interface.ACSRecordDao,
	recordGroupDao record_group_dao_interface.ACSRecordGroupDao,
) {
	s := grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", env.GRPCServePort))
	if err != nil {
		panic(err)
	}
	// auth token
	pd3.RegisterAuthTokenServerServer(s, &auth_token_serve.GRPCAuthTokenServer{Env: env})
	// record
	pd.RegisterRecordServerServer(s, &record_server.GRPCRecordServer{Dao: recordDao, Env: env})
	// record group
	pd2.RegisterRecordGroupServerServer(s, &record_group_server.GRPCRecordGroupServer{Dao: recordGroupDao, Env: env})

	fmt.Println("GRPC Server is running on port " + env.GRPCServePort)
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

}
