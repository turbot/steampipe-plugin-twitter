---
title: "Steampipe Table: twitter_user_following - Query Twitter User Followings using SQL"
description: "Allows users to query Twitter User Followings, providing a list of users that a specific Twitter user is following."
---

# Table: twitter_user_following - Query Twitter User Followings using SQL

Twitter User Following is a feature within Twitter that allows a user to follow other Twitter users, thereby receiving their tweets and updates. It is a way for users to curate their Twitter feed based on their interests and preferences. This feature forms a fundamental part of the Twitter social media platform, enabling user interaction and engagement.

## Table Usage Guide

The `twitter_user_following` table provides insights into the users that a specific Twitter user is following. As a Social Media Analyst, explore user-specific details through this table, including the list of users they follow, their interests, and their engagement. Utilize it to understand user behavior, their preferences, and to create targeted marketing strategies.

## Examples

### List the follows for a user
Discover the segments that a specific user is following on Twitter, providing insights into their interests and affiliations. This can be useful for social media analysis and targeted marketing strategies.

```sql+postgres
select
  *
from
  twitter_user_following
where
  user_id = '1318177503995985921' -- @steampipeio
```

```sql+sqlite
select
  *
from
  twitter_user_following
where
  user_id = '1318177503995985921' -- @steampipeio
```

### List follows by username
Discover the segments that are followed by a specific Twitter user, enabling you to gain insights into their interests and connections. This could be beneficial in understanding their network and potential influence.
Via subselect:

```sql+postgres
select
  uf.id,
  uf.username
from
  twitter_user_following as uf
where
  uf.user_id in
  (
    select
      id
    from
      twitter_user
    where
      username = 'steampipeio'
  )
```

```sql+sqlite
select
  uf.id,
  uf.username
from
  twitter_user_following as uf
where
  uf.user_id in
  (
    select
      id
    from
      twitter_user
    where
      username = 'steampipeio'
  )
```

### Find all users followed by both turbothq and nathanwallace
Discover the shared connections between two specific Twitter users to understand their common interests or potential collaborations. This can be beneficial in social media analysis or targeted marketing strategies.

```sql+postgres
select
  id,
  username
from
  twitter_user_following
where
  user_id in (select id from twitter_user where username = 'turbothq')

intersect

select
  id,
  username
from
  twitter_user_following
where
  user_id in (select id from twitter_user where username = 'nathanwallace')
```

```sql+sqlite
select
  id,
  username
from
  twitter_user_following
where
  user_id in (select id from twitter_user where username = 'turbothq')

union

select
  id,
  username
from
  twitter_user_following
where
  user_id in (select id from twitter_user where username = 'nathanwallace')
```

### Find users who follow you, but you don't follow them
Discover the segments of your Twitter audience that follow your account, but whom you have not followed back. This can be useful for identifying potential influencers or key accounts to engage with for networking and community-building purposes.

```sql+postgres
with account as 
(
  select
    id 
  from
    twitter_user 
  where
    username = 'turbothq' 
)
,
following as 
(
  select
    id,
    username 
  from
    twitter_user_following 
  where
    user_id in 
    (
      select
        id 
      from
        account
    )
)
,
followers as 
(
  select
    id,
    username 
  from
    twitter_user_follower 
  where
    user_id in 
    (
      select
        id 
      from
        account
    )
)
select
  username 
from
  followers 
where
  username not in 
  (
    select
      username 
    from
      following
  )
order by
  username
```

```sql+sqlite
with account as 
(
  select
    id 
  from
    twitter_user 
  where
    username = 'turbothq' 
)
,
following as 
(
  select
    id,
    username 
  from
    twitter_user_following 
  where
    user_id in 
    (
      select
        id 
      from
        account
    )
)
,
followers as 
(
  select
    id,
    username 
  from
    twitter_user_follower 
  where
    user_id in 
    (
      select
        id 
      from
        account
    )
)
select
  username 
from
  followers 
where
  username not in 
  (
    select
      username 
    from
      following
  )
order by
  username
```