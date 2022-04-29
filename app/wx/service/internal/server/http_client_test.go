package server

import (
	"github.com/go-resty/resty/v2"
	"testing"
)

func TestSendGetReq(t *testing.T) {
	//at := assert.New(t)
	hc := resty.New()
	resp, err := hc.R().SetQueryParams(map[string]string{
		"grant_type": "client_credential",
		"appid":      "wx8253ecdcdaf841f5",
		"secret":     "9f77cb018789bbdd5d268a346628caec",
	}).Get("https://api.weixin.qq.com/cgi-bin/token")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(resp)
	}
}

func TestSendPostReq(t *testing.T) {
	hc := resty.New()
	resp, err := hc.R().SetBody(`
		 {
			 "button":[
			 {	
				  "type":"click",
				  "name":"今日歌曲",
				  "key":"V1001_TODAY_MUSIC"
			  },
			  {
				   "name":"菜单",
				   "sub_button":[
				   {	
					   "type":"view",
					   "name":"搜索",
					   "url":"https://www.baidu.com/"
					},
					{
						 "type":"view",
						 "name":"wxa",
						 "url":"http://mp.weixin.qq.com"
					 },
					{
					   "type":"click",
					   "name":"赞一下我们",
					   "key":"V1001_GOOD"
					}]
			   }]
		 }
	`).SetQueryParam(
		"access_token",
		"56_8MLND_bMmS8zPzNGNqpWgCesUoL-V81IdEHOry350s33EHoWm3ZWTdvejnhnqeLXoB0lwVz41RCnKerR_T5ocR1Ta_Q6E3AZfXq1KzdDqZ6h84YfFrc-rCjwWnl1_amRJx1ezmgYWzatXWRXHSTcAJABMN",
	).Post("https://api.weixin.qq.com/cgi-bin/menu/create")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(resp)
	}
}
