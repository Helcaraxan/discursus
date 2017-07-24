package categories

import (
	"github.com/helcaraxan/discursus/atoms"
	"github.com/helcaraxan/discursus/groups"
	"github.com/helcaraxan/discursus/topics"
	"github.com/helcaraxan/discursus/users"
)

type ListRequest struct{}

type ListResponse struct {
	atoms.UserPermissions
	atoms.DraftInfo
	Categories []*Category `json:"categories"`
}

type CreateRequest struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
	TextColor string `json:"text_color"`
}

type CreateResponse struct {
	Category
}

type ListTopicsRequest struct{}

type ListTopicsResponse struct {
	Users     []*users.User    `json:"users"`
	TopicList topics.TopicList `json:"topic_list"`
}

type UpdateRequest CreateRequest

type UpdateResponse struct {
	Success  string   `json:"success"`
	Category Category `json:"category"`
}

type Category struct {
	atoms.CategoryID
	atoms.CategoryInfo
	atoms.CategoryStatistics
	atoms.UserPermissions

	Color             string `json:"color"`
	TextColor         string `json:"text_color"`
	DefaultView       string `json:"default_view"`
	NotificationLevel string `json:"notification_level"`

	UploadedLogo       atoms.Logo `json:"uploaded_logo"`
	UploadedBackground string     `json:"uploaded_background"`

	SortOrder            string `json:"sort_order"`
	SortAscending        string `json:"sort_ascending"`
	Position             int    `json:"position"`
	HasChildren          bool   `json:"has_children"`
	ShowSubcategoryList  bool   `json:"show_subcategory_list"`
	SubcategoryListStyle string `json:"subcategory_list_style"`
	NumFeaturedTopics    int    `json:"num_featured_topics"`
	DefaultTopPeriod     string `json:"default_top_period"`
	SuppressFromHomepage bool   `json:"suppress_from_homepage"`

	TopicTemplate string `json:"topic_template"`

	CannotDeleteReason    string `json:"cannot_delete_reason"`
	Permission            int    `json:"permission"`
	ReadRestricted        bool   `json:"read_restricted"`
	AllowBadges           bool   `json:"allow_badges"`
	EmailIn               bool   `json:"email_in"`
	EmailInAllowStrangers bool   `json:"email_in_allow_strangers"`

	AutoCloseHours           int  `json:"auto_close_hours"`
	AutoCloseBasedOnLastPost bool `json:"auto_close_based_on_last_past"`

	AvailableGroups  []*groups.Group `json:"available_groups"`
	GroupPermissions []interface{}   `json:"group_permissions"`

	CustomFields interface{} `json:"custom_fields"`
}
