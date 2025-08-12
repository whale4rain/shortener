package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"shortener/internal/svc"
	"shortener/internal/types"
	"shortener/pkg/connect"
	"shortener/pkg/md5"
	"shortener/pkg/urltool"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1.校验数据
	// 1.1 参数合法 @./handler/converthandler.go
	// 1.2 判断url是否可访问
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("无效连接")
	}
	// 1.3 是否转链过
	md5Value := md5.Sum([]byte(req.LongUrl))

	u, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{
		String: md5Value,
		Valid:  true,
	})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, fmt.Errorf("该链接已转为%s", u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5", logx.LogField{
			Key:   "error",
			Value: err.Error(),
		})
		return nil, err
	}
	// 1.4 输入不能是短链接(循环转链)
	basePath, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urltool.GetBasePath failed", logx.LogField{
			Key:   "lurl",
			Value: req.LongUrl,
		}, logx.LogField{
			Key:   "error",
			Value: err.Error(),
		})
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{
		String: basePath,
		Valid:  true,
	})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, errors.New("该链接已经是短链了")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl", logx.LogField{
			Key:   "error",
			Value: err.Error(),
		})
		return nil, err
	}
	return

}
