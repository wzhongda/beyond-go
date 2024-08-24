package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"beyond-go/application/article/rpc/internal/svc"
	"beyond-go/application/article/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleDetailLogic {
	return &ArticleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticleDetailLogic) ArticleDetail(in *service.ArticleDetailRequest) (*service.ArticleDetailResponse, error) {
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, uint64(in.ArticleId))
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return &service.ArticleDetailResponse{}, nil
		}
		return nil, err
	}
	return &service.ArticleDetailResponse{
		Article: &service.ArticleItem{
			Id:          int64(article.Id),
			Title:       article.Title,
			Content:     article.Content,
			Description: article.Description,
			Cover:       article.Cover,
			LikeCount:   article.LikeNum,
			PublishTime: article.PublishTime.Unix(),
		},
	}, nil
}
