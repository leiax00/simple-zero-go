package biz

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	v1 "github.com/simple-zero-go/api/wx/service/v1"
	"github.com/simple-zero-go/app/wx/service/internal/conf"
	_const "github.com/simple-zero-go/app/wx/service/internal/const"
	"github.com/simple-zero-go/pkg/utils"
	"sort"
	"strings"
)

type WxSysUseCase struct {
	biz  *Biz
	log  *log.Helper
	conf *conf.WxConf
	repo WxSysRepo
}

func NewWxSysUseCase(biz *Biz, logger log.Logger, wx *conf.WxConf, repo WxSysRepo) *WxSysUseCase {
	return &WxSysUseCase{
		biz:  biz,
		log:  log.NewHelper(log.With(logger, "module", "useCase/wxSys")),
		conf: wx,
		repo: repo,
	}
}

type WxSysRepo interface {
	SaveToken(ctx context.Context, data *v1.TokenReply) error
}

func (uc *WxSysUseCase) AuthServer(ctx context.Context, data *v1.AuthServerReq) (string, error) {
	tmpArray := []string{data.Nonce, data.Timestamp, uc.conf.ServerToken}
	sort.Strings(tmpArray)
	h := sha1.New()
	h.Write([]byte(strings.Join(tmpArray, "")))
	calcVal := hex.EncodeToString(h.Sum(nil))
	if calcVal == data.Signature {
		return data.Echostr, nil
	}
	return "", nil
}

func (uc *WxSysUseCase) GetAccessToken(ctx context.Context, data *v1.TokenReq) (*v1.TokenReply, error) {
	token, err := uc.biz.rc.Get(ctx, _const.WX_ACCESS_TOKEN).Result()
	if err != nil && err != redis.Nil {
		return &v1.TokenReply{
			AccessToken: token,
		}, nil
	}
	resp, err := uc.biz.hc.R().
		SetContext(ctx).
		SetResult(&v1.TokenReply{}).
		SetQueryParams(map[string]string{
			"grant_type": utils.If(data.GrantType == "", "client_credential", data.GrantType),
			"appid":      utils.If(data.Appid == "", uc.conf.Appid, data.Appid),
			"secret":     utils.If(data.Secret == "", uc.conf.Secret, data.Secret),
		}).Get("https://api.weixin.qq.com/cgi-bin/token")
	if err != nil {
		log.Info(err)
		return nil, err
	}
	reply := resp.Result().(*v1.TokenReply)
	err = uc.repo.SaveToken(ctx, reply)
	return reply, err
}
