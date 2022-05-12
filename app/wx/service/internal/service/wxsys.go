package service

import (
	"context"
	"encoding/xml"
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

func (s *WxSysService) DispatchMsg(ctx context.Context, req *wxPb.MsgReq) (*wxPb.StringReply, error) {
	rst, err := s.uc.DispatchMsg(ctx, req)
	if err != nil {
		return &wxPb.StringReply{Msg: ""}, err
	}
	bytes, err := xml.Marshal(rst)
	s.log.Info(string(bytes))
	return &wxPb.StringReply{Msg: string(bytes)}, err
}

func (s *WxSysService) GetAccessToken(ctx context.Context, tokenObj *wxPb.TokenReq) (*wxPb.TokenReply, error) {
	return s.uc.GetAccessToken(ctx, tokenObj)
}

func (s *WxSysService) CreateMenu(ctx context.Context, menu *wxPb.Menu) (*wxPb.CommonReply, error) {
	return s.uc.CreateMenu(ctx, menu)
}
