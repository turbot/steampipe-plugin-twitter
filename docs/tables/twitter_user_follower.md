---
title: "Steampipe Table: twitter_user_follower - Query Twitter User Followers using SQL"
description: "Allows users to query Twitter User Followers, providing insights into the followers of a specific Twitter user."
---

# Table: twitter_user_follower - Query Twitter User Followers using SQL

Twitter is a social media platform that allows users to post and interact with messages known as "tweets". A significant aspect of Twitter is the concept of "followers". A follower on Twitter is another user who has subscribed to your tweets, meaning your tweets will appear in their timeline.

## Table Usage Guide

The `twitter_user_follower` table provides insights into the followers of a specific Twitter user. As a social media analyst, you can explore follower-specific details through this table, including follower count, location, and associated metadata. Utilize it to uncover information about followers, such as their geographical distribution, follower count trends, and the verification status of followers.

**Important Notes**
- The `user_id` field must be set in the `where` clause.

## Examples

### List followers for a user
Explore which Twitter users are following a specific account. This can be useful for understanding the reach and influence of that account.

```sql+postgres
select
  *
from
  twitter_user_follower
where
  user_id = '1318177503995985921'; -- @steampipeio
```

```sql+sqlite
select
  *
from
  twitter_user_follower
where
  user_id = '1318177503995985921'; -- @steampipeio;
```

### List followers by username
Identify instances where you can find out who is following a specific user on Twitter. This can be useful for understanding the demographics of your audience or identifying influential followers.
Via subselect:

```sql+postgres
select
  uf.id,
  uf.username
from
  twitter_user_follower as uf
where
  uf.user_id in
  (
    select
      id
    from
      twitter_user
    where
      username = 'steampipeio'
  );
```

```sql+sqlite
select
  uf.id,
  uf.username
from
  twitter_user_follower as uf
where
  uf.user_id in
  (
    select
      id
    from
      twitter_user
    where
      username = 'steampipeio'
  );
```

Via join:
```sql+postgres
select
  uf.id,
  uf.username
from
  twitter_user_follower as uf,
  twitter_user as u
where
  uf.user_id = u.id
  and u.username = 'steampipeio';
```

```sql+sqlite
select
  uf.id,
  uf.username
from
  twitter_user_follower as uf,
  twitter_user as u
where
  uf.user_id = u.id
  and u.username = 'steampipeio';
```

### Find the top 10 followers for a user
Determine the top ten users who are following a specific user on Twitter, ranked by the number of their own followers. This can be useful for identifying influential followers and understanding the reach of your social network.
Via join:

```sql+postgres
select
  uf.id,
  uf.username,
  (
    uf.public_metrics ->> 'followers_count'
  )
  ::int as follower_count
from
  twitter_user_follower as uf,
  twitter_user as u
where
  uf.user_id = u.id
  and u.username = 'steampipeio'
order by
  follower_count desc
limit
  10;
```

```sql+sqlite
select
  uf.id,
  uf.username,
  cast(json_extract(uf.public_metrics, '$.followers_count') as integer) as follower_count
from
  twitter_user_follower as uf,
  twitter_user as u
where
  uf.user_id = u.id
  and u.username = 'steampipeio'
order by
  follower_count desc
limit
  10;
```