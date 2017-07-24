package atoms

type AutoCloseInfo struct {
	AutoCloseAt struct {
		Value string `json:"auto_close_at"`
	} `json:"auto_close_at"`
	AutoCloseHours struct {
		Value int `json:"auto_close_hours"`
	} `json:"auto_close_hours"`
	AutoCloseBasedOnLastPost bool `json:"auto_close_based_on_last_post"`
}

type CategoryID struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	DescriptionText    string `json:"description_text"`
	DescriptionExcerpt string `json:"description_excerpt"`
}

type CategoryInfo struct {
	Slug          string `json:"slug"`
	TopicURL      string `json:"topic_url"`
	LogoURL       string `json:"logo_url"`
	BackgroundURL string `json:"background_url"`
}

type CategoryStatistics struct {
	TopicCount    int `json:"topic_count"`
	PostCount     int `json:"post_count"`
	TopicsDay     int `json:"topics_day"`
	TopicsWeek    int `json:"topics_week"`
	TopicsMonth   int `json:"topics_month"`
	TopicsYear    int `json:"topics_year"`
	TopicsAllTime int `json:"topics_all_time"`
}

type DraftInfo struct {
	Draft         bool   `json:"draft"`
	DraftKey      string `json:"draft_key"`
	DraftSequence int    `json:"draft_sequence"`
}

type Logo struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type ModificationInfo struct {
	CreatedAt    string `json:"created_at"`
	Cooked       string `json:"cooked"`
	UpdatedAt    int    `json:"updated_at"`
	LastPostedAt string `json:"last_posted_at"`
	Bumped       bool   `json:"bumped"`
	BumpedAt     string `json:"bumped_at"`
	EditReason   string `json:"edit_reason"`
	DeletedAt    string `json:"deleted_at"`
	AvgTime      string `json:"avg_time"`
	Version      int    `json:"version"`
}

type PinnedInformation struct {
	Unpinned       bool   `json:"unpinned"`
	Pinned         bool   `json:"pinned"`
	PinnedGlobally bool   `json:"pinned_globally"`
	PinnedAt       string `json:"pinned_at"`
	PinnedUntil    string `json:"pinned_until"`
}

type Poster struct {
	Extras         string `json:"extras"`
	Description    string `json:"description"`
	UserID         int    `json:"user_id"`
	PrimaryGroupID int    `json:"primary_group_id"`
}

type PostID struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PostNumber        int    `json:"post_number"`
	PostType          int    `json:"post_type"`
	TrustLevel        int    `json:"trust_level"`
	ReplyToPostNumber int    `json:"reply_to_post_number"`
}

type PostStatistics struct {
	ReplyCount        int `json:"reply_count"`
	QuoteCount        int `json:"quote_count"`
	IncomingLinkCount int `json:"incoming_link_count"`
	Reads             int `json:"reads"`
	Score             int `json:"score"`
}

type PrimaryGroup struct {
	PrimaryGroupName         string `json:"primary_group_name"`
	PrimaryGroupFlairURL     string `json:"primary_group_flair_url"`
	PrimaryGroupFlairBGColor string `json:"primary_group_flair_bg_color"`
	PrimaryGroupFlairColor   string `json:"primary_group_flair_color"`
}

type PrimaryGroupReference struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	FlairURL     string `json:"flair_url"`
	FlairBGColor string `json:"flair_bg_color"`
	FlairColor   string `json:"flair_color"`
}

type TopicID struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	FancyTitle string `json:"fancy_title"`
	Excerpt    string `json:"excerpt"`
}

type TopicInformation struct {
	Slug      string `json:"slug"`
	ImageURL  string `json:"image_url"`
	Archetype string `json:"archetype"`
}

type TopicPermissions struct {
	CanMovePosts          bool `json:"can_move_posts"`
	CanEdit               bool `json:"can_edit"`
	CanDelete             bool `json:"can_delete"`
	CanRemoveAllowedUsers bool `json:"can_remove_allowed_users"`
	CanInviteTo           bool `json:"can_invite_to"`
	CanInviteViaEmail     bool `json:"can_invite_via_email"`
	CanCreatePost         bool `json:"can_create_post"`
	CanReplyAsNewTopic    bool `json:"can_reply_as_new_topic"`
	CanFlagTopic          bool `json:"can_flag_topic"`
}

type TopicReference struct {
	TopicID   int    `json:"topic_id"`
	TopicSlug string `json:"topic_slug"`
}

type TopicStatistics struct {
	WordCount         int `json:"word_count"`
	PostsCount        int `json:"posts_count"`
	ReplyCount        int `json:"reply_count"`
	HighestPostNumber int `json:"highest_post_number"`
	Views             int `json:"views"`
}

type UserID struct {
	UserID          int    `json:"user_id"`
	Username        string `json:"username"`
	UserTitle       string `json:"user_title"`
	DisplayUsername string `json:"display_username"`
	AvatarTemplate  string `json:"avatar_template"`
}

type UserPermissions struct {
	CanCreateCategory  bool `json:"can_create_category"`
	CanCreateTopic     bool `json:"can_create_topic"`
	CanDelete          bool `json:"can_delete"`
	CanEdit            bool `json:"can_edit"`
	CanRecover         bool `json:"can_recover"`
	CanViewEditHistory bool `json:"can_view_edit_history"`
	CanWiki            bool `json:"can_wiki"`
}

type UserReference struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	AvatarTemplate string `json:"avatar_template"`
}

type UserType struct {
	Moderator bool `json:"moderator"`
	Staff     bool `json:"staff"`
	Admin     bool `json:"admin"`
}
