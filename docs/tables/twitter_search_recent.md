# Table: twitter_search

Search the last 7 days of public tweets using the [Twitter search syntax and operators](https://developer.twitter.com/en/docs/twitter-api/tweets/search/integrate/build-a-query).

Notes:
* The `query` field must be set in the `where` clause.
* To prevent excess API quota use, results are limited to `max_results_per_query` by default.

## Examples

### Tweets matching a hashtag

```sql
select
  *
from
  twitter_search
where
  query = '#rowscoloredglasses'
```

### Tweets mentioning a user

```sql
select
  *
from
  twitter_search
where
  query = '@steampipeio'
```

### Tweets by a given author

```sql
select
  *
from
  twitter_search
where
  query = 'from:steampipeio'
```

### Tweets in reply to a user

```sql
select
  *
from
  twitter_search
where
  query = 'to:steampipeio'
```

### Tweets using a URL (matches expanded form inside short links)

```sql
select
  *
from
  twitter_search
where
  query = 'url:steampipe.io'
```
