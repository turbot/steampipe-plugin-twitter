# Table: twitter_tweet

Get information about a public tweet.

To query lists of tweets see
[twitter_search_recent](../twitter_search_recent) or
[twitter_user_tweet](../twitter_user_tweet) instead.

Note: The `id` field must be set in the `where` clause.

## Examples

### Get tweet by ID

```sql
select
  *
from
  twitter_tweet
where
  id = '1373134228620214275'
```
