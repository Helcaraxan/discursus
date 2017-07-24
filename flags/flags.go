package flags

import (
	"github.com/helcaraxan/discursus/atoms"
	"github.com/helcaraxan/discursus/messages"
	"github.com/helcaraxan/discursus/topics"
)

type ListRequest struct{}

type ListResponse struct {
	Posts  []*messages.Message    `json:"posts"`
	Topics []*topics.Topic        `json:"topics"`
	Users  []*atoms.UserReference `json:"users"`
}
