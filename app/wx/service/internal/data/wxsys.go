package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/simple-zero-go/api/wx/service/v1"
	"github.com/simple-zero-go/app/wx/service/internal/biz"
	_const "github.com/simple-zero-go/app/wx/service/internal/const"
	"time"
)

type WxSysRepo struct {
	data *Data
	log  *log.Helper
}

func NewWxSysRepo(data *Data, logger log.Logger) biz.WxSysRepo {
	return &WxSysRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/wxSys")),
	}
}

func (wp WxSysRepo) SaveToken(ctx context.Context, data *v1.TokenReply) error {
	return wp.data.rc.
		Set(ctx, _const.WX_ACCESS_TOKEN, data.AccessToken, time.Duration(7200)*time.Second).Err()
}
