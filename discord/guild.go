package discord

// Guild struct
type Guild struct {
	Features                 []string  `json:"features"`
	ID                       string    `json:"id"`
	Name                     string    `json:"name"`
	Icon                     string    `json:"icon"`
	Region                   string    `json:"region"`
	AfkChannelID             string    `json:"afk_channel_id"`
	OwnerID                  string    `json:"owner_id"`
	WidgetChannelID          string    `json:"widget_channel_id"`
	SystemChannelID          string    `json:"system_channel_id"`
	RulesChannelID           string    `json:"rules_channel_id"`
	VanityURLCode            string    `json:"vanity_url_code"`
	Description              string    `json:"description"`
	Banner                   string    `json:"banner"`
	DiscoverySplash          string    `json:"discovery_splash"`
	Splash                   string    `json:"splash"`
	ApplicationID            string    `json:"application_id"`
	PreferredLocale          string    `json:"preferred_locale"`
	PublicUpdatesChannelID   string    `json:"public_updates_channel_id"`
	Members                  []Member `json:"members"`
	Owner                    bool      `json:"owner"`
	WidgetEnabled            bool      `json:"widget_enabled"`
	Large                    bool      `json:"large"`
	Unavailable              bool      `json:"unavailable"`
	AfkTimeout               int       `json:"afk_timeout"`
	MemberCount              int       `json:"member_count"`
	MaxPresences             int       `json:"max_presences"`
	MaxMembers               int       `json:"max_members"`
	ExplicitContentFilter    int       `json:"explicit_content_filter"`
	MfaLevel                 int       `json:"mfa_level"`
	PremiumSubscriptionCount int       `json:"premium_subscription_count"`
	MaxVideoChannelUsers     int       `json:"max_video_channel_users"`
	ApproximateMemberCount   int       `json:"approximate_member_count"`
	ApproximatePresenceCount int       `json:"approximate_presence_count"`
	Permissions              int64     `json:"permissions,string"`
}
