package biz

import (
	v1 "github.com/simple-zero-go/api/wx/service/v1"
	_const "github.com/simple-zero-go/app/wx/service/internal/const"
)

// GetMsgReplyByReq 目前收到什么消息就回复什么消息
func GetMsgReplyByReq(msg *v1.MsgReq) MsgReply {
	var obj MsgReply
	switch msg.MsgType {
	case _const.MsgTypeText:
		obj = &TextMsg{}
	case _const.MsgTypeImage:
		obj = &PicMsg{}
	case _const.MsgTypeVoice:
		obj = &VoiceMsg{}
	case _const.MsgTypeVideo:
		obj = &VideoMsg{}
	case _const.MsgTypeMusic:
		obj = &MusicMsg{}
	case _const.MsgTypeNews:
		obj = &ArticleMsg{}
	default:
		obj = &TextMsg{}
	}
	obj.FillByReq(msg)
	return obj
}
