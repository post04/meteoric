package discord

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

// Session struct
type Session struct {
	Token      string
	State      *Ready
	Connection *websocket.Conn
	Mux        sync.Mutex
	Sequence   int
	Intents    int
}

// Payload struct
type Payload struct {
	Type string          `json:"t"`
	Data json.RawMessage `json:"d"`
	Op   int             `json:"op"`
}

// Ready event struct
type Ready struct {
	SessionID string  `json:"session_id"`
	User      User    `json:"user"`
	Guilds    []Guild `json:"guilds"`
}

// New creates a discord client.
func New(Token string, Intents int) (sess *Session, err error) {
	sess = &Session{
		Token:    Token,
		Sequence: 0,
		Intents:  Intents,
	}

	if c, _, err := websocket.DefaultDialer.Dial("wss://gateway.discord.gg?v=8&encoding=json", nil); err == nil {
		sess.Connection = c
	} else {
		return nil, err
	}

	return sess, nil
}

// HandleEvent hanldes event and registers.
func (s *Session) HandleEvent(eventName string, eventData []byte) {
	// Find a better way to do this.
	switch eventName {

	case "READY":
		err := json.Unmarshal(eventData, &s.State)
		if err != nil {
			fmt.Println("unmarshalling ready event:", err)
		}

		event.Register(eventName, s, s.State)

	case "MESSAGE_CREATE":
		var dmsg *MessageCreate
		err := json.Unmarshal(eventData, &dmsg)
		if err != nil {
			fmt.Println("unmarshalling message_create:", err)
		}

		event.Register(eventName, s, dmsg)

	default:
		event.Register(eventName, s, eventData)
	}
}

// Open connection and listen for events.
func (s *Session) Open() {
	for {
		if _, message, err := s.Connection.ReadMessage(); err == nil {
			jsonParsed, _ := gabs.ParseJSON(message)
			switch jsonParsed.Path("op").String() {
			case "0":
				sequence := jsonParsed.Path("s").Data().(float64)
				if int(sequence) > s.Sequence {
					s.Sequence = int(sequence)
				}

				eventName := jsonParsed.Path("t").Data().(string)
				data := jsonParsed.Path("d").Bytes()
				go s.HandleEvent(eventName, data)

			case "1":
				s.SendHeartbeat()

			case "10":
				s.SendIdentify()
				go s.StartHeart(int(jsonParsed.Path("d.heartbeat_interval").Data().(float64)))

			}
		}
	}
}

// StartHeart starts sending heartbeats.
func (s *Session) StartHeart(Interval int) {
	tick := time.NewTicker(time.Millisecond * time.Duration(int(Interval)))
	defer tick.Stop()

	for range tick.C {
		s.SendHeartbeat()
	}
}

// SendHeartbeat sends heartbeat payload.
func (s *Session) SendHeartbeat() {
	s.Mux.Lock()
	s.Connection.WriteJSON(bson.M{
		"op": 1,
		"d":  s.Sequence,
	})
	s.Mux.Unlock()
}

// SendIdentify sends identify payload.
func (s *Session) SendIdentify() {
	s.Mux.Lock()
	s.Connection.WriteJSON(bson.M{
		"op": 2,
		"d": bson.M{
			"intents":      s.Intents,
			"capabilities": 61,
			"token":        s.Token,
			"properties": bson.M{
				"os":                       "Windows",
				"browser":                  "Chrome",
				"device":                   "",
				"browser_user_agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.0.0 Safari/537.36",
				"browser_version":          "88.0.0.0",
				"os_version":               "10",
				"referrer":                 "",
				"referring_domain":         "",
				"referrer_current":         "",
				"referring_domain_current": "",
				"release_channel":          "stable",
				"client_build_number":      76069,
				"client_event_source":      nil,
			},
			"presence": bson.M{
				"status":     "online",
				"since":      0,
				"activities": []string{},
				"afk":        false,
			},
			"compress": false,
			"client_state": bson.M{
				"guild_hashes":                bson.M{},
				"highest_last_message_id":     "0",
				"read_state_version":          0,
				"user_guild_settings_version": -1,
			},
		},
	})
	s.Mux.Unlock()
}
