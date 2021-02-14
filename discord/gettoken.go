package discord

import (
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/gorilla/websocket"
)

var (
	socketURL = "ws://127.0.0.1:6463/?v=1&encoding=json"
	headers   = http.Header{}
)

type sendReady struct {
	Cmd   string      `json:"cmd"`
	Args  interface{} `json:"args"`
	Evt   string      `json:"evt"`
	Nonce string      `json:"nonce"`
}

type connectArgs struct {
	Type string `json:"type"`
	Pid  int    `json:"pid"`
}

type connect struct {
	Cmd   string       `json:"cmd"`
	Args  *connectArgs `json:"args"`
	Nonce string       `json:"nonce"`
}

// GetToken - attempts to get users token from discord rpc ipc
func GetToken() (string, error) {
	var payloadSend = false
	var token string
	headers["Origin"] = []string{"https://discord.com"}
	var ready sendReady
	ready.Cmd = "SUBSCRIBE"
	ready.Evt = "OVERLAY"
	ready.Nonce = "test"

	var con = &connect{
		Cmd: "OVERLAY",
		Args: &connectArgs{
			Type: "CONNECT",
			Pid:  4,
		},
		Nonce: "test",
	}

	if c, _, err := websocket.DefaultDialer.Dial(socketURL, headers); err == nil {
		for {
			if payloadSend {
				break
			}
			if _, message, err := c.ReadMessage(); err == nil {
				jsonParsed, _ := gabs.ParseJSON(message)
				switch jsonParsed.Path("evt").String() {
				case "\"READY\"":
					c.WriteJSON(ready)
					c.WriteJSON(con)
					break
				default:
					if jsonParsed.Path("cmd").String() == "\"DISPATCH\"" && jsonParsed.Path("data.type").String() == "\"DISPATCH\"" && jsonParsed.Path("data.pid").String() == "4" {
						test, _ := jsonParsed.Path("data.payloads").ArrayElement(0)
						token = test.Path("token").String()
						token = strings.ReplaceAll(token, "\"", "")
						payloadSend = true
					}
					break
				}
			}
		}
		c.Close()
	} else {
		return token, err
	}
	return token, nil
}
