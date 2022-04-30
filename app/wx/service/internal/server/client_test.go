package server

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"testing"
	"time"
)

func TestHttpSendGetReq(t *testing.T) {
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

func TestHttpSendPostReq(t *testing.T) {
	hc := resty.New()
	resp, err := hc.R().SetBody(`
		{
		  "button": [
		    {
		      "type": "view",
		      "name": "百度一下",
		      "url": "https://www.baidu.com/"
		    },
		    {
		      "type": "view",
		      "name": "开发文档",
		      "url": "https://developers.weixin.qq.com/doc"
		    },
		    {
		      "name": "菜单",
		      "sub_button": [
		        {
		          "type": "view",
		          "name": "搜索",
		          "url": "https://www.baidu.com/"
		        },
		        {
		          "type": "view",
		          "name": "wxa",
		          "url": "http://mp.weixin.qq.com"
		        },
		        {
		          "type": "click",
		          "name": "赞一下我们",
		          "key": "V1001_GOOD"
		        }
		      ]
		    }
		  ]
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

func TestNewRedisClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rc := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
	result, err := rc.Set(ctx, "test_key", "Hello, baby!!", redis.KeepTTL).Result()
	_, _ = rc.Set(ctx, "test_key_1", "Hello, baby!!", 50).Result()
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(result)
	}
	result, err = rc.Get(ctx, "test_key").Result()
	t.Logf("%s, %v", result, err)
	result, err = rc.Get(ctx, "test_key_2").Result()
	t.Logf("%s, %v", result, err == redis.Nil)
}
