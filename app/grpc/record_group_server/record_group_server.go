package record_group_server

import (
	"APICallerStats/app/grpc/auth_help"
	pd "APICallerStats/app/grpc/record_group_server/recordgroupserver"
	"APICallerStats/common/record_group_dao_interface"
	"APICallerStats/model"
	"context"
	"google.golang.org/grpc/metadata"
)

type GRPCRecordGroupServer struct {
	pd.RecordGroupServerServer
	Dao record_group_dao_interface.ACSRecordGroupDao
	Env *model.Env
}

func (s *GRPCRecordGroupServer) RecordGroupList(ctx context.Context, in *pd.RecordGroupListRequest) (
	*pd.RecordGroupListResponse, error,
) {
	md, _ := metadata.FromIncomingContext(ctx)
	if !auth_help.CheckAuth(md, s.Env) {
		return &pd.RecordGroupListResponse{
			Code: 100,
			Msg:  "Access Denied!",
		}, nil
	}

	rt := &pd.RecordGroupListResponse{}
	list, err := s.Dao.List()
	if err != nil {
		rt.Code = 100
		rt.Msg = err.Error()
		return rt, nil
	}
	info := make([]*pd.RecordGroupModel, 0)
	for _, val := range list {
		info = append(
			info, &pd.RecordGroupModel{
				Id:        int32(val.Id),
				Name:      val.Name,
				Desc:      val.Desc,
				CreatedAt: val.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdateAt:  val.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		)
	}
	rt.Code = 200
	rt.Msg = "Success!"
	rt.Data = info
	return rt, nil

}
func (s *GRPCRecordGroupServer) RecordGroupUpdate(ctx context.Context, in *pd.RecordGroupUpRequest) (
	*pd.RecordGroupResponse, error,
) {
	md, _ := metadata.FromIncomingContext(ctx)
	if !auth_help.CheckAuth(md, s.Env) {
		return &pd.RecordGroupResponse{
			Code: 100,
			Msg:  "Access Denied!",
		}, nil
	}
	rt := &pd.RecordGroupResponse{}
	if in.GetId() <= 0 {
		rt.Code = 100
		rt.Msg = "参数错误"
		return rt, nil
	}
	if in.GetId() == 1 {
		rt.Code = 100
		rt.Msg = "默认分组不支持修改"
		return rt, nil
	}

	err := s.Dao.UpByGroupId(
		&model.BaseRecordGroupModel{
			Id:   int(in.GetId()),
			Name: in.GetName(),
			Desc: in.GetDesc(),
		},
	)
	if err != nil {
		rt.Code = 100
		rt.Msg = err.Error()
		return rt, nil
	}
	rt.Code = 200
	rt.Msg = "Success!"
	return rt, nil
}
func (s *GRPCRecordGroupServer) RecordGroupAdd(ctx context.Context, in *pd.RecordGroupAddRequest) (
	*pd.RecordGroupResponse, error,
) {
	md, _ := metadata.FromIncomingContext(ctx)
	if !auth_help.CheckAuth(md, s.Env) {
		return &pd.RecordGroupResponse{
			Code: 100,
			Msg:  "Access Denied!",
		}, nil
	}
	rt := &pd.RecordGroupResponse{}
	err := s.Dao.Add(
		&model.BaseRecordGroupModel{
			Name: in.GetName(),
			Desc: in.GetDesc(),
		},
	)
	if err != nil {
		rt.Code = 100
		rt.Msg = err.Error()
		return rt, nil
	}
	rt.Code = 200
	rt.Msg = "Success!"
	return rt, nil

}
func (s *GRPCRecordGroupServer) RecordGroupDel(ctx context.Context, in *pd.RecordGroupDelRequest) (
	*pd.RecordGroupResponse, error,
) {
	md, _ := metadata.FromIncomingContext(ctx)
	if !auth_help.CheckAuth(md, s.Env) {
		return &pd.RecordGroupResponse{
			Code: 100,
			Msg:  "Access Denied!",
		}, nil
	}
	rt := &pd.RecordGroupResponse{}
	if in.GetId() <= 0 {
		rt.Code = 100
		rt.Msg = "参数错误"
		return rt, nil
	}
	if in.GetId() == 1 {
		rt.Code = 100
		rt.Msg = "默认分组不支持修改"
		return rt, nil
	}
	err := s.Dao.DelByGroupId(int(in.GetId()))
	if err != nil {
		rt.Code = 100
		rt.Msg = err.Error()
		return rt, nil
	}
	rt.Code = 200
	rt.Msg = "Success!"
	return rt, nil
}
