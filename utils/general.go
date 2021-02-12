package utils

import (
	"github.com/valyala/fasthttp"
)

// Find a value inside the specified list.
func Find(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
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
