package models

import "time"

// Application represents a Discord application structure.
type Application struct {
	ID                  string   `json:"id,omitempty"`
	Name                string   `json:"name"`
	Icon                string   `json:"icon,omitempty"`
	Description         string   `json:"description,omitempty"`
	RPCOrigins          []string `json:"rpc_origins,omitempty"`
	BotPublic           bool     `json:"bot_public,omitempty"`
	BotRequireCodeGrant bool     `json:"bot_require_code_grant,omitempty"`
	TermsOfServiceURL   string   `json:"terms_of_service_url"`
	PrivacyProxyURL     string   `json:"privacy_policy_url"`
	Owner               *User    `json:"owner"`
	Summary             string   `json:"summary"`
	VerifyKey           string   `json:"verify_key"`
	Team                *Team    `json:"team"`
	GuildID             string   `json:"guild_id"`
	PrimarySKUID        string   `json:"primary_sku_id"`
	Slug                string   `json:"slug"`
	CoverImage          string   `json:"cover_image"`
	Flags               int      `json:"flags,omitempty"`
}

// ApplicationCommand represents a command structure in Discord.
type ApplicationCommand struct {
	ID                       string                      `json:"id"`                                   // Unique ID of the command
	Type                     int                         `json:"type,omitempty"`                       // Type of command, defaults to 1
	ApplicationID            string                      `json:"application_id"`                       // ID of the parent application
	GuildID                  *string                     `json:"guild_id,omitempty"`                   // Guild ID of the command, if not global
	Name                     string                      `json:"name"`                                 // Name of the command, 1-32 characters
	NameLocalizations        *map[string]string          `json:"name_localizations,omitempty"`         // Localization dictionary for name field
	Description              string                      `json:"description"`                          // Description for CHAT_INPUT commands, 1-100 characters. Empty for USER and MESSAGE commands
	DescriptionLocalizations *map[string]string          `json:"description_localizations,omitempty"`  // Localization dictionary for description field
	Options                  []*ApplicationCommandOption `json:"options,omitempty"`                    // Parameters for the command, max of 25 (CHAT_INPUT commands)
	DefaultMemberPermissions *string                     `json:"default_member_permissions,omitempty"` // Set of permissions represented as a bit set
	DMPermission             *bool                       `json:"dm_permission,omitempty"`              // Deprecated: Indicates if the command is available in DMs for global commands
	DefaultPermission        *bool                       `json:"default_permission,omitempty"`         // Deprecated: Indicates if the command is enabled by default when the app is added to a guild
	NSFW                     *bool                       `json:"nsfw,omitempty"`                       // Indicates whether the command is age-restricted, defaults to false
	IntegrationTypes         []string                    `json:"integration_types,omitempty"`          // Installation contexts for the command (global commands)
	Contexts                 *[]string                   `json:"contexts,omitempty"`                   // Interaction context(s) where the command can be used (global commands)
	Version                  string                      `json:"version"`                              // Auto-incrementing version identifier updated during substantial record changes
	Handler                  *string                     `json:"handler,omitempty"`                    // Determines whether the interaction is handled by the app's handler or by Discord
}

// ApplicationCommandOption represents an option for an Application Command.
type ApplicationCommandOption struct {
	Type                     int                               `json:"type"`                                // Type of the option
	Name                     string                            `json:"name"`                                // 1-32 character name
	NameLocalizations        *map[string]string                `json:"name_localizations,omitempty"`        // Optional localization dictionary for the name field
	Description              string                            `json:"description"`                         // 1-100 character description
	DescriptionLocalizations *map[string]string                `json:"description_localizations,omitempty"` // Optional localization dictionary for the description field
	Required                 *bool                             `json:"required,omitempty"`                  // Whether the option is required, default false
	Choices                  []*ApplicationCommandOptionChoice `json:"choices,omitempty"`                   // Optional array of choices for the user to pick from
	Options                  []*ApplicationCommandOption       `json:"options,omitempty"`                   // Nested options if the option is a subcommand or subcommand group
	ChannelTypes             []int                             `json:"channel_types,omitempty"`             // Restrict shown channels to these types
	MinValue                 *float64                          `json:"min_value,omitempty"`                 // Minimum value permitted (integer for INTEGER, double for NUMBER)
	MaxValue                 *float64                          `json:"max_value,omitempty"`                 // Maximum value permitted (integer for INTEGER, double for NUMBER)
	MinLength                *int                              `json:"min_length,omitempty"`                // Minimum allowed length (min 0, max 6000)
	MaxLength                *int                              `json:"max_length,omitempty"`                // Maximum allowed length (min 1, max 6000)
	Autocomplete             *bool                             `json:"autocomplete,omitempty"`              // Whether autocomplete is enabled for this option
}

