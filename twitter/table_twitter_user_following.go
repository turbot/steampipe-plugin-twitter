package twitter

import (
	"context"

	twitter "github.com/g8rswimmer/go-twitter/v2"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableTwitterUserFollowing(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twitter_user_following",
		Description: "List of users the specified user ID is following.",
		List: &plugin.ListConfig{
			Hydrate:    listUserFollowing,
			KeyColumns: plugin.SingleColumn("user_id"),
		},
		Columns: append(
			userColumns("user_id"),
		),
	}
}

func listUserFollowing(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twitter_search.listUserFollowing", "connection_error", err)
		return nil, err
	}
	maxItems := maxItemsPerQuery(ctx, d)
	opts := twitter.UserFollowingLookupOpts{
		Expansions:  userExpansions(),
		UserFields:  userFields(),
		TweetFields: tweetFields(),
		MaxResults:  min(100, maxItems),
	}
	keyQuals := d.KeyColumnQuals
	userID := keyQuals["user_id"].GetStringValue()
	count := 0
	for {
		result, err := conn.UserFollowingLookup(ctx, userID, opts)
		if err != nil {
			plugin.Logger(ctx).Error("twitter_search.listUserFollowing", "query_error", err, "userID", userID, "opts", opts)
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
