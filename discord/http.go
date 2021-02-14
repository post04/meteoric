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
func BuildRequest(Method, URL string, Body []byte) (req *fasthttp.Request, resp *fasthttp.Response) {
	req = fasthttp.AcquireRequest()
	req.Header.SetMethod(Method)
	req.SetRequestURI(URL)
	req.SetBody(Body)
	resp = fasthttp.AcquireResponse()
	return req, resp
}

// ClaimCode claims a discord gift code and returns http response.
func (s *Session) ClaimCode(Code, ChannelID, Token string) (body *fasthttp.Response, err error) {
	req, resp := BuildRequest("POST", fmt.Sprintf(baseURL+"/entitlements/gift-codes/%v/redeem", Code), []byte(fmt.Sprintf(`{"channel_id": "%v"}`, ChannelID)))
	req.Header.SetContentType("application/json")
	req.Header.Set("authorization", Token)
	if err := fasthttp.Do(req, resp); err == nil {
		return resp, nil
	}
	return nil, errors.New("claimer: error claiming nitro: " + err.Error())
}

// CheckToken - checks to make sure the token claiming the nitro is a valid token.
func CheckToken(token string) bool {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.SetRequestURI("https://discordapp.com/api/v7/users/@me")
	req.Header.Set("Authorization", token)
	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, resp); err != nil {
		return false
	}
	if resp.StatusCode() == 200 {
		return true
	}
	return false
}
