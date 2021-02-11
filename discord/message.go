package discord

import "time"

// MessageCreate event struct
type MessageCreate struct {
	ChannelID       string    `json:"channel_id"`
	Content         string    `json:"content"`
	GuildID         string    `json:"guild_id"`
	ID              string    `json:"id"`
	Nonce           string    `json:"nonce"`
	Author          Author    `json:"author"`
	Member          Member    `json:"member"`
	Timestamp       time.Time `json:"timestamp"`
	MentionEveryone bool      `json:"mention_everyone"`
	Pinned          bool      `json:"pinned"`
	Tts             bool      `json:"tts"`
	Flags           int       `json:"flags"`
	Type            int       `json:"type"`
}
