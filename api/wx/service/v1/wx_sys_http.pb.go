// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type WxSysHTTPServer interface {
	AuthServer(context.Context, *AuthServerReq) (*AuthServerResp, error)
	DispatchMsg(context.Context, *MsgReq) (*StringReply, error)
}

func RegisterWxSysHTTPServer(s *http.Server, srv WxSysHTTPServer) {
	r := s.Route("/")
	r.GET("/wx/v1/handle", _WxSys_AuthServer0_HTTP_Handler(srv))
	r.POST("/wx/v1/handle", _WxSys_DispatchMsg0_HTTP_Handler(srv))
}

func _WxSys_AuthServer0_HTTP_Handler(srv WxSysHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthServerReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/wx.service.v1.WxSys/AuthServer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuthServer(ctx, req.(*AuthServerReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthServerResp)
		return ctx.Result(200, reply.Echostr)
	}
}

func _WxSys_DispatchMsg0_HTTP_Handler(srv WxSysHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MsgReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/wx.service.v1.WxSys/DispatchMsg")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DispatchMsg(ctx, req.(*MsgReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StringReply)
		return ctx.Result(200, reply.Msg)
	}
}

type WxSysHTTPClient interface {
	AuthServer(ctx context.Context, req *AuthServerReq, opts ...http.CallOption) (rsp *AuthServerResp, err error)
	DispatchMsg(ctx context.Context, req *MsgReq, opts ...http.CallOption) (rsp *StringReply, err error)
}

type WxSysHTTPClientImpl struct {
	cc *http.Client
}

func NewWxSysHTTPClient(client *http.Client) WxSysHTTPClient {
	return &WxSysHTTPClientImpl{client}
}

func (c *WxSysHTTPClientImpl) AuthServer(ctx context.Context, in *AuthServerReq, opts ...http.CallOption) (*AuthServerResp, error) {
	var out AuthServerResp
	pattern := "/wx/v1/handle"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/wx.service.v1.WxSys/AuthServer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out.Echostr, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WxSysHTTPClientImpl) DispatchMsg(ctx context.Context, in *MsgReq, opts ...http.CallOption) (*StringReply, error) {
	var out StringReply
	pattern := "/wx/v1/handle"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/wx.service.v1.WxSys/DispatchMsg"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out.Msg, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
