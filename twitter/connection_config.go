package twitter

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type twitterConfig struct {
	// OAuth 2.0
	BearerToken *string `cty:"bearer_token"`
	// OAuth 1.0
	ConsumerKey    *string `cty:"consumer_key"`
	ConsumerSecret *string `cty:"consumer_secret"`
	AccessToken    *string `cty:"access_token"`
	AccessSecret   *string `cty:"access_secret"`
	// Limits
	MaxItemsPerQuery *int `cty:"max_items_per_query"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"bearer_token": {
		Type: schema.TypeString,
	},
	"consumer_key": {
		Type: schema.TypeString,
	},
	"consumer_secret": {
		Type: schema.TypeString,
	},
	"access_token": {
		Type: schema.TypeString,
	},
	"access_secret": {
		Type: schema.TypeString,
	},
	"max_items_per_query": {
		Type: schema.TypeInt,
	},
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
