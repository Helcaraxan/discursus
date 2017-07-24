package badges

type UserListRequest struct{}

type UserListResponse struct {
	BadgeTypes []*UserBadge
}

type UserBadge struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	BadgeID int `json:"badge_id"`

	Count       int    `json:"count"`
	GrantedAt   string `json:"granted_at"`
	GrantedByID int    `json:"granted_by_id"`

	PostNumber int `json:"post_number"`
	PostID     int `json:"post_id"`
	TopicID    int `json:"topic_id"`
}

type Badge struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Image       *struct{} `json:"image,omitempty"`

	BadgeGroupingID int `json:"badge_grouping_id"`
	BadgeTypeID     int `json:"badge_type_id"`

	GrantCount int `json:"grant_count"`

	AllowTitle    bool `json:"allow_title"`
	MultipleGrant bool `json:"multiple_grant"`
	Listable      bool `json:"listable"`
	Enabled       bool `json:"enabled"`
	System        bool `json:"system"`
}

type BadgeType struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}
