package twitter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	twitter "github.com/g8rswimmer/go-twitter/v2"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

type authorizeOAuth1 struct{}

func (a authorizeOAuth1) Add(req *http.Request) {}

type authorizeOAuth2 struct {
	Token string
}

func (a authorizeOAuth2) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func connect(_ context.Context, d *plugin.QueryData) (*twitter.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "twitter"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*twitter.Client), nil
	}

	bearerToken := os.Getenv("TWITTER_BEARER_TOKEN")
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	// First, use the token config
	twitterConfig := GetConfig(d.Connection)
	if &twitterConfig != nil {
		if twitterConfig.BearerToken != nil {
			bearerToken = *twitterConfig.BearerToken
		}
		if twitterConfig.ConsumerKey != nil {
			consumerKey = *twitterConfig.ConsumerKey
		}
		if twitterConfig.ConsumerSecret != nil {
			consumerSecret = *twitterConfig.ConsumerSecret
		}
		if twitterConfig.AccessToken != nil {
			accessToken = *twitterConfig.AccessToken
		}
		if twitterConfig.AccessSecret != nil {
			accessSecret = *twitterConfig.AccessSecret
		}
	}

	var conn *twitter.Client

	// First, try to use the bearer token and OAuth 2.0
	if bearerToken != "" {
		conn = &twitter.Client{
			Authorizer: authorizeOAuth2{
				Token: bearerToken,
			},
			Client: http.DefaultClient,
			Host:   "https://api.twitter.com",
		}
	} else if consumerKey != "" {
		config := oauth1.NewConfig(consumerKey, consumerSecret)
		token := oauth1.NewToken(accessToken, accessSecret)
		httpClient := config.Client(oauth1.NoContext, token)
		conn = &twitter.Client{
			Authorizer: authorizeOAuth1{},
			Client:     httpClient,
			Host:       "https://api.twitter.com",
		}
	} else {
		// Credentials not set
		return nil, errors.New("bearer_token (or consumer_key etc) must be configured")
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}

func maxItemsPerQuery(ctx context.Context, d *plugin.QueryData) int {
	// First, use the config
	twitterConfig := GetConfig(d.Connection)
	if &twitterConfig != nil {
		if twitterConfig.MaxItemsPerQuery != nil {
			return *twitterConfig.MaxItemsPerQuery
		}
	}
	// Second, return a default
	return 1000
}

func queryString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	q := quals["query"].GetStringValue()
	return q, nil
}

func userIDString(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	i := quals["user_id"].GetStringValue()
	return i, nil
}

func referenced(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	refType := d.Param.(string)
	t := d.Value.(twitter.TweetObj)
	for _, ref := range t.ReferencedTweets {
		if ref.Type == refType {
			return ref.ID, nil
		}
	}
	return nil, nil
}

