package httpdata

import "time"

/*
{
  "errcode" : 0,
  "errmsg" : "ok",
  "invaliduser" : "userid1|userid2",
  "invalidparty" : "partyid1|partyid2",
  "invalidtag": "tagid1|tagid2",
  "msgid": "xxxx",
  "response_code": "xyzxyz"
}
*/

type WechatErr struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type SendMessageResponse struct {
	WechatErr
	InValidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
	MsgId        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}

type AccessTokenResponse struct {
	WechatErr
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type AccessTokenData struct {
	AccessToken string
	NextTime    time.Time
}