// ApplicationCommandOptionChoice represents a choice for an ApplicationCommandOption.
type ApplicationCommandOptionChoice struct {
	Name  string      `json:"name"`  // The name of the choice
	Value interface{} `json:"value"` // The value of the choice (can be string, integer, or number)
}

// User represents a Discord user structure.
type User struct {
	ID                   string                `json:"id"`                               // Snowflake ID of the user
	Username             string                `json:"username"`                         // Username of the user (not unique)
	Discriminator        string                `json:"discriminator"`                    // User's Discord tag (four-digit identifier)
	GlobalName           *string               `json:"global_name,omitempty"`            // Optional display name (for bots, the application name)
	Avatar               *string               `json:"avatar,omitempty"`                 // Optional avatar hash of the user
	Bot                  *bool                 `json:"bot,omitempty"`                    // Optional flag indicating if the user is a bot
	System               *bool                 `json:"system,omitempty"`                 // Optional flag indicating if the user is a system user
	MFEnabled            *bool                 `json:"mfa_enabled,omitempty"`            // Optional flag indicating if 2FA is enabled
	Banner               *string               `json:"banner,omitempty"`                 // Optional banner hash of the user
	AccentColor          *int                  `json:"accent_color,omitempty"`           // Optional banner color as an integer
	Locale               *string               `json:"locale,omitempty"`                 // Optional user's chosen language
	Verified             *bool                 `json:"verified,omitempty"`               // Optional flag indicating if the email is verified
	Email                *string               `json:"email,omitempty"`                  // Optional user's email
	Flags                *int                  `json:"flags,omitempty"`                  // Optional flags on the user's account
	PremiumType          *int                  `json:"premium_type,omitempty"`           // Optional Nitro subscription type
	PublicFlags          *int                  `json:"public_flags,omitempty"`           // Optional public flags on the user's account
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data,omitempty"` // Optional avatar decoration data
}

// AvatarDecorationData represents data for a user's avatar decoration.
type AvatarDecorationData struct {
	Decoration string `json:"decoration"` // Example field for decoration
}

