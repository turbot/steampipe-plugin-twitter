# Table: twitter_search_recent

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
  twitter_search_recent
where
  query = '#rowscoloredglasses'
```

### Tweets mentioning a user

```sql
select
  *
from
  twitter_search_recent
where
  query = '@steampipeio'
```

### Tweets by a given author

```sql
select
  *
from
  twitter_search_recent
where
  query = 'from:steampipeio'
```

### Tweets in reply to a user

```sql
select
  *
from
  twitter_search_recent
where
  query = 'to:steampipeio'
```

### Tweets using a URL (matches expanded form inside short links)

```sql
select
  *
from
  twitter_search_recent
where
  query = 'url:steampipe.io'
```

### Place, author, and text for tweets about the weather in Vermont

```
select 
  place ->> 'full_name' as place,
  author ->> 'username' as author,  
  text 
from 
  twitter_search_recent
where 
  query = 'weather' 
  and place ->> 'full_name' ~* ' vt$'  -- regex matches 'Barre VT' etc
```

### Tweets about weather within 10 miles of a lat/lon location 

Note: `point_radius` and related operators are not available with a basic ("Essential") account, see [operators by product](https://developer.twitter.com/en/docs/twitter-api/enterprise/rules-and-filtering/operators-by-product).

```
select 
  *
from 
  twitter_search_recent
where 
  query = 'weather point_radius:[-105.292778 40.019444 10mi]' 
```
