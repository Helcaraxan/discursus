package users

import (
	"github.com/helcaraxan/discursus/atoms"
	"github.com/helcaraxan/discursus/groups"
	"github.com/helcaraxan/discursus/topics"
)

type UsersEmailInviteRequest struct {
	Email         string `json:"email"`
	GroupNames    string `json:"group_names"`
	CustomMessage string `json:"custom_message"`
}

type UsersEmailInviteResponse struct {
	Success string `json:"success"`
}

type UserGenerateInviteLinkRequest UsersEmailInviteRequest

type UserGenerateInviteLinkResponse string

type UserGetPrivateMessagesRequest struct{}

type UserGetPrivateMessagesResponse struct {
	TopicList topics.TopicList `json:"topic_list"`
}

type UserGetPrivateMessageSentRequest struct{}

type UserGetPrivateMessageSentResponse struct {
	Users         []*atoms.UserReference         `json:"users"`
	PrimaryGroups []*atoms.PrimaryGroupReference `json:"primary_groups"`
	TopicList     topics.TopicList               `json:"topic_list"`
}

type User struct {
	atoms.UserReference
	atoms.UserType
	atoms.PrimaryGroup

	Title       string `json:"title"`
	Email       string `json:"email"`
	BioRaw      string `json:"bio_raw"`
	BioCooked   string `json:"bio_cooked"`
	BioExcerpt  string `json:"bio_excerpt"`
	Website     string `json:"website"`
	WebsiteName string `json:"website_name"`
	Location    string `json:"location"`
	Locale      string `json:"locale"`
	DateOfBirth string `json:"date_of_birth"`

	CreatedAt    string `json:"created_at"`
	InvitedBy    string `json:"invited_by"`
	LastPostedAt string `json:"last_posted_at"`
	LastSeenAt   string `json:"last_seen_at"`

	CardBadge            string `json:"card_badge"`
	BadgeCount           int    `json:"badge_count"`
	FeaturedUserBadgeIDs []int  `json:"featured_user_badge_ids"`
	ProfileViewCount     int    `json:"profile_view_count"`
	PendingCount         int    `json:"pending_count"`
	PostCount            int    `json:"post_count"`

	TrustLevel             int  `json:"trust_level"`
	CanEdit                bool `json:"can_edit"`
	CanEditUsername        bool `json:"can_edit_username"`
	CanEditEmail           bool `json:"can_edit_email"`
	CanChangeBio           bool `json:"can_change_bio"`
	CanSendPrivateMessages bool `json:"can_send_private_messages"`
	CanDeleteAllPosts      bool `json:"can_delete_all_posts"`

	CanBeDeleted                bool `json:"can_be_deleted"`
	CanSentPrivateMessageToUser bool `json:"can_send_private_message_to_user"`
	HasTitleBadges              bool `json:"has_title_badges"`

	CustomFields []struct{} `json:"custom_fields"`

	MutedCategoryIDs            []int `json:"muted_category_ids"`
	TrackedCategoryIDs          []int `json:"tracked_category_ids"`
	WatchedCategoryIDs          []int `json:"watched_category_ids"`
	WatchedFirstPostCategoryIDs []int `json:"watched_first_post_category_ids"`

	MutedTags             []string `json:"muted_tags"`
	TrackedTags           []string `json:"tracked_tags"`
	WatchedTags           []string `json:"watched_tags"`
	WatchingFirstPostTags []string `json:"watching_first_post_tags"`

	MutedUsernames []string `json:"muted_usernames"`

	MailingListPostsPerDay int `json:"mailing_list_posts_per_day"`

	SystemAvatarUploadID   int    `json:"system_avatar_upload_id"`
	SystemAvatarTemplate   string `json:"system_avatar_template"`
	GravatarAvatarUploadID int    `json:"gravatar_avatar_upload_id"`
	GravatarAvatarTemplate string `json:"gravatar_avatar_template"`
	CustomAvatarUploadID   int    `json:"custom_avatar_upload_id"`
	CustomAvatarTemplate   string `json:"custom_avatar_template"`
	UploadedAvatarID       int    `json:"uploaded_avatar_id"`

	Groups     []*groups.Group     `json:"groups"`
	GroupUsers []*groups.GroupUser `json:"group_users"`

	UserOption UserOptions `json:"user_option"`

	UserAPIKeys *struct{} `json:"user_api_keys"`
}

type UserOptions struct {
	UserID         int  `json:"user_id"`
	DynamicFavicon bool `json:"dynamic_favicon"`

	EmailAlways          bool `json:"email_always"`
	EmailDigests         bool `json:"email_digets"`
	EmailPrivateMessages bool `json:"email_private_messages"`
	EmailDirect          bool `json:"email_direct"`
	EmailPreviousReplies int  `json:"email_previous_replies"`
	EmailInReplyTo       bool `json:"email_in_reply_to"`

	MailingListMode          bool `json:"mailing_list_mode"`
	MailingListModeFrequency int  `json:"mailing_list_mode_frequency"`

	DigestAfterMinutes  int  `json:"digest_after_minutes"`
	IncludeTL0InDigests bool `json:"include_tl0_in_digests"`

	DisableJumpReply      bool `json:"disable_jump_reply"`
	EnableQuoting         bool `json:"enable_quoting"`
	ExternalLinksInNewTab bool `json:"external_links_in_new_tab"`

	AutoTrackTopicsAfterMillisecs int  `json:"auto_track_topics_after_msecs"`
	NewTopicDurationMinutes       int  `json:"new_topic_duration_minutes"`
	AutomaticallyUnpinTopics      bool `json:"automatically_unpin_topics"`

	NotificationLevelWhenReplying int `json:"notification_level_when_replying"`
	LikeNotificationFrequency     int `json:"like_notification_frequency"`

	ThemeKey    string `json:"theme_key"`
	ThemeKeySeq int    `json:"theme_key_seq"`
}
