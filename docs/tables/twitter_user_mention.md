---
title: "Steampipe Table: twitter_user_mention - Query Twitter User Mentions using SQL"
description: "Allows users to query Twitter User Mentions, specifically the details of the users who have mentioned the authenticated user in their tweets."
---

# Table: twitter_user_mention - Query Twitter User Mentions using SQL

Twitter User Mentions are instances when a Twitter user includes another user's Twitter handle in their tweet. This is often done to draw the mentioned user's attention to the tweet or to engage them in a conversation. User mentions are a key part of the interactive nature of Twitter, allowing users to engage with each other and fostering discussions on the platform.

## Table Usage Guide

The `twitter_user_mention` table provides insights into Twitter User Mentions, specifically the details of the users who have mentioned the authenticated user in their tweets. As a social media analyst, explore user-specific details through this table, including the frequency of mentions, the context of mentions, and associated metadata. Utilize it to uncover information about user interactions, such as the most active users, the nature of interactions, and the reach of your tweets.

**Important Notes**
- The `user_id` field must be set in the `where` clause.

## Examples

### Last 5 tweets that mention the user
Explore the most recent social media mentions of a specific user to stay updated with the latest conversations involving them. This can be particularly useful for monitoring brand reputation, tracking user engagement, or responding to customer feedback.

```sql+postgres
select
  id,
  text
from
  twitter_user_mention
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
  twitter_user_mention
where
  user_id = '1318177503995985921' -- @steampipeio
order by
  created_at desc
limit
  5;
```

### Get mention timeline by username
Discover the specific instances where a particular username is mentioned on Twitter. This can be useful in monitoring brand mentions, tracking conversations involving a specific user, or understanding user engagement trends.
Via subselect:

```sql+postgres
select
  id,
  text
from
  twitter_user_mention as t
where
  t.user_id in (select id from twitter_user where username = 'steampipeio');
```

```sql+sqlite
select
  id,
  text
from
  twitter_user_mention as t
where
  t.user_id in (select id from twitter_user where username = 'steampipeio');
```

Via join:
```sql+postgres
select
  t.id,
  t.text
from
  twitter_user_mention as t,
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
  twitter_user_mention as t,
  twitter_user as u
where
  t.user_id = u.id
  and u.username = 'steampipeio';
```