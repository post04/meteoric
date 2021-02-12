package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"./discord"
	"./utils"
	"github.com/gookit/color"
)

func main() {
	cfg := utils.GetConfig()

	if cfg.GeneralConfig.EnvConfig.UseEnvFile {
		Token = os.Getenv(cfg.GeneralConfig.EnvConfig.TokenValueName)
	} else {
		Token = cfg.GeneralConfig.Token
	}

	if sess, err := discord.New(Token, 512); err == nil {
		wHook = discord.NewWebhook(cfg.GeneralConfig.Logging.WebhookURL)
		discord.On("READY", onReady)
		discord.On("MESSAGE_CREATE", onMessage)
		sess.Open()
	}
}

func onReady(s *discord.Session, r *discord.Ready) {
	color.Printf("<magenta>%v</> | Logged in as %v#%v\n | Guilds: %v\n", time.Now().Format(TimeFormat), s.State.User.Username, s.State.User.Discriminator, len(s.Guilds()))
}

func onMessage(s *discord.Session, m *discord.MessageCreate) {
	if r.MatchString(m.Content) {
		code := r.FindStringSubmatch(m.Content)
		found := utils.Find(CachedNitro, code[3])
		if !found && len(code[3]) >= 16 && len(code[3]) <= 24 {
			now := time.Now()
			if nresp, err := s.ClaimCode(code[3], m.ChannelID); err == nil {
				elapsed := time.Since(now).String()

				var authorUsername = m.Author.Username + "#" + m.Author.Discriminator
				var clientUsername = s.ClientUsername()

				switch nresp.StatusCode() {
				case 404:
					color.Printf("<magenta>%v</> | CODE: <danger>%v</> — STATUS: %v (Unknown) — ELAPSED: %v\n", time.Now().Format(TimeFormat), code[3], nresp.StatusCode(), elapsed)
					wHook.LogInfo("Unknown", code[3], authorUsername, elapsed, clientUsername, 12632256)
				case 400:
					color.Printf("<magenta>%v</> | CODE: <yellow>%v</> — STATUS: %v (Already claimed) — ELAPSED: %v\n", time.Now().Format(TimeFormat), code[3], nresp.StatusCode(), elapsed)
					wHook.LogInfo("Already Claimed", code[3], authorUsername, elapsed, clientUsername, 16776960)
				case 200:
					color.Printf("<magenta>%v</> | CODE: <suc>%v</> — STATUS: %v (Valid) — ELAPSED: %v\n", time.Now().Format(TimeFormat), code[3], nresp.StatusCode(), elapsed)
					wHook.LogInfo("Valid", code[3], authorUsername, elapsed, clientUsername, 65390)
				default:
					color.Printf("<magenta>%v</> | CODE: <err>%v</> — STATUS: %v (???) — ELAPSED: %v\n", time.Now().Format(TimeFormat), code[3], nresp.StatusCode(), elapsed)
					wHook.LogInfo("???", code[3], authorUsername, elapsed, clientUsername, 0)
				}
				CachedNitro = append(CachedNitro, code[3])
			} else {
				fmt.Println(err)
			}
		}
	}
}

// Variables
var (
	TimeFormat  = "15:04:05 — 01/02-06"
	wHook       *discord.Webhook
	Token       string
	CachedNitro []string
	r = regexp.MustCompile(`(discord|discordapp)(\.gift\/|\.com\/gifts\/)([a-zA-Z0-9]+)`)
)
