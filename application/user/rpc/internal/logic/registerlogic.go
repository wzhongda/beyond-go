package logic

import (
	"beyond-go/application/user/rpc/internal/code"
	"beyond-go/application/user/rpc/model"
	"context"
	"time"

	"beyond-go/application/user/rpc/internal/svc"
	"beyond-go/application/user/rpc/pb/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {
	if len(in.Username) == 0 {
		return nil, code.RegisterNameEmpty
	}

	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:   in.Username,
		Mobile:     in.Mobile,
		Avatar:     in.Avatar,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		logx.Errorf("register req:%v error:%v", in, err)
	}
	userId, err := ret.LastInsertId()
	if err != nil {
		logx.Errorf("lastinterid error:%v", err)
		return nil, err
	}
	return &service.RegisterResponse{UserId: userId}, nil

	return &service.RegisterResponse{}, nil
}
