package weixin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/mylog"

	"github.com/silenceper/wechat"
)

const (
	_getTicket = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=wx_card&access_token="
	_getJsurl  = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="
)

// GetAccessToken 获取微信accesstoken
func GetAccessToken() (accessToken string, err error) {
	//先从缓存中获取
	cache := mycache.OnGetCache("weixin_token")
	var tp interface{}
	tp, b := cache.Value("base")
	if b {
		accessToken = tp.(string)
	} else {
		wc := wechat.NewWechat(&cfg)
		accessToken, err = wc.GetAccessToken()
		//保存缓存
		cache.Add("base", accessToken, 7000*time.Second)
	}
	return
}

// GetAPITicket 获取微信卡券ticket
func GetAPITicket() (ticket string, err error) {
	//先从缓存中获取
	cache := mycache.OnGetCache("weixin_card_ticket")
	var tp interface{}
	tp, b := cache.Value("base")
	if b {
		ticket = tp.(string)
	} else {
		accessToken, e := GetAccessToken()
		if e != nil {
			mylog.Error(e)
			err = e
			return
		}
		var url = _getTicket + accessToken

		resp, e1 := http.Get(url)
		if e1 != nil {
			mylog.Error(e1)
			err = e1
			return
		}
		defer resp.Body.Close()
		body, e2 := ioutil.ReadAll(resp.Body)
		if e2 != nil {
			mylog.Error(e2)
			err = e2
			return
		}
		var result APITicket
		json.Unmarshal(body, &result)
		ticket = result.Ticket
		//保存缓存
		cache.Add("base", ticket, 7000*time.Second)
	}
	return
}

// GetJsTicket 获取微信js ticket
func GetJsTicket() (ticket string, err error) {
	//先从缓存中获取
	cache := mycache.OnGetCache("weixin_js_ticket")
	var tp interface{}
	tp, b := cache.Value("base")
	if b {
		ticket = tp.(string)
	} else {
		accessToken, e := GetAccessToken()
		if e != nil {
			mylog.Error(e)
			err = e
			return
		}
		var url = _getJsurl + accessToken

		resp, e1 := http.Get(url)
		if e1 != nil {
			mylog.Error(e1)
			err = e1
			return
		}
		defer resp.Body.Close()
		body, e2 := ioutil.ReadAll(resp.Body)
		if e2 != nil {
			mylog.Error(e2)
			err = e2
			return
		}
		var result APITicket
		json.Unmarshal(body, &result)
		ticket = result.Ticket
		//保存缓存
		cache.Add("base", ticket, 7000*time.Second)
	}
	return
}
