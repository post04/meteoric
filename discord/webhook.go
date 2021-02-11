package discord

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// Webhook struct...
type Webhook struct {
	webhookURL   string
	usingWebhook bool
}

// NewWebhook creates a new webhook client.
func NewWebhook(WebhookURL string) (whook *Webhook) {
	whook = &Webhook{
		webhookURL:   WebhookURL,
		usingWebhook: false,
	}

	if whook.CheckWebhookExists() {
		whook.usingWebhook = true
	}
	
	return whook
}

// CheckWebhookExists checks if a webhook is valid or not.
func (w *Webhook) CheckWebhookExists() bool {
	req, resp := BuildRequest("GET", w.webhookURL, nil)
	if err := fasthttp.Do(req, resp); err == nil {
		if resp.StatusCode() == 200 {
			return true
		}
		return false
	}

	return false
}

// LogInfo ...
func (w *Webhook) LogInfo(NitroState string, Code string, AuthorName string, Elapsed string, StateUsername string, Color int) error {
	if w.usingWebhook {
		return w.SendMessage([]byte(fmt.Sprintf(`{"content": null, "embeds": [{"title": "%v", "description": "CODE: **%v**\nAUTHOR: **%v**\nELAPSED: **%v**\n\n", "color": %v, "author": { "name": "%v"}}]}`, NitroState, Code, AuthorName, Elapsed, Color, StateUsername)))
	}

	return nil
}

// SendMessage through the webhook to a discord channel.
func (w *Webhook) SendMessage(Content []byte) error {
	req, resp := BuildRequest("POST", w.webhookURL, Content)
	req.Header.SetContentType("application/json")
	if err := fasthttp.Do(req, resp); err != nil {
		return err
	}

	return nil
}
