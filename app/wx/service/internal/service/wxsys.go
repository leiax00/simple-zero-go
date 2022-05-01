package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	wxPb "github.com/simple-zero-go/api/wx/service/v1"
	"github.com/simple-zero-go/app/wx/service/internal/biz"
)

type WxSysService struct {
	wxPb.UnimplementedWxSysServer

	uc  *biz.WxSysUseCase
	log *log.Helper
}

func NewWxSysService(uc *biz.WxSysUseCase, logger log.Logger) *WxSysService {
	return &WxSysService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/wxSys")),
	}
}

func (s *WxSysService) AuthServer(ctx context.Context, req *wxPb.AuthServerReq) (*wxPb.AuthServerResp, error) {
	str, err := s.uc.AuthServer(ctx, req)
	return &wxPb.AuthServerResp{Echostr: str}, err
}

func (s *WxSysService) GetAccessToken(ctx context.Context, tokenObj *wxPb.TokenReq) (*wxPb.TokenReply, error) {
	return s.uc.GetAccessToken(ctx, tokenObj)
}

func (s *WxSysService) CreateMenu(ctx context.Context, menu *wxPb.Menu) (*wxPb.CommonReply, error) {
	return s.uc.CreateMenu(ctx, menu)
}
