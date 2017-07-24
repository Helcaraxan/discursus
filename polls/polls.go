package polls

type VoteRequest struct {
	PostID   int    `json:"post_id"`
	PollName string `json:"poll_name"`
	Options  string `json:"options[]"`
}

type VoteResponse struct {
	Poll   []*PollOption `json:"poll"`
	Voters int           `json:"voters"`
	Status string        `json:"status"`
	Name   string        `json:"name"`
	Type   string        `json:"type"`
}

type PollOption struct {
	ID    string `json:"id"`
	HTML  string `json:"html"`
	Votes int    `json:"votes"`
}
