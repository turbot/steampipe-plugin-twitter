package twitter

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type twitterConfig struct {
	// OAuth 2.0
	BearerToken *string `hcl:"bearer_token"`
	// OAuth 1.0
	ConsumerKey    *string `hcl:"consumer_key"`
	ConsumerSecret *string `hcl:"consumer_secret"`
	AccessToken    *string `hcl:"access_token"`
	AccessSecret   *string `hcl:"access_secret"`
	// Limits
	MaxItemsPerQuery *int `hcl:"max_items_per_query"`
}

func ConfigInstance() interface{} {
	return &twitterConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) twitterConfig {
	if connection == nil || connection.Config == nil {
		return twitterConfig{}
	}
	config, _ := connection.Config.(twitterConfig)
	return config
}
