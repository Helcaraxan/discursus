package notifications

type ListRequest struct{}

type ListResponse struct {
	Notifications          []*Notification `json:"notifications"`
	TotalRowsNotifications int             `json:"total_rows_notifications"`
	SeenNotificationID     int             `json:"seen_notification_id"`
	LoadMoreNotifications  string          `json:"load_more_notifications"`
}

type Notification struct {
	ID               int    `json:"id"`
	NotificationType int    `json:"notification_type"`
	Read             bool   `json:"read"`
	CreatedAt        string `json:"created_at"`
	PostNumber       int    `json:"post_number"`
	TopicID          int    `json:"topic_id"`
	FancyTitle       string `json:"fancy_title"`
	Slug             string `json:"slug"`
	Data             Data   `json:"data"`
}

type Data struct {
	TopicTitle       string `json:"topic_title"`
	OriginalPostID   int    `json:"original_post_id"`
	OriginalPostType int    `json:"original_post_type"`
	OriginalUsername string `json:"original_username"`
	DisplayUsername  string `json:"display_username"`
}
