package biz

import (
	"encoding/xml"
	v1 "github.com/simple-zero-go/api/wx/service/v1"
	"time"
)

type MsgReply interface {
	FillByReq(msg *v1.MsgReq)
	FillMsg(msg Msg)
	Reply(msg *v1.MsgReq) error
}

type Msg interface{}

type MsgBase struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `protobuf:"bytes,1,opt,name=toUserName,proto3" json:"toUserName,omitempty"`
	FromUserName string   `protobuf:"bytes,2,opt,name=fromUserName,proto3" json:"fromUserName,omitempty"`
	CreateTime   int64    `protobuf:"bytes,3,opt,name=createTime,proto3" json:"createTime,omitempty"`
	MsgType      string   `protobuf:"bytes,4,opt,name=msgType,proto3" json:"msgType,omitempty"`
}

func (m *MsgBase) Reply(msg *v1.MsgReq) error {
	panic("implement me")
}

func (m *MsgBase) FillMsg(msg Msg) {
	panic("implement me")
}

func (m *MsgBase) FillByReq(msg *v1.MsgReq) {
	m.FromUserName = msg.ToUserName
	m.ToUserName = msg.FromUserName
	m.CreateTime = time.Now().Unix()
	m.MsgType = msg.MsgType
}

type Text string

var _ Msg = (*Text)(nil)

type TextMsg struct {
	MsgBase
	Content Text `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (t *TextMsg) Reply(msg *v1.MsgReq) error {
	t.FillMsg(Text(msg.Content))
	return nil
}

func (t *TextMsg) FillMsg(msg Msg) {
	t.Content = msg.(Text)
}

type MediaBase struct {
	MediaId string `protobuf:"bytes,8,opt,name=mediaId,proto3" json:"mediaId,omitempty"`
}

var _ Msg = (*MediaBase)(nil)

type PicMsg struct {
	MsgBase
	Image *MediaBase `json:"image,omitempty"`
}

func (p *PicMsg) Reply(msg *v1.MsgReq) error {
	p.FillMsg(MediaBase{
		MediaId: msg.MediaId,
	})
	return nil
}

func (p *PicMsg) FillMsg(msg Msg) {
	t := msg.(MediaBase)
	p.Image = &t
}

type VoiceMsg struct {
	MsgBase
	Voice *MediaBase `json:"voice,omitempty"`
}

func (v *VoiceMsg) Reply(msg *v1.MsgReq) error {
	v.FillMsg(MediaBase{
		MediaId: msg.MediaId,
	})
	return nil
}

func (v *VoiceMsg) FillMsg(msg Msg) {
	t := msg.(MediaBase)
	v.Voice = &t
}

type MediaVideo struct {
	MediaBase
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type VideoMsg struct {
	MsgBase
	Video *MediaVideo `json:"video,omitempty"`
}

func (v *VideoMsg) Reply(msg *v1.MsgReq) error {
	v.FillMsg(MediaVideo{
		MediaBase:   MediaBase{MediaId: msg.MsgId},
		Title:       "Hello World",
		Description: "This is a test video!!!",
	})
	return nil
}

func (v *VideoMsg) FillMsg(msg Msg) {
	tmp := msg.(MediaVideo)
	v.Video = &tmp
}

type MediaMusic struct {
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	MusicUrl     string `json:"musicUrl,omitempty"`
	HQMusicUrl   string `json:"HQMusicUrl,omitempty"`
	ThumbMediaId string `json:"thumbMediaId,omitempty"`
}

type MusicMsg struct {
	MsgBase
	Music *MediaMusic `json:"music,omitempty"`
}

func (m *MusicMsg) Reply(msg *v1.MsgReq) error {
	m.FillMsg(MediaMusic{
		Title:        "曾经你说",
		Description:  "this is music!!!",
		MusicUrl:     "https://i.y.qq.com/v8/playsong.html?songid=314846798#webchat_redirect",
		HQMusicUrl:   "https://i.y.qq.com/v8/playsong.html?songid=314846798#webchat_redirect",
		ThumbMediaId: "",
	})
	return nil
}

func (m *MusicMsg) FillMsg(msg Msg) {
	tmp := msg.(MediaMusic)
	m.Music = &tmp
}

type Article struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	PicUrl      string `json:"PicUrl,omitempty"`
	Url         string `json:"url,omitempty"`
}

type ArticlePack struct {
	Item Article `json:"item,omitempty"`
}

type ArticleList struct {
	Articles []ArticlePack `json:"articles,omitempty"`
}

var _ Msg = (*ArticleList)(nil)

type ArticleMsg struct {
	MsgBase
	ArticleList
	ArticleCount int `json:"articleCount,omitempty"`
}

func (a *ArticleMsg) Reply(msg *v1.MsgReq) error {
	a.FillMsg(ArticleList{
		Articles: []ArticlePack{
			{
				Item: Article{
					Title:       "百度一下",
					Description: "百度一下",
					PicUrl:      "www.baidu.com/img/PCfb_5bf082d29588c07f842ccde3f97243ea.png",
					Url:         "https://www.baidu.com/",
				},
			},
		},
	})
	return nil
}

func (a *ArticleMsg) FillMsg(msg Msg) {
	t := msg.(ArticleList)
	a.ArticleList = t
	a.ArticleCount = len(t.Articles)
}
