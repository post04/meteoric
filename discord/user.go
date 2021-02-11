package discord

import (
	"time"
)

// Member struct
type Member struct {
	Roles       []string  `json:"roles"`
	HoistedRole string    `json:"hoisted_role"`
	JoinedAt    time.Time `json:"joined_at"`
	Deaf        bool      `json:"deaf"`
	Mute        bool      `json:"mute"`
}

// Author struct
type Author struct {
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	ID            string `json:"id"`
	Username      string `json:"username"`
	PublicFlags   int    `json:"public_flags"`
}

// User struct
type User struct {
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	ID            string `json:"id"`
	Phone         string `json:"phone"`
	Desktop       bool   `json:"desktop"`
	MfaEnabled    bool   `json:"mfa_enabled"`
	Mobile        bool   `json:"mobile"`
	NsfwAllowed   bool   `json:"nsfw_allowed"`
	Premium       bool   `json:"premium"`
	Verified      bool   `json:"verified"`
	Flags         int    `json:"flags"`
}

// ClientUsername returns full client username.
func (s *Session) ClientUsername() string {
	return s.State.User.Username + "#" + s.State.User.Discriminator
}

// Guilds return user guilds.
func (s *Session) Guilds() []Guild {
	return s.State.Guilds
}
