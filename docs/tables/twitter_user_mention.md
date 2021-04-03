# Table: twitter_user_mention

Tweets where the specified user is mentioned.

Note: The `user_id` field must be set in the `where` clause.

## Examples

### Last 5 tweets that mention the user

```sql
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
  5
```

### Get mention timeline by username

Via subselect:
```sql
select
  id,
  text
from
  twitter_user_mention as t
where
  t.user_id in (select id from twitter_user where username = 'steampipeio')
```

Via join:
```sql
select
  t.id,
  t.text
from
  twitter_user_mention as t,
  twitter_user as u
where
  t.user_id = u.id
  and u.username = 'steampipeio'
```
