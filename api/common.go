package api

import (
	"appcenter-wechat/httpdata"
	"appcenter-wechat/klog"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

var (
	AccessToken = httpdata.AccessTokenData{}
	JsonHeader  = map[string]string{"Content-Type": "application/json"}
	IpList      = []string{}
)

func sendRequest(url, method string, body io.Reader, params, headers map[string]string) (httpcode int, data io.Reader, err error) {
	req := &http.Request{}

	if req, err = http.NewRequest(url, method, body); err != nil {
		return
	}

	if params != nil {
		u := req.URL.Query()
		for k, v := range params {
			u.Add(k, v)
		}
		req.URL.RawQuery = u.Encode()
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DisableKeepAlives: true,
			// MaxIdleConns:      100, n * MaxCoonsPerHost
			MaxConnsPerHost: 100,
			IdleConnTimeout: 60 * time.Second,
		},
	}

	resp := &http.Response{}
	if resp, err = client.Do(req); err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("send http reqeust failed, err: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return http.StatusInternalServerError, nil, fmt.Errorf("send http request failed, statuscode[%d] is not 200", resp.StatusCode)
	}

	buffer := new(bytes.Buffer)

	_, err = io.Copy(buffer, resp.Body)

	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("http write response failed, err: %s", err.Error())
	}

	return http.StatusOK, buffer, nil
}

func getToken() (err error) {
	tokenUrl := fmt.Sprintf(accessTokenUrl, corpId, corpSecret)

	if AccessToken.AccessToken != "" && time.Now().Before(AccessToken.NextTime) {
		return nil
	}

	httpcode, data, err := sendRequest(tokenUrl, http.MethodGet, nil, nil, JsonHeader)

	if err != nil {
		klog.Logger.Errorf("%s, resp: %s, httpcode: %d", err.Error(), data, httpcode)
		return
	}

	accessTokenResponse := httpdata.AccessTokenResponse{}

	if err = json.NewDecoder(data).Decode(&accessTokenResponse); err != nil {
		return fmt.Errorf("decode data failed, err: %s", err.Error())
	}

	AccessToken.AccessToken = accessTokenResponse.AccessToken
	AccessToken.NextTime = time.Now().Add(time.Duration(accessTokenResponse.ExpiresIn) * time.Second)

	klog.Logger.Info("accessToken: ", AccessToken.AccessToken, AccessToken.NextTime)
	return
}

func getWechatServerIpList() (err error) {

	return
}
