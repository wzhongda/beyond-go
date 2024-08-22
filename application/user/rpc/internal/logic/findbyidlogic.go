package logic

import (
	"context"

	"beyond-go/application/user/rpc/internal/svc"
	"beyond-go/application/user/rpc/pb/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *service.FindByIdRequest) (*service.FindByIdResponse, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.UserId))
	if err != nil {
		logx.Errorf("findby id %d error:%v", in.UserId, err)
		return nil, err
	}

	return &service.FindByIdResponse{
		UserId:   int64(user.Id),
		Username: user.Username,
		Avatar:   user.Avatar,
	}, nil
}
