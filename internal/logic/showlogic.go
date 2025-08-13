package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	Err404 = errors.New("404")
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// 查看短连接
	// 布隆过滤器
	// - 基于内存，缺点 服务重启后重新加载
	// - 基于Redis
	exist, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err != nil {
		logx.Errorw("Bloom Filter failed", logx.LogField{
			Value: err.Error(),
			Key:   "err",
		})
	}
	// 不存在的短链接
	if !exist {
		return nil, Err404
	}

	fmt.Println("check cache and DB ...")
	u, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{Valid: true, String: req.ShortUrl})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, Err404
		}
		logx.Errorw("Show failed", logx.LogField{Value: err.Error(), Key: "err"})
		return nil, err
	}
	// 查询到， 调用handler处理
	return &types.ShowResponse{LongUrl: u.Lurl.String}, nil
}
