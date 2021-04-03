# Table: twitter_user_tweet

Tweets published by the specified user ID (author).

Note: The `user_id` field must be set in the `where` clause.

## Examples

### Last 5 tweets by the author

```sql
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
  5
```

### Find all tweets in a user timeline related to open source

```sql
select
  id,
  ca->'entity'->>'name' as context,
  text
from
  twitter_user_tweet,
  jsonb_array_elements(context_annotations) as ca
where
  user_id = '8092452' -- @turbothq
  and ca->'entity'->>'name' = 'Open source'
```

### Get tweet timeline by username

Via subselect:
```sql
select
  id,
  text
from
  twitter_user_tweet as t
where
  t.user_id in (select id from twitter_user where username = 'steampipeio')
```

Via join:
```sql
select
  t.id,
  t.text
from
  twitter_user_tweet as t,
  twitter_user as u
where
  t.user_id = u.id
  and u.username = 'steampipeio'
```
