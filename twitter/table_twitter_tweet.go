package twitter

import (
	"context"
	"errors"
	"fmt"
	"strings"

	twitter "github.com/g8rswimmer/go-twitter/v2"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableTwitterTweet(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twitter_tweet",
		Description: "Lookup a specific tweet by ID.",
		List: &plugin.ListConfig{
			Hydrate:    listTweet,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: tweetColumns(),
	}
}

func listTweet(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twitter_tweet.listTweet", "connection_error", err)
		return nil, err
	}
	opts := twitter.TweetLookupOpts{
		Expansions:  tweetExpansions(),
		TweetFields: tweetFields(),
		UserFields:  userFields(),
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	result, err := conn.TweetLookup(ctx, []string{id}, opts)
	// Hard error
	if err != nil {
		plugin.Logger(ctx).Error("twitter_tweet.listTweet", "query_error", err, "opts", opts)
		return nil, err
	}
	// Soft error, e.g. 404
	if len(result.Raw.Errors) > 0 {
		errMsgs := []string{}
		for _, e := range result.Raw.Errors {
			plugin.Logger(ctx).Error("twitter_user.listTweet", "query_error", e, "opts", opts)
			errMsgs = append(errMsgs, fmt.Sprintf("%s: %s", e.Title, e.Detail))
			if e.Title == "Not Found Error" {
				return nil, nil
			}
		}
		// Return the full set of error messages
		return nil, errors.New(strings.Join(errMsgs, "\n"))
	}
	for _, i := range result.Raw.TweetDictionaries() {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
