package logic

import (
	"api-share/rpc/internal/svc"
	"api-share/rpc/types/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileIsUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileIsUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileIsUserLogic {
	return &FileIsUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileIsUserLogic) FileIsUser(in *user.FileIsUserReq) (res *user.FileIsUserReply, err error) {
	// todo: add your logic here and delete this line
	res = new(user.FileIsUserReply)
	res.IsUser = true
	if in.FID != 1 {
		res.IsUser = false
	}
	return res, nil
}
