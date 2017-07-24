package messages

import (
	"github.com/helcaraxan/discursus/atoms"
)

type MessageCreateRequest struct {
	Title           string `json:"title"`
	Raw             string `json:"raw"`
	TargetUsernames string `json:"target_usernames"`
	Archetype       string `json:"archetype"`
	CreatedAt       string `json:"created_at"`
}

type MessageCreateResponse struct {
	Message
}

type Message struct {
	atoms.PostID
	atoms.TopicReference
	atoms.UserID
	atoms.UserType
	atoms.UserPermissions
	atoms.DraftInfo
	atoms.ModificationInfo
	atoms.PostStatistics
	atoms.PrimaryGroup

	MessageAnswerInfo

	Posters []*atoms.Poster `json:"posters"`

	UserDeleted bool `json:"user_deleted"`
	Wiki        bool `json:"wiki"`

	Yours          bool `json:"yours"`
	Hidden         bool `json:"hidden"`
	HiddenReasonID int  `json:"hidden_reason_id"`
}

type MessageAnswerInfo struct {
	CanAcceptAnswer   bool `json:"can_accept_answer"`
	CanUnacceptAnswer bool `json:"can_unaccept_answer"`
	AcceptedAnswer    bool `json:"accepted_answer"`
}
