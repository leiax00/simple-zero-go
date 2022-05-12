package biz

import (
	"encoding/json"
	"encoding/xml"
	v1 "github.com/simple-zero-go/api/wx/service/v1"
	"testing"
)

func TestMenuStr(t *testing.T) {
	rst := &v1.Menu{}
	err := json.Unmarshal([]byte(menu), rst)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(rst.Button[0].Name)
	}

}

func TestWxObj(t *testing.T) {
	tmp := &MsgBase{
		ToUserName:   "asdasd",
		FromUserName: "dfgdbdbdf",
		CreateTime:   12433543754,
		MsgType:      "text",
	}
	str, _ := xml.Marshal(tmp)
	t.Log(string(str))

	tmp1 := TextMsg{
		MsgBase: *tmp,
		Content: "sdfsdkmkvsdvdshgsgsdg",
	}
	str, _ = xml.Marshal(tmp1)
	t.Log(string(str))
}

func TestTextMsg_FillByReq(t *testing.T) {
	req := &v1.MsgReq{
		ToUserName:   "asdafasfas",
		FromUserName: "sdgsdgasfas",
		CreateTime:   23534654363,
		MsgType:      "text",
		MsgId:        "safdsgsdgsd",
		Content:      "dfgdfhfdhdnbdndfb",
	}
	obj := &TextMsg{}
	obj.FillByReq(req)
	obj.FillMsg(Text(req.Content))
	bytes, err := xml.Marshal(obj)
	t.Log(string(bytes), err)
}

func TestVoiceMsg_FillByReq(t *testing.T) {
	req := &v1.MsgReq{
		ToUserName:   "asdafasfas",
		FromUserName: "sdgsdgasfas",
		CreateTime:   23534654363,
		MsgType:      "text",
		MsgId:        "safdsgsdgsd",
		MediaId:      "asfafasd",
		Format:       "sdfsdvsdv",
		Recognition:  "sdvnujvnsi",
	}
	obj := &VoiceMsg{}
	obj.FillByReq(req)
	obj.FillMsg(MediaBase{
		MediaId: req.MediaId,
	})
	str, _ := xml.Marshal(obj)
	t.Log(string(str))
	t.Log(Text(str))
}
