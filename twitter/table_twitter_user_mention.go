package twitter

import (
	"context"

	twitter "github.com/g8rswimmer/go-twitter/v2"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableTwitterUserMention(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twitter_user_mention",
		Description: "The user mention timeline lists Tweets mentioning a specific Twitter user, for example, if a Twitter account mentioned @TwitterDev within a Tweet. This will also include replies to Tweets by the user requested.",
		List: &plugin.ListConfig{
			Hydrate:    listUserMention,
			KeyColumns: plugin.SingleColumn("user_id"),
		},
		Columns: tweetColumns("user_id"),
	}
}

func listUserMention(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twitter_user_mention.listUserMention", "connection_error", err)
		return nil, err
	}
	maxItems := maxItemsPerQuery(ctx, d)
	opts := twitter.UserMentionTimelineOpts{
		Expansions:  tweetExpansions(),
		TweetFields: tweetFields(),
		UserFields:  userFields(),
		MaxResults:  min(100, maxItems),
	}

	keyQuals := d.EqualsQuals
	userID := keyQuals["user_id"].GetStringValue()

	/*
		// TODO - This works, but only when STEAMPIPE_CACHE=false.
		quals := d.QueryContext.GetQuals()
		if quals["created_at"] != nil {
			for _, q := range quals["created_at"].Quals {
				ts, e := ptypes.Timestamp(q.Value.GetTimestampValue())
				if e != nil {
					plugin.Logger(ctx).Error("twitter_user_mention.listUserMention", "parse_error", e, "userID", userID, "opts", opts)
					continue
				}
				switch q.GetStringValue() {
				case ">", ">=":
					opts.StartTime = ts
				case "<", "<=":
					opts.EndTime = ts
				}
			}
		}
	*/

	count := 0
	for {
		result, err := conn.UserMentionTimeline(ctx, userID, opts)
		if err != nil {
			plugin.Logger(ctx).Error("twitter_user_mention.listUserMention", "query_error", err, "userID", userID, "opts", opts)
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
		opts.PaginationToken = result.Meta.NextToken
	}
	return nil, nil
}
