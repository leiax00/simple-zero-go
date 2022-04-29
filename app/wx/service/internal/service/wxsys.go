package service

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	wxPb "github.com/simple-zero-go/api/wx/service/v1"
	"sort"
	"strings"
)

type WxSysService struct {
	wxPb.UnimplementedWxSysServer
}

func NewWxSysService() *WxSysService {
	return &WxSysService{}
}

func (s *WxSysService) AuthServer(ctx context.Context, req *wxPb.AuthServerReq) (*wxPb.AuthServerResp, error) {
	token := "lax4832"
	tmpArray := []string{req.Nonce, req.Timestamp, token}
	sort.Strings(tmpArray)
	h := sha1.New()
	h.Write([]byte(strings.Join(tmpArray, "")))
	calcVal := hex.EncodeToString(h.Sum(nil))
	if calcVal == req.Signature {
		return &wxPb.AuthServerResp{Echostr: req.Echostr}, nil
	}
	return &wxPb.AuthServerResp{Echostr: ""}, nil
}