// Channel represents a Discord channel structure.
type Channel struct {
	ID                            string                 `json:"id"`                                           // Snowflake ID of the channel
	Type                          int                    `json:"type"`                                         // Type of the channel
	GuildID                       *string                `json:"guild_id,omitempty"`                           // Optional Guild ID if the channel is part of a guild
	Position                      *int                   `json:"position,omitempty"`                           // Optional sorting position of the channel
	PermissionOverwrites          *[]PermissionOverwrite `json:"permission_overwrites,omitempty"`              // Optional explicit permission overwrites for members and roles
	Name                          *string                `json:"name,omitempty"`                               // Optional name of the channel (1-100 characters)
	Topic                         *string                `json:"topic,omitempty"`                              // Optional topic of the channel (up to 4096 characters for forum/media channels, 1024 for others)
	NSFW                          *bool                  `json:"nsfw,omitempty"`                               // Optional flag indicating if the channel is NSFW
	LastMessageID                 *string                `json:"last_message_id,omitempty"`                    // Optional ID of the last message sent in the channel
	Bitrate                       *int                   `json:"bitrate,omitempty"`                            // Optional bitrate (in bits) of the voice channel
	UserLimit                     *int                   `json:"user_limit,omitempty"`                         // Optional user limit of the voice channel
	RateLimitPerUser              *int                   `json:"rate_limit_per_user,omitempty"`                // Optional rate limit per user (in seconds)
	Recipients                    *[]User                `json:"recipients,omitempty"`                         // Optional list of recipients in a DM
	Icon                          *string                `json:"icon,omitempty"`                               // Optional icon hash for group DM
	OwnerID                       *string                `json:"owner_id,omitempty"`                           // Optional owner ID for group DM or thread
	ApplicationID                 *string                `json:"application_id,omitempty"`                     // Optional application ID if bot-created group DM
	Managed                       *bool                  `json:"managed,omitempty"`                            // Optional flag indicating if the group DM is managed by an application
	ParentID                      *string                `json:"parent_id,omitempty"`                          // Optional ID of the parent category for guild channels
	LastPinTimestamp              *string                `json:"last_pin_timestamp,omitempty"`                 // Optional timestamp of when the last pinned message was pinned
	RTCRegion                     *string                `json:"rtc_region,omitempty"`                         // Optional voice region ID for the voice channel
	VideoQualityMode              *int                   `json:"video_quality_mode,omitempty"`                 // Optional video quality mode for the voice channel
	MessageCount                  *int                   `json:"message_count,omitempty"`                      // Optional count of messages in a thread (excludes initial and deleted messages)
	MemberCount                   *int                   `json:"member_count,omitempty"`                       // Optional count of users in a thread
	ThreadMetadata                *ThreadMetadata        `json:"thread_metadata,omitempty"`                    // Optional thread-specific fields
	Member                        *ThreadMember          `json:"member,omitempty"`                             // Optional thread member object for the current user
	DefaultAutoArchiveDuration    *int                   `json:"default_auto_archive_duration,omitempty"`      // Optional default duration (in minutes) for auto-archiving threads
	Permissions                   *string                `json:"permissions,omitempty"`                        // Optional computed permissions for the invoking user in the channel
	Flags                         *int                   `json:"flags,omitempty"`                              // Optional bitfield of channel flags
	TotalMessagesSent             *int                   `json:"total_message_sent,omitempty"`                 // Optional total number of messages ever sent in a thread
	AvailableTags                 *[]Tag                 `json:"available_tags,omitempty"`                     // Optional set of tags available in a forum/media channel
	AppliedTags                   *[]string              `json:"applied_tags,omitempty"`                       // Optional IDs of the tags applied to a thread in a forum/media channel
	DefaultReactionEmoji          *DefaultReaction       `json:"default_reaction_emoji,omitempty"`             // Optional default reaction emoji for threads
	DefaultThreadRateLimitPerUser *int                   `json:"default_thread_rate_limit_per_user,omitempty"` // Optional initial rate limit per user for newly created threads
	DefaultSortOrder              *int                   `json:"default_sort_order,omitempty"`                 // Optional default sort order for forum/media channels
	DefaultForumLayout            *int                   `json:"default_forum_layout,omitempty"`               // Optional default layout view for forum channels
}

// Invite represents a Discord invite structure.
type Invite struct {
	Guild             *Guild       `json:"guild"`
	Channel           *Channel     `json:"channel"`
	Inviter           *User        `json:"inviter"`
	Code              string       `json:"code"`
	CreatedAt         time.Time    `json:"created_at"`
	MaxAge            int          `json:"max_age"`
	Uses              int          `json:"uses"`
	MaxUses           int          `json:"max_uses"`
	Revoked           bool         `json:"revoked"`
	Temporary         bool         `json:"temporary"`
	Unique            bool         `json:"unique"`
	TargetUser        *User        `json:"target_user"`
	TargetType        uint8        `json:"target_type"`
	TargetApplication *Application `json:"target_application"`

	// will only be filled when using InviteWithCounts
	ApproximatePresenceCount int `json:"approximate_presence_count"`
	ApproximateMemberCount   int `json:"approximate_member_count"`

	ExpiresAt *time.Time `json:"expires_at"`
}

// PermissionOverwrite represents an overwrite object for a channel's permissions.
type PermissionOverwrite struct {
	ID    string `json:"id"`    // Snowflake ID of the overwrite (role or user)
	Type  int    `json:"type"`  // Type of overwrite (0 = role, 1 = member)
	Allow string `json:"allow"` // Allowed permissions bit set
	Deny  string `json:"deny"`  // Denied permissions bit set
}

