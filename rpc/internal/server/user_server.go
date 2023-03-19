// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"api-share/rpc/internal/logic"
	"api-share/rpc/internal/svc"
	"api-share/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) FileIsUser(ctx context.Context, in *user.FileIsUserReq) (*user.FileIsUserReply, error) {
	l := logic.NewFileIsUserLogic(ctx, s.svcCtx)
	return l.FileIsUser(in)
}