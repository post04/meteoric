package discord

import (
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

var (
	baseURL = "https://discord.com/api/v6"
)

// BuildRequest builds a fasthttp request and returns request and response.
func BuildRequest(Method string, URL string, Body []byte) (req *fasthttp.Request, resp *fasthttp.Response) {
	req = fasthttp.AcquireRequest()
	req.Header.SetMethod(Method)
	req.SetRequestURI(URL)
	req.SetBody(Body)
	resp = fasthttp.AcquireResponse()
	return req, resp
}

// ClaimCode claims a discord gift code and returns http response.
func (s *Session) ClaimCode(Code string, ChannelID string) (body *fasthttp.Response, err error) {
	req, resp := BuildRequest("POST", fmt.Sprintf(baseURL+"/entitlements/gift-codes/%v/redeem", Code), []byte(fmt.Sprintf(`{"channel_id": "%v"}`, ChannelID)))
	req.Header.SetContentType("application/json")
	req.Header.Set("authorization", s.Token)
	if err := fasthttp.Do(req, resp); err == nil {
		return resp, nil
	}
	return nil, errors.New("claimer: error claiming nitro: " + err.Error())
}
