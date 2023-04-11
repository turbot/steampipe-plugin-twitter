package twitter

import (
	"context"

	twitter "github.com/g8rswimmer/go-twitter/v2"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableTwitterSearchRecent(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twitter_search_recent",
		Description: "Search public Tweets posted over the last 7 days.",
		List: &plugin.ListConfig{
			Hydrate:    listSearchRecent,
			KeyColumns: plugin.SingleColumn("query"),
		},
		Columns: tweetColumns("query"),
	}
}

func listSearchRecent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twitter_search.listSearchRecent", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	q := quals["query"].GetStringValue()
	maxItems := maxItemsPerQuery(ctx, d)
	opts := twitter.TweetRecentSearchOpts{
		Expansions:  tweetExpansions(),
		TweetFields: tweetFields(),
		UserFields:  userFields(),
		MaxResults:  min(100, maxItems),
	}
	count := 0
	for {
		result, err := conn.TweetRecentSearch(ctx, q, opts)
		if err != nil {
			plugin.Logger(ctx).Error("twitter_search.listSearchRecent", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, i := range result.Raw.TweetDictionaries() {
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
		opts.NextToken = result.Meta.NextToken
	}
	return nil, nil
}
