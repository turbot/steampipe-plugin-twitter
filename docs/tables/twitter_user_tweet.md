---
title: "Steampipe Table: twitter_user_tweet - Query Twitter User Tweets using SQL"
description: "Allows users to query Twitter User Tweets, focusing on the specific tweets made by a user, providing insights into tweet content, engagement metrics, and user details."
---

# Table: twitter_user_tweet - Query Twitter User Tweets using SQL

Twitter is a social networking service where users post and interact with messages known as "tweets". It provides a platform for public conversation and content distribution, allowing users to post tweets and interact with others. The tweets can contain photos, videos, links, and text, making it a rich source of data for analysis and insights.

## Table Usage Guide

The `twitter_user_tweet` table provides insights into tweets made by specific Twitter users. As a data analyst or social media manager, explore tweet-specific details through this table, including content, engagement metrics, and user details. Utilize it to uncover information about tweet patterns, engagement rates, and the impact of specific tweets, facilitating better understanding and decision-making for social media strategies.

## Examples

### Last 5 tweets by the author
Explore the most recent social media activity of a specific user to understand their latest posts and updates. This is useful for tracking the latest updates from a user, potentially for monitoring or engagement purposes.

```sql+postgres
select
  id,
  text
from
  twitter_user_tweet
where
  user_id = '1318177503995985921' -- @steampipeio
order by
  created_at desc
limit
  5;
```

```sql+sqlite
select
  id,
  text
from
  twitter_user_tweet
where
  user_id = '1318177503995985921' -- @steampipeio
order by
  created_at desc
limit
  5;
```

### Find all tweets in a user timeline related to open source
Explore user timelines on Twitter to uncover tweets related to the topic of open source. This could be beneficial in identifying discussions or sentiments about open source in a specific user's timeline.

```sql+postgres
select
  id,
  ca->'entity'->>'name' as context,
  text
from
  twitter_user_tweet,
  jsonb_array_elements(context_annotations) as ca
where
  user_id = '8092452' -- @turbothq
  and ca->'entity'->>'name' = 'Open source';
```

```sql+sqlite
select
  id,
  json_extract(ca.value, '$.entity.name') as context,
  text
from
  twitter_user_tweet,
  json_each(context_annotations) as ca
where
  user_id = '8092452' -- @turbothq
  and json_extract(ca.value, '$.entity.name') = 'Open source';
```

### Get tweet timeline by username
Explore the tweets made by a specific Twitter user. This could be useful for understanding the user's posting habits, recent activity, or content preferences.
Via subselect:

```sql+postgres
select
  id,
  text
from
  twitter_user_tweet as t
where
  t.user_id in (select id from twitter_user where username = 'steampipeio');
```

```sql+sqlite
select
  id,
  text
from
  twitter_user_tweet as t
where
  t.user_id in (select id from twitter_user where username = 'steampipeio');
```

Via join:
```sql+postgres
select
  t.id,
  t.text
from
  twitter_user_tweet as t,
  twitter_user as u
where
  t.user_id = u.id
  and u.username = 'steampipeio';
```

```sql+sqlite
select
  t.id,
  t.text
from
  twitter_user_tweet as t,
  twitter_user as u
where
  t.user_id = u.id
  and u.username = 'steampipeio';
```