// ThreadMetadata represents metadata specific to threads.
type ThreadMetadata struct {
	Archived            bool   `json:"archived"`              // Whether the thread is archived
	AutoArchiveDuration int    `json:"auto_archive_duration"` // Duration in minutes to auto-archive
	ArchiveTimestamp    string `json:"archive_timestamp"`     // ISO8601 timestamp when the thread was archived
	Locked              bool   `json:"locked"`                // Whether the thread is locked
}

// ThreadMember represents a member in a thread.
type ThreadMember struct {
	ID            string `json:"id"`             // Snowflake ID of the thread member
	UserID        string `json:"user_id"`        // Snowflake ID of the user
	JoinTimestamp string `json:"join_timestamp"` // ISO8601 timestamp when the member joined
	Flags         int    `json:"flags"`          // Thread member flags
}

// Tag represents a tag object for forum/media channels.
type Tag struct {
	ID    string `json:"id"`    // Snowflake ID of the tag
	Name  string `json:"name"`  // Name of the tag
	Emoji string `json:"emoji"` // Optional emoji associated with the tag
}

// DefaultReaction represents the default reaction emoji for threads in a forum/media channel.
type DefaultReaction struct {
	EmojiID   string `json:"emoji_id,omitempty"`   // Snowflake ID of the emoji
	EmojiName string `json:"emoji_name,omitempty"` // Name of the emoji
}

// Guild represents a Discord guild (server) structure.
type Guild struct {
	ID                          string         `json:"id"`                                   // Snowflake ID of the guild
	Name                        string         `json:"name"`                                 // Name of the guild (2-100 characters)
	Icon                        *string        `json:"icon,omitempty"`                       // Optional icon hash
	IconHash                    *string        `json:"icon_hash,omitempty"`                  // Optional icon hash returned in the template object
	Splash                      *string        `json:"splash,omitempty"`                     // Optional splash hash
	DiscoverySplash             *string        `json:"discovery_splash,omitempty"`           // Optional discovery splash hash for discoverable guilds
	Owner                       *bool          `json:"owner,omitempty"`                      // Optional flag indicating if the user is the owner of the guild
	OwnerID                     string         `json:"owner_id"`                             // ID of the owner
	Permissions                 *string        `json:"permissions,omitempty"`                // Optional total permissions for the user in the guild
	Region                      *string        `json:"region,omitempty"`                     // Optional voice region ID for the guild (deprecated)
	AfkChannelID                *string        `json:"afk_channel_id,omitempty"`             // Optional ID of AFK channel
	AfkTimeout                  int            `json:"afk_timeout"`                          // AFK timeout in seconds
	WidgetEnabled               *bool          `json:"widget_enabled,omitempty"`             // Optional flag indicating if the server widget is enabled
	WidgetChannelID             *string        `json:"widget_channel_id,omitempty"`          // Optional channel ID for widget invite
	VerificationLevel           int            `json:"verification_level"`                   // Verification level required for the guild
	DefaultMessageNotifications int            `json:"default_message_notifications"`        // Default message notifications level
	ExplicitContentFilter       int            `json:"explicit_content_filter"`              // Explicit content filter level
	Roles                       []Role         `json:"roles"`                                // Roles in the guild
	Emojis                      []Emoji        `json:"emojis"`                               // Custom guild emojis
	Features                    []string       `json:"features"`                             // Enabled guild features
	MFALevel                    int            `json:"mfa_level"`                            // Required MFA level for the guild
	ApplicationID               *string        `json:"application_id,omitempty"`             // Optional application ID if bot-created
	SystemChannelID             *string        `json:"system_channel_id,omitempty"`          // Optional system channel ID for notices
	SystemChannelFlags          int            `json:"system_channel_flags"`                 // System channel flags
	RulesChannelID              *string        `json:"rules_channel_id,omitempty"`           // Optional channel ID for community rules
	MaxPresences                *int           `json:"max_presences,omitempty"`              // Optional maximum presences for the guild
	MaxMembers                  int            `json:"max_members"`                          // Maximum number of members for the guild
	VanityURLCode               *string        `json:"vanity_url_code,omitempty"`            // Optional vanity URL code for the guild
	Description                 *string        `json:"description,omitempty"`                // Optional description of the guild
	Banner                      *string        `json:"banner,omitempty"`                     // Optional banner hash
	PremiumTier                 int            `json:"premium_tier"`                         // Premium tier (Server Boost level)
	PremiumSubscriptionCount    *int           `json:"premium_subscription_count,omitempty"` // Optional number of boosts
	PreferredLocale             string         `json:"preferred_locale"`                     // Preferred locale of the community guild
	PublicUpdatesChannelID      *string        `json:"public_updates_channel_id,omitempty"`  // Optional public updates channel ID
	MaxVideoChannelUsers        int            `json:"max_video_channel_users"`              // Maximum users in a video channel
	MaxStageVideoChannelUsers   int            `json:"max_stage_video_channel_users"`        // Maximum users in a stage video channel
	ApproximateMemberCount      *int           `json:"approximate_member_count,omitempty"`   // Optional approximate number of members
	ApproximatePresenceCount    *int           `json:"approximate_presence_count,omitempty"` // Optional approximate non-offline members
	WelcomeScreen               *WelcomeScreen `json:"welcome_screen,omitempty"`             // Optional welcome screen object
	NSFWLevel                   int            `json:"nsfw_level"`                           // NSFW level of the guild
	Stickers                    []Sticker      `json:"stickers"`                             // Custom guild stickers
	PremiumProgressBarEnabled   bool           `json:"premium_progress_bar_enabled"`         // Flag for boost progress bar enabled
	SafetyAlertsChannelID       *string        `json:"safety_alerts_channel_id,omitempty"`   // Optional channel ID for safety alerts
}

