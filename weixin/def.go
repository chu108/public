package weixin

type UserInfo struct {
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int32    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

// APITicket ...
type APITicket struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

// WxInfo 微信配置信息
type WxInfo struct {
	AppID          string // 微信公众平台应用ID
	AppSecret      string // 微信支付商户平台商户号
	APIKey         string // 微信支付商户平台API密钥
	MchID          string
	NotifyURL      string
	ShearURL       string
	Token          string
	EncodingAESKey string
}
