---
title: "Steampipe Table: twitter_search_recent - Query Twitter Recent Searches using SQL"
description: "Allows users to query recent searches on Twitter, providing insights into the most recent search queries and related metadata."
---

# Table: twitter_search_recent - Query Twitter Recent Searches using SQL

Twitter is a social media platform that allows users to post and interact with messages known as "tweets". Twitter API provides programmatic access to Twitter data, including user profile information, tweets, and more. The recent searches resource provides information about the most recent search queries made by a user.

## Table Usage Guide

The `twitter_search_recent` table provides insights into the most recent search queries made on Twitter. As a data analyst, you can explore query-specific details through this table, including the query text, the timestamp of the query, and other related metadata. Use it to understand user behavior, track trending topics, or analyze the popularity of certain keywords or hashtags.

## Examples

### Tweets matching a hashtag
Identify instances where specific hashtags are used on Twitter. This can help in tracking the popularity and reach of a marketing campaign or trending topic.

```sql+postgres
select
  *
from
  twitter_search_recent
where
  query = '#rowscoloredglasses'
```

```sql+sqlite
select
  *
from
  twitter_search_recent
where
  query = '#rowscoloredglasses'
```

### Tweets mentioning a user
Discover the segments that mention a specific user on Twitter. This can be useful to understand the public sentiment or perception about the user in the Twitter community.

```sql+postgres
select
  *
from
  twitter_search_recent
where
  query = '@steampipeio'
```

```sql+sqlite
select
  *
from
  twitter_search_recent
where
  query = '@steampipeio'
```

### Tweets by a given author
Explore the recent posts made by a specific Twitter user. This can be useful for understanding the content and frequency of their tweets, which could be beneficial for social media analysis or competitor research.

```sql+postgres
select
  *
from
  twitter_search_recent
where
  query = 'from:steampipeio'
```

```sql+sqlite
select
  *
from
  twitter_search_recent
where
  query = 'from:steampipeio'
```

### Tweets in reply to a user
Explore which tweets are in response to a specific user to understand their social engagement and interactions on Twitter. This can be useful in identifying trends, gauging public sentiment, or tracking customer service interactions.

```sql+postgres
select
  *
from
  twitter_search_recent
where
  query = 'to:steampipeio'
```

```sql+sqlite
select
  *
from
  twitter_search_recent
where
  query = 'to:steampipeio'
```

### Tweets using a URL (matches expanded form inside short links)
Explore the tweets that contain a specific URL to gain insights into the public discussion related to a particular website, in this case, 'steampipe.io'. This could be useful for monitoring brand mentions, tracking campaign performance, or understanding audience engagement.

```sql+postgres
select
  *
from
  twitter_search_recent
where
  query = 'url:steampipe.io'
```

```sql+sqlite
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