// Role represents a role object in the guild.
type Role struct {
	// Define fields for Role structure based on your needs
	ID          string `json:"id"`          // Snowflake ID of the role
	Name        string `json:"name"`        // Name of the role
	Color       int    `json:"color"`       // Color of the role
	Hoist       bool   `json:"hoist"`       // Whether the role is hoisted in the user list
	Position    int    `json:"position"`    // Position of the role
	Permissions string `json:"permissions"` // Permissions for the role
	Managed     bool   `json:"managed"`     // Whether the role is managed by an application
	Mentionable bool   `json:"mentionable"` // Whether the role is mentionable
}

// Emoji represents an emoji object in the guild.
type Emoji struct {
	// Define fields for Emoji structure based on your needs
	ID             string `json:"id"`              // Snowflake ID of the emoji
	Name           string `json:"name"`            // Name of the emoji
	Roles          []Role `json:"roles"`           // Roles allowed to use the emoji
	User           *User  `json:"user,omitempty"`  // Optional user object that created the emoji
	RequiresColons bool   `json:"requires_colons"` // Whether the emoji requires colons
	Managed        bool   `json:"managed"`         // Whether the emoji is managed by an application
	Animated       bool   `json:"animated"`        // Whether the emoji is animated
}

// WelcomeScreen represents the welcome screen for community guilds.
type WelcomeScreen struct {
	// Define fields relevant to the welcome screen as needed
	Description     *string          `json:"description,omitempty"` // Optional description for welcome screen
	WelcomeChannels []WelcomeChannel `json:"welcome_channels"`      // Channels to show on the welcome screen
}

// WelcomeChannel represents a channel in the welcome screen.
type WelcomeChannel struct {
	ChannelID   string  `json:"channel_id"`            // ID of the channel
	Description *string `json:"description,omitempty"` // Optional description for the channel
	Emoji       *string `json:"emoji,omitempty"`       // Optional emoji for the channel
}

// Sticker represents a sticker object in the guild.
type Sticker struct {
	// Define fields for Sticker structure based on your needs
	ID          string  `json:"id"`                    // Snowflake ID of the sticker
	PackID      string  `json:"pack_id"`               // ID of the sticker pack
	Name        string  `json:"name"`                  // Name of the sticker
	FormatType  int     `json:"format_type"`           // Format type of the sticker
	Description *string `json:"description,omitempty"` // Optional description of the sticker
}

type GuildBan struct {
	Reason string `json:"reason"`
	User   *User  `json:"user"`
}

