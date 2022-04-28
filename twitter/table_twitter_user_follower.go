package twitter

import (
	"context"

	twitter "github.com/g8rswimmer/go-twitter/v2"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableTwitterUserFollower(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twitter_user_follower",
		Description: "List of users the specified user ID is follower.",
		List: &plugin.ListConfig{
			Hydrate:    listUserFollower,
			KeyColumns: plugin.SingleColumn("user_id"),
		},
		Columns: userColumns("user_id"),
	}
}

func listUserFollower(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twitter_search.listUserFollower", "connection_error", err)
		return nil, err
	}
	maxItems := maxItemsPerQuery(ctx, d)
	opts := twitter.UserFollowersLookupOpts{
		Expansions:  userExpansions(),
		UserFields:  userFields(),
		TweetFields: tweetFields(),
		MaxResults:  min(100, maxItems),
	}
	keyQuals := d.KeyColumnQuals
	userID := keyQuals["user_id"].GetStringValue()
	count := 0
	for {
		result, err := conn.UserFollowersLookup(ctx, userID, opts)
		if err != nil {
			plugin.Logger(ctx).Error("twitter_search.listUserFollower", "query_error", err, "userID", userID, "opts", opts)
			return nil, err
		}
		for _, i := range result.Raw.UserDictionaries() {
			d.StreamListItem(ctx, i)
			count++
		}
		// Only check the max items after each page, we've already taken the cost
		// of getting the page of results anyway so no point in throwing them away
		if maxItems >= 0 && count >= maxItems {
			break
		}
		// Is there another page?
		if result.Meta.NextToken == "" {
			break
		}
		opts.PaginationToken = result.Meta.NextToken
	}
	return nil, nil
}
