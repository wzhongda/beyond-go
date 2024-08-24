package logic

import (
	"beyond-go/application/article/api/internal/code"
	"beyond-go/application/article/rpc/service"
	"beyond-go/pkg/xcode"
	"context"
	"encoding/json"

	"beyond-go/application/article/api/internal/svc"
	"beyond-go/application/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const minContentLen = 80

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishRequest) (resp *types.PublishResponse, err error) {
	if len(req.Title) == 0 {
		return nil, code.ArticleTitleEmpty
	}
	if len(req.Content) < minContentLen {
		return nil, code.ArticleContentTooFewWords
	}
	if len(req.Cover) == 0 {
		return nil, code.ArticleCoverEmpty
	}
	userId, err := l.ctx.Value("UserId").(json.Number).Int64()
	if err != nil {
		logx.Errorf("l.ctx.value err:%v", err)
		return nil, xcode.NoLogin
	}
	pret, err := l.svcCtx.ArticleRPC.Publish(l.ctx, &service.PublishRequest{
		UserId:      userId,
		Title:       req.Title,
		Content:     req.Content,
		Description: req.Description,
		Cover:       req.Cover,
	})
	if err != nil {
		logx.Errorf("publish req:%v userid:%d error:%v", req, userId, err)
		return nil, err
	}

	return &types.PublishResponse{
		ArticleId: pret.ArticleId,
	}, nil
}
