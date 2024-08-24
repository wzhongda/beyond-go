package logic

import (
	"beyond-go/application/article/rpc/internal/code"
	"beyond-go/application/article/rpc/internal/types"
	"beyond-go/pkg/xcode"
	"context"

	"beyond-go/application/article/rpc/internal/svc"
	"beyond-go/application/article/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleDeleteLogic {
	return &ArticleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleDeleteLogic) ArticleDelete(in *service.ArticleDeleteRequest) (*service.ArticleDeleteResponse, error) {
	if in.UserId <= 0 {
		return nil, code.UserIdInvalid
	}
	if in.ArticleId <= 0 {
		return nil, code.ArticleIdInvalid
	}
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, uint64(in.ArticleId))
	if err != nil {
		l.Logger.Errorf("article findone req:%v error:%v", in, err)
		return nil, err
	}
	if int64(article.AuthorId) != in.UserId {
		return nil, xcode.AccessDenied
	}
	err = l.svcCtx.ArticleModel.UpdateArticleStatus(l.ctx, in.ArticleId, types.ArticleStatusDelete)
	if err != nil {
		l.Logger.Errorf("updatestatus req:%v,error:%v", in, err)
		return nil, err
	}

	_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, articlesKey(in.UserId, types.SortPublishTime), in.ArticleId)
	if err != nil {
		l.Logger.Errorf("zremctx req:%v,errror:%v", in, err)
	}
	_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, articlesKey(in.UserId, types.SortLikeCount), in.ArticleId)
	if err != nil {
		l.Logger.Errorf("zremctx req:%v error:%v", in, err)
	}
	return &service.ArticleDeleteResponse{}, nil

}
