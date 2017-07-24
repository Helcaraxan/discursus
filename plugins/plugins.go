package plugins

type PluginsGetListRequest struct{}

type PluginsGetListResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Version        string `json:"version"`
	URL            string `json:"url"`
	Enabled        bool   `json:"enabled"`
	EnabledSetting string `json:"enabled_setting"`
}