// Member structure representing a user in a guild.
type Member struct {
	User                       *User                 `json:"user,omitempty"`                         // The user this guild member represents
	Nick                       *string               `json:"nick,omitempty"`                         // This user's guild nickname
	Avatar                     *string               `json:"avatar,omitempty"`                       // The member's guild avatar hash
	Roles                      []string              `json:"roles"`                                  // Array of role object IDs
	JoinedAt                   time.Time             `json:"joined_at"`                              // When the user joined the guild
	PremiumSince               *time.Time            `json:"premium_since,omitempty"`                // When the user started boosting the guild
	Deaf                       bool                  `json:"deaf"`                                   // Whether the user is deafened in voice channels
	Mute                       bool                  `json:"mute"`                                   // Whether the user is muted in voice channels
	Flags                      int                   `json:"flags"`                                  // Guild member flags represented as a bit set, defaults to 0
	Pending                    *bool                 `json:"pending,omitempty"`                      // Whether the user has not yet passed the guild's Membership Screening requirements
	Permissions                *string               `json:"permissions,omitempty"`                  // Total permissions of the member in the channel, including overwrites
	CommunicationDisabledUntil *time.Time            `json:"communication_disabled_until,omitempty"` // When the user's timeout will expire
	AvatarDecorationData       *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`       // Data for the member's guild avatar decoration
}

// Message structure representing a message sent in a channel.
type Message struct {
	ID                   string         `json:"id"`                               // ID of the message
	ChannelID            string         `json:"channel_id"`                       // ID of the channel the message was sent in
	Author               *User          `json:"author"`                           // The author of this message (not guaranteed to be a valid user)
	Content              string         `json:"content"`                          // Contents of the message
	Timestamp            time.Time      `json:"timestamp"`                        // When this message was sent
	EditedTimestamp      *time.Time     `json:"edited_timestamp,omitempty"`       // When this message was edited (or null if never)
	TTS                  bool           `json:"tts"`                              // Whether this was a TTS message
	MentionEveryone      bool           `json:"mention_everyone"`                 // Whether this message mentions everyone
	Mentions             []*User        `json:"mentions"`                         // Users specifically mentioned in the message
	MentionRoles         []string       `json:"mention_roles"`                    // Roles specifically mentioned in this message
	MentionChannels      []*interface{} `json:"mention_channels,omitempty"`       // Channels specifically mentioned in this message
	Attachments          []*interface{} `json:"attachments,omitempty"`            // Any attached files
	Embeds               []*interface{} `json:"embeds,omitempty"`                 // Any embedded content
	Reactions            []*Reaction    `json:"reactions,omitempty"`              // Reactions to the message
	Nonce                interface{}    `json:"nonce,omitempty"`                  // Used for validating a message was sent
	Pinned               bool           `json:"pinned"`                           // Whether this message is pinned
	WebhookID            *string        `json:"webhook_id,omitempty"`             // If the message is generated by a webhook
	Type                 int            `json:"type"`                             // Type of message
	Activity             *interface{}   `json:"activity,omitempty"`               // Sent with Rich Presence-related chat embeds
	Application          *interface{}   `json:"application,omitempty"`            // Sent with Rich Presence-related chat embeds
	ApplicationID        *string        `json:"application_id,omitempty"`         // ID of the application if the message is an Interaction or application-owned webhook
	Flags                int            `json:"flags"`                            // Message flags combined as a bitfield
	MessageReference     *interface{}   `json:"message_reference,omitempty"`      // Data showing the source of a crosspost, channel follow add, pin, or reply message
	MessageSnapshots     []*interface{} `json:"message_snapshots,omitempty"`      // The message associated with the message_reference
	ReferencedMessage    *Message       `json:"referenced_message,omitempty"`     // The message associated with the message_reference
	InteractionMetadata  *interface{}   `json:"interaction_metadata,omitempty"`   // Sent if the message is sent as a result of an interaction
	Interaction          *interface{}   `json:"interaction,omitempty"`            // Deprecated in favor of interaction_metadata
	Thread               *Channel       `json:"thread,omitempty"`                 // The thread that was started from this message
	Components           []*interface{} `json:"components,omitempty"`             // Sent if the message contains components like buttons, action rows, etc.
	StickerItems         []*interface{} `json:"sticker_items,omitempty"`          // Sent if the message contains stickers
	Stickers             []*Sticker     `json:"stickers,omitempty"`               // Deprecated the stickers sent with the message
	Position             int            `json:"position,omitempty"`               // Approximate position of the message in a thread
	RoleSubscriptionData *interface{}   `json:"role_subscription_data,omitempty"` // Data of the role subscription purchase or renewal
	Resolved             *interface{}   `json:"resolved,omitempty"`               // Data for users, members, channels, and roles in the message's auto-populated select menus
	Poll                 *interface{}   `json:"poll,omitempty"`                   // A poll!
	Call                 *interface{}   `json:"call,omitempty"`                   // The call associated with the message
}

// Reaction structure representing a reaction to a message.
type Reaction struct {
	Count        int           `json:"count"`                   // Total number of times this emoji has been used to react (including super reacts)
	CountDetails *CountDetails `json:"count_details,omitempty"` // Reaction count details object
	Me           bool          `json:"me"`                      // Whether the current user reacted using this emoji
	MeBurst      bool          `json:"me_burst"`                // Whether the current user super-reacted using this emoji
	Emoji        *PartialEmoji `json:"emoji"`                   // Emoji information
	BurstColors  []string      `json:"burst_colors,omitempty"`  // HEX colors used for super reaction
}

// CountDetails structure representing details about reaction counts.
type CountDetails struct {
	// Add specific fields for count details if needed
}

// PartialEmoji structure representing an emoji used in reactions.
type PartialEmoji struct {
	ID       string `json:"id,omitempty"`       // ID of the emoji (if it's a custom emoji)
	Name     string `json:"name"`               // Name of the emoji
	Animated bool   `json:"animated,omitempty"` // Whether the emoji is animated
}

// RoleParms structure representing the parameters for a role in a guild.
type RoleParams struct {
	Name         string `json:"name"`                    // Name of the role, max 100 characters
	Permissions  string `json:"permissions"`             // Bitwise value of the enabled/disabled permissions
	Color        int    `json:"color"`                   // RGB color value
	Hoist        bool   `json:"hoist"`                   // Whether the role should be displayed separately in the sidebar
	Icon         string `json:"icon,omitempty"`          // The role's icon image (if the guild has the ROLE_ICONS feature)
	UnicodeEmoji string `json:"unicode_emoji,omitempty"` // The role's unicode emoji as a standard emoji (if the guild has the ROLE_ICONS feature)
	Mentionable  bool   `json:"mentionable"`             // Whether the role should be mentionable
}

// TeamMember structure representing a member of a team.
type Team struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Icon        string        `json:"icon"`
	OwnerID     string        `json:"owner_user_id"`
	Members     []*TeamMember `json:"members"`
}

// TeamMember structure representing a member of a team.
type TeamMember struct {
	User            *User    `json:"user"`
	TeamID          string   `json:"team_id"`
	MembershipState int      `json:"membership_state"`
	Permissions     []string `json:"permissions"`
}

type GuildParams struct {
	Name                        string   `json:"name,omitempty"`
	Region                      string   `json:"region,omitempty"`
	VerificationLevel           int      `json:"verification_level,omitempty"`
	DefaultMessageNotifications int      `json:"default_message_notifications,omitempty"` // TODO: Separate type?
	ExplicitContentFilter       int      `json:"explicit_content_filter,omitempty"`
	AfkChannelID                string   `json:"afk_channel_id,omitempty"`
	AfkTimeout                  int      `json:"afk_timeout,omitempty"`
	Icon                        string   `json:"icon,omitempty"`
	OwnerID                     string   `json:"owner_id,omitempty"`
	Splash                      string   `json:"splash,omitempty"`
	DiscoverySplash             string   `json:"discovery_splash,omitempty"`
	Banner                      string   `json:"banner,omitempty"`
	SystemChannelID             string   `json:"system_channel_id,omitempty"`
	SystemChannelFlags          int      `json:"system_channel_flags,omitempty"`
	RulesChannelID              string   `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      string   `json:"public_updates_channel_id,omitempty"`
	PreferredLocale             string   `json:"preferred_locale,omitempty"`
	Features                    []string `json:"features,omitempty"`
	Description                 string   `json:"description,omitempty"`
	PremiumProgressBarEnabled   *bool    `json:"premium_progress_bar_enabled,omitempty"`
}
