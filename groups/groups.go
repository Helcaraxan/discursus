package groups

import (
	"github.com/helcaraxan/discursus/atoms"
)

type ListRequest struct{}

type ListResponse struct {
	Groups []*Group `json:"groups"`
	Extras struct {
		GroupUserIDs []int `json:"group_user_ids"`
	} `json:"extras"`
	TotalRowsGroups int    `json:"total_rows_groups"`
	LoadMoreGroups  string `json:"load_more_groups"`
}

type CreateRequest struct {
	GroupName string `json:"group[name]"`
}

type CreateResponse struct {
	BasicGroup Group `json:"basic_group"`
}

type DeleteRequest struct{}

type DeleteResponse struct {
	Success string `json:"success"`
}

type UpdateRequest struct {
	Group
}

type UpdateResponse struct {
	BasicGroup Group `json:"basic_group"`
}

type GetMembersRequest struct{}

type GetMembersResponse struct {
	Members []*Member `json:"members"`
	Owners  []*Member `json:"owners"`
	Meta    struct {
		Total  int `json:"total"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	} `json:"meta"`
}

type AddMembersRequest struct {
	Usernames string `json:"usernames"`
}

type AddMembersResponse struct {
	Success string `json:"success"`
}

type RemoveMemberRequest struct {
	UserID int `json:"user_id"`
}

type RemoveMemberResponse struct {
	Success string `json:"success"`
}

type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	FullName    string `json:"full_name"`
	Title       string `json:"title"`
	BioRaw      string `json:"bio_raw"`
	BioCooked   string `json:"bio_cooked"`

	IsMember    bool `json:"is_member"`
	Mentionable bool `json:"mentionable"`

	FlairURL     string `json:"flair_url"`
	FlairBGColor string `json:"flair_bg_color"`
	FlairColor   string `json:"flair_color"`

	UserCount       int `json:"user_count"`
	AliasLevel      int `json:"alias_level"`
	VisibilityLevel int `json:"visibility_level"`

	Automatic                       bool   `json:"automatic"`
	AutomaticMembershipEmailDomains string `json:"automatic_membership_email_domains"`
	AutomaticMembershipRetroactive  bool   `json:"automatic_membership_retroactive"`
	AllowMembershipRequests         bool   `json:"allow_membership_requests"`

	PrimaryGroup             bool   `json:"primary_group"`
	Public                   bool   `json:"public"`
	GrantTrustLevel          int    `json:"grant_trust_level"`
	IncomingEmail            string `json:"incoming_email"`
	HasMessages              bool   `json:"has_messages"`
	DefaultNotificationLevel int    `json:"default_notification_level"`
}

type GroupUser struct {
	GroupID           int  `json:"group_id"`
	UserID            int  `json:"user_id"`
	NotificationLevel int  `json:"notification_level"`
	Owner             bool `json:"owner"`
}

type Member struct {
	atoms.UserReference

	Title        string `json:"title"`
	LastPostedAt string `json:"last_posted_at"`
	LastSeenAt   string `json:"last_seen_at"`
}
