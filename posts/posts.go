package posts

import "github.com/helcaraxan/discursus/messages"

type PostCreateRequest struct {
	TopicID   int    `json:"topic_id"`
	Raw       string `json:"raw"`
	CreatedAt string `json:"created_atg"`
}

type PostCreateResponse struct {
	messages.Message
}

type PostGetRequest struct{}

type PostGetResponse PostCreateResponse

type PostUpdateRequest struct {
	RawPost string `json:"post[raw]"`
}

type PostUpdateResponse PostCreateResponse

type PostLikeRequest struct {
	ID               int  `json:"id"`
	PostActionTypeID int  `json:"post_action_type_id"`
	FlagTopic        bool `json:"flag_topic"`
}

type PostLikeResponse struct{}

type PostUnlikeRequest struct {
	PostActionTypeID int `json:"post_action_type_id"`
}

type PostUnlikeResponse struct{}