func userColumns(optionalCols ...string) []*plugin.Column {
	cols := []*plugin.Column{
		// Top columns
		{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ID"), Description: "The unique identifier of this user."},
		{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Name"), Description: "The name of the user, as they’ve defined it on their profile. Not necessarily a person’s name."},
		{Name: "username", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.UserName"), Description: "The Twitter screen name, handle, or alias that this user identifies themselves with. Usernames are unique but subject to change."},
		// Other columns
		{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("User.CreatedAt"), Description: "The UTC datetime that the user account was created on Twitter."},
		{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Description"), Description: "The text of this user's profile description (also known as bio), if the user provided one."},
		{Name: "entities", Type: proto.ColumnType_JSON, Transform: transform.FromField("User.Entities"), Description: "Entities are JSON objects that provide additional information about hashtags, urls, user mentions, and cashtags associated with the description."},
		{Name: "location", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Location"), Description: "The location specified in the user's profile, if the user provided one. As this is a freeform value, it may not indicate a valid location, but it may be fuzzily evaluated when performing searches with location queries."},
		{Name: "pinned_tweet", Type: proto.ColumnType_JSON, Transform: transform.FromField("PinnedTweet"), Description: "Contains withholding details for withheld content, if applicable."},
		{Name: "pinned_tweet_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.PinnedTweetID"), Description: "Unique identifier of this user's pinned Tweet."},
		{Name: "profile_image_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.ProfileImageURL"), Description: "The URL to the profile image for this user, as shown on the user's profile."},
		{Name: "protected", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.Protected"), Description: "Indicates if this user has chosen to protect their Tweets (in other words, if this user's Tweets are private)."},
		{Name: "public_metrics", Type: proto.ColumnType_JSON, Transform: transform.FromField("User.PublicMetrics"), Description: "Contains details about activity for this user."},
		{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("User.URL"), Description: "The URL specified in the user's profile, if present."},
		{Name: "verified", Type: proto.ColumnType_BOOL, Transform: transform.FromField("User.Verified"), Description: "Indicates if this user is a verified Twitter User."},
		{Name: "withheld", Type: proto.ColumnType_JSON, Transform: transform.FromField("User.WithHeld"), Description: "Contains withholding details for withheld content, if applicable."},
	}
	for _, col := range optionalCols {
		switch col {
		case "user_id":
			cols = append(cols, &plugin.Column{Name: "user_id", Type: proto.ColumnType_STRING, Hydrate: userIDString, Transform: transform.FromValue(), Description: "ID of the user who is followed by these users."})
		}
	}
	return cols
}

func tweetColumns(optionalCols ...string) []*plugin.Column {
	cols := []*plugin.Column{
		// Top columns
		{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet.ID"), Description: "Unique identifier of this Tweet."},
		{Name: "text", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet.Text"), Description: "The content of the Tweet."},
		// Other columns
		{Name: "author_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet.AuthorID"), Description: "Unique identifier of the author of the Tweet."},
		{Name: "conversation_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet.ConversationID"), Description: "The Tweet ID of the original Tweet of the conversation (which includes direct replies, replies of replies)."},
		{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Tweet.CreatedAt"), Description: "Creation time of the Tweet."},

		{Name: "in_reply_to_user_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet.InReplyToUserID"), Description: "If this Tweet is a Reply, indicates the user ID of the parent Tweet's author."},
		// Simple, custom fields for references.
		// TODO - Check that there can only be one of each type, which is unclear to me.
		{Name: "replied_to", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet").TransformP(referenced, "replied_to"), Description: "If this Tweet is a Reply, indicates the ID of the Tweet it is a reply to."},
		{Name: "retweeted", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet").TransformP(referenced, "retweeted"), Description: "If this Tweet is a Retweet, indicates the ID of the orginal Tweet."},
		{Name: "quoted", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet").TransformP(referenced, "quoted"), Description: "If this Tweet is a Quote Tweet, indicates the ID of the original Tweet."},
		// Object form for the simple fields, removed for now.
		//{Name: "referenced_tweets", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.WithHeld"), Description: "A list of Tweets this Tweet refers to. For example, if the parent Tweet is a Retweet, a Retweet with comment (also known as Quoted Tweet) or a Reply, it will include the related Tweet referenced to by its parent."},

		{Name: "mentions", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet").Transform(mentions), Description: "List of users (e.g. steampipeio) mentioned in the Tweet."},
		{Name: "hashtags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet").Transform(hashtags), Description: "List of hashtags (e.g. #sql) mentioned in the Tweet."},
		{Name: "urls", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet").Transform(urls), Description: "List of URLs (e.g. https://steampipe.io) mentioned in the Tweet."},
		{Name: "cashtags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet").Transform(cashtags), Description: "List of cashtags (e.g. $TWTR) mentioned in the Tweet."},
		{Name: "entities", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.Entities"), Description: "Contains details about text that has a special meaning in a Tweet."},

		{Name: "attachments", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.Attachments"), Description: "Specifies the type of attachments (if any) present in this Tweet."},
		{Name: "geo", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.Geo"), Description: "Contains details about the location tagged by the user in this Tweet, if they specified one."},
		{Name: "context_annotations", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.ContextAnnotations"), Description: "Contains context annotations for the Tweet."},
		{Name: "withheld", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.Withheld"), Description: "Contains withholding details for withheld content."},
		{Name: "public_metrics", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.PublicMetrics"), Description: "Engagement metrics for the Tweet at the time of the request."},

		// These fields require extra permissions, so hidden for now
		//{Name: "non_public_metrics", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.NonPublicMetrics"), Description: "Non-public engagement metrics for the Tweet at the time of the request. This is a private metric, and requires the use of OAuth 1.0a User Context authentication."},
		//{Name: "organic_metrics", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.OrganicMetrics"), Description: "Organic engagement metrics for the Tweet at the time of the request. Requires user context authentication."},
		//{Name: "promoted_metrics", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.PromotedMetrics"), Description: "Engagement metrics for the Tweet at the time of the request in a promoted context. Requires user context authentication."},

		{Name: "possibly_sensitive", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Tweet.PossibySensitive"), Description: "Indicates if this Tweet contains URLs marked as sensitive, for example content suitable for mature audiences."},
		{Name: "lang", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet.Language"), Description: "Language of the Tweet, if detected by Twitter. Returned as a BCP47 language tag."},
		//{Name: "reply_settings", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet.ReplySettings"), Description: "Shows who can reply to this Tweet. Fields returned are everyone, mentionedUsers, and following."},
		{Name: "source", Type: proto.ColumnType_STRING, Transform: transform.FromField("Tweet.Source"), Description: "The name of the app the user Tweeted from."},
		//{Name: "includes", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tweet.Includes"), Description: "If you include an expansion parameter, the referenced objects will be returned if available."},

		{Name: "author", Type: proto.ColumnType_JSON, Transform: transform.FromField("Author"), Description: "Author of the Tweet."},
		{Name: "in_reply_user", Type: proto.ColumnType_JSON, Transform: transform.FromField("InReplyUser"), Description: "User the Tweet was in reply to."},
		{Name: "place", Type: proto.ColumnType_JSON, Transform: transform.FromField("Place"), Description: "Place where the Tweet was created."},
		{Name: "attachment_polls", Type: proto.ColumnType_JSON, Transform: transform.FromField("AttachmentPolls"), Description: "Polls attached to the Tweet."},
		{Name: "mentions_obj", Type: proto.ColumnType_JSON, Transform: transform.FromField("Mentions"), Description: "Users mentioned in the Tweet."},
		{Name: "referenced_tweets", Type: proto.ColumnType_JSON, Transform: transform.FromField("ReferencedTweets"), Description: "Tweets referenced in this Tweet."},
	}
	for _, col := range optionalCols {
		switch col {
		case "user_id":
			cols = append(cols, &plugin.Column{Name: "user_id", Type: proto.ColumnType_STRING, Hydrate: userIDString, Transform: transform.FromValue(), Description: "ID of the user the tweets are related to."})
		case "query":
			cols = append(cols, &plugin.Column{Name: "query", Type: proto.ColumnType_STRING, Hydrate: queryString, Transform: transform.FromValue(), Description: "Query string for the exploit search."})
		}
	}
	return cols
}

func userExpansions() []twitter.Expansion {
	return []twitter.Expansion{
		twitter.ExpansionPinnedTweetID,
	}
}

func tweetExpansions() []twitter.Expansion {
	return []twitter.Expansion{
		twitter.ExpansionAttachmentsPollIDs,
		twitter.ExpansionAttachmentsMediaKeys,
		twitter.ExpansionAuthorID,
		twitter.ExpansionEntitiesMentionsUserName,
		twitter.ExpansionGeoPlaceID,
		twitter.ExpansionInReplyToUserID,
		twitter.ExpansionReferencedTweetsID,
		twitter.ExpansionReferencedTweetsIDAuthorID,
		//twitter.ExpansionPinnedTweetID,
	}
}

func tweetFields() []twitter.TweetField {
	return []twitter.TweetField{
		twitter.TweetFieldID,
		twitter.TweetFieldText,
		twitter.TweetFieldAttachments,
		twitter.TweetFieldAuthorID,
		twitter.TweetFieldContextAnnotations,
		twitter.TweetFieldConversationID,
		twitter.TweetFieldCreatedAt,
		twitter.TweetFieldEntities,
		twitter.TweetFieldGeo,
		twitter.TweetFieldInReplyToUserID,
		twitter.TweetFieldLanguage,
		//twitter.TweetFieldNonPublicMetrics,
		twitter.TweetFieldPublicMetrics,
		//twitter.TweetFieldOrganicMetrics,
		//twitter.TweetFieldPromotedMetrics,
		twitter.TweetFieldPossiblySensitve,
		twitter.TweetFieldReferencedTweets,
		twitter.TweetFieldSource,
		twitter.TweetFieldWithHeld,
	}
}

func userFields() []twitter.UserField {
	return []twitter.UserField{
		twitter.UserFieldCreatedAt,
		twitter.UserFieldDescription,
		twitter.UserFieldEntities,
		twitter.UserFieldID,
		twitter.UserFieldLocation,
		twitter.UserFieldName,
		twitter.UserFieldPinnedTweetID,
		twitter.UserFieldProfileImageURL,
		twitter.UserFieldProtected,
		twitter.UserFieldPublicMetrics,
		twitter.UserFieldURL,
		twitter.UserFieldUserName,
		twitter.UserFieldVerified,
		twitter.UserFieldWithHeld,
	}
}

func mentions(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	t := d.Value.(twitter.TweetObj)
	items := []string{}
	if t.Entities != nil {
		for _, i := range t.Entities.Mentions {
			items = append(items, i.UserName)
		}
	}
	return items, nil
}

func urls(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	t := d.Value.(twitter.TweetObj)
	items := []string{}
	if t.Entities != nil {
		for _, i := range t.Entities.URLs {
			items = append(items, i.URL)
		}
	}
	return items, nil
}

func hashtags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	t := d.Value.(twitter.TweetObj)
	items := []string{}
	if t.Entities != nil {
		for _, i := range t.Entities.HashTags {
			items = append(items, i.Tag)
		}
	}
	return items, nil
}

func cashtags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	t := d.Value.(twitter.TweetObj)
	items := []string{}
	if t.Entities != nil {
		for _, i := range t.Entities.CashTags {
			items = append(items, i.Tag)
		}
	}
	return items, nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
