package twitter

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-twitter",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"twitter_search_recent":  tableTwitterSearchRecent(ctx),
			"twitter_tweet":          tableTwitterTweet(ctx),
			"twitter_user":           tableTwitterUser(ctx),
			"twitter_user_follower":  tableTwitterUserFollower(ctx),
			"twitter_user_following": tableTwitterUserFollowing(ctx),
			"twitter_user_mention":   tableTwitterUserMention(ctx),
			"twitter_user_tweet":     tableTwitterUserTweet(ctx),
		},
	}
	return p
}
