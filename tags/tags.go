package tags

import (
	"github.com/helcaraxan/discursus/atoms"
	"github.com/helcaraxan/discursus/topics"
)

type GetTagRequest struct{}

type GetTagResponse struct {
	Users         []*atoms.UserReference         `json:"users"`
	PrimaryGroups []*atoms.PrimaryGroupReference `json:"primary_groups"`
	TopicList     topics.TopicList               `json:"topic_list"`
}
