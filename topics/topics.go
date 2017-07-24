package topics

import (
	"github.com/helcaraxan/discursus/atoms"
	"github.com/helcaraxan/discursus/messages"
)

type TopicCreateRequest struct {
	Title     string `json:"title"`
	Raw       string `json:"raw"`
	Category  int    `json:"category"`
	CreatedAt string `json:"created_at"`
}

type TopicCreateResponse struct {
	messages.Message
}

type TopicGetRequest struct{}

type TopicGetResponse struct {
	PostStream struct {
		Posts  []*messages.Message `json:"posts"`
		Stream []int               `json:"stream"`
	} `json:"post_stream"`
	TimelineLookup []*[]int `json:"timeline_lookup"`
	Topic
}

type TopicDeleteRequest struct{}

type TopicDeleteResponse struct{}

type TopicUpdateRequest struct {
	Title      string `json:"title"`
	CategoryID int    `json:"category_id"`
}

type TopicUpdateResponse struct {
	BasicTopic struct {
		ID         int    `json:"id"`
		Title      string `json:"title"`
		FancyTitle string `json:"fancy_title"`
		Slug       string `json:"slug"`
		PostsCount int    `json:"posts_count"`
	} `json:"basic_topic"`
}

type TopicInviteUserRequest struct {
	Username string `json:"username"`
}

type TopicInviteUserResponse struct {
	User atoms.UserReference `json:"user"`
}

type TopicBookmarkRequest struct{}

type TopicBookmarkResponse struct{}

type TopicCloseRequest struct {
	Enabled bool   `json:"enabled"`
	Status  string `json:"status"`
}

type TopicCloseResponse struct {
	Success           string   `json:"success"`
	TopicStatusUpdate struct{} `json:"topic_status_update"`
}

type TopicGetLatestRequest struct{}

type TopicGetLatestResponse struct {
	Users     []*atoms.UserReference `json:"users"`
	TopicList TopicList              `json:"topic_list"`
}

type TopicGetTopRequest struct{}

type TopicGetTopResponse TopicGetLatestResponse

type TopicCreateTimedRequest struct {
	Time           string `json:"time"`
	TimezoneOffset int    `json:"timezone_offset"`
	StatusType     string `json:"status_type"`
	CategoryID     int    `json:"category_id"`
}

type TopicCreateTimedResponse struct{}

type TopicUpdateTimestampRequest struct {
	Timestamp int `json:"timestamp"`
}

type TopicUpdateTimestampResponse struct {
	Success string `json:"success"`
}

type TopicSetNotificationLevelRequest struct {
	NotificationLevel int `json:"notification_level"`
}

type TopicSetNotificationLevelResponse TopicUpdateTimestampResponse

type TopicList struct {
	atoms.UserPermissions
	atoms.DraftInfo
	PerPage int      `json:"per_page"`
	Tags    []string `json:"tags"`
	Topics  []*Topic `json:"topics"`
}

type Topic struct {
	atoms.TopicID
	atoms.TopicInformation
	atoms.TopicStatistics
	atoms.ModificationInfo
	atoms.PinnedInformation

	TopicAnswerInfo

	Details Details `json:"details"`

	Tags       []string `json:"tags"`
	HasSummary bool     `json:"has_summary"`
	HasDeleted bool     `json:"has_deleted"`

	ActionsSummary []*Action `json:"actions_summary"`

	UserID             int             `json:"user_id"`
	Posters            []*atoms.Poster `json:"posters"`
	ParticipantCount   int             `json:"participant_count"`
	LastPosterUsername string          `json:"last_poster_username"`

	Visible            bool `json:"visible"`
	Unseen             bool `json:"unseen"`
	Unread             int  `json:"unread"`
	NewPosts           int  `json:"new_posts"`
	LastReadPostID     int  `json:"last_read_post_id"`
	LastReadPostNumber int  `json:"last_read_post_number"`
	NotificationLevel  int  `json:"notification_level"`

	TopicTimer       *struct{} `json:"topic_timer"`
	MessageBusLastID int       `json:"message_bus_last_id"`
	ChunkSize        int       `json:"chunk_size"`

	Bookmarked bool `json:"bookmarked"`
	Closed     bool `json:"close"`
	Archived   bool `json:"archived"`

	Liked     bool `json:"liked"`
	LikeCount int  `json:"like_count"`

	FeaturedLink string `json:"featured_link"`

	CanVote   bool `json:"can_vote"`
	UserVoted bool `json:"user_voted"`
	VoteCount int  `json:"vote_count"`
}

type TopicAnswerInfo struct {
	CanHaveAnswer     bool `json:"can_have_answer"`
	HasAcceptedAnswer bool `json:"has_accepted_answer"`
}
type Details struct {
	atoms.AutoCloseInfo
	atoms.TopicPermissions

	CreatedBy            atoms.UserReference `json:"created_by"`
	LastPoster           atoms.UserReference `json:"last_poster"`
	Participants         []*Participant      `json:"participants"`
	SuggestedTopics      []*Topic            `json:"suggested_topics"`
	Links                []*Link             `json:"links"`
	NotificationLevel    int                 `json:"notification_level"`
	NoticicationReasonID int                 `json:"notification_reason_id"`
}

type Participant struct {
	atoms.UserReference
	atoms.PrimaryGroup

	PostCount int `json:"post_count"`
}

type Link struct {
	URL        string `json:"url"`
	Title      string `json:"title"`
	FancyTitle string `json:"fancy_title"`
	Internal   bool   `json:"internal"`
	Attachment bool   `json:"attachment"`
	Reflection bool   `json:"reflection"`
	Clicks     int    `json:"clicks"`
	UserID     int    `json:"user_id"`
	Domain     string `json:"domain"`
}

type Action struct {
	ID     int  `json:"id"`
	Count  int  `json:"count"`
	Hidden bool `json:"hidden"`
	CanAct bool `json:"can_act"`
}
