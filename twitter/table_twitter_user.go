package twitter

import (
	"context"
	"errors"
	"fmt"
	"strings"

	twitter "github.com/g8rswimmer/go-twitter/v2"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableTwitterUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twitter_user",
		Description: "Lookup a specific user by ID or username.",
		List: &plugin.ListConfig{
			Hydrate:    listUser,
			KeyColumns: plugin.AnyColumn([]string{"id", "username"}),
		},
		Columns: userColumns(),
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twitter_user.listUser", "connection_error", err)
		return nil, err
	}
	opts := twitter.UserLookupOpts{
		Expansions:  userExpansions(),
		UserFields:  userFields(),
		TweetFields: tweetFields(),
	}
	quals := d.KeyColumnQuals
	var result *twitter.UserLookupResponse
	var lookupErr error
	if quals["id"] == nil && quals["username"] != nil {
		un := quals["username"].GetStringValue()
		result, lookupErr = conn.UserNameLookup(ctx, []string{un}, opts)
	} else {
		id := quals["id"].GetStringValue()
		result, lookupErr = conn.UserLookup(ctx, []string{id}, opts)
	}
	// Hard error
	if lookupErr != nil {
		plugin.Logger(ctx).Error("twitter_user.listUser", "query_error", lookupErr, "opts", opts)
		return nil, lookupErr
	}
	// Soft error, e.g. 404
	if len(result.Raw.Errors) > 0 {
		errMsgs := []string{}
		for _, e := range result.Raw.Errors {
			plugin.Logger(ctx).Error("twitter_user.listUser", "query_error", e, "opts", opts)
			errMsgs = append(errMsgs, fmt.Sprintf("%s: %s", e.Title, e.Detail))
			if e.Title == "Not Found Error" {
				return nil, nil
			}
		}
		// Return the full set of error messages
		return nil, errors.New(strings.Join(errMsgs, "\n"))
	}
	for _, i := range result.Raw.UserDictionaries() {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
