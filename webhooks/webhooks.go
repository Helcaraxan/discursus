package webhooks

type CreateRequest struct {
	PayloadURL          string `json:"payload_url"`
	ContentType         int    `json:"content_type"`
	Secret              string `json:"secret"`
	WildcardWebhook     bool   `json:"wildcard_webhook"`
	VerifyCertificate   bool   `json:"verify_certificate"`
	Active              bool   `json:"active"`
	WebhookEventTypeIDs []int  `json:"webhook_event_types_ids"`
	CategoryIDs         []int  `json:"category_ids"`
	GroupIDs            []int  `json:"group_ids"`
}

type CreateResponse struct {
	ID                int    `json:"id"`
	PayloadURL        string `json:"payload_url"`
	ContentType       int    `json:"content_type"`
	Secret            string `json:"secret"`
	WildcardWebhook   bool   `json:"wildcard_webhook"`
	VerifyCertificate bool   `json:"verify_certificate"`
	Active            bool   `json:"active"`
	WebhookEventTypes []*struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"webhook_event_types_ids"`
	CategoryIDs []int `json:"category_ids"`
	GroupIDs    []int `json:"group_ids"`
}
