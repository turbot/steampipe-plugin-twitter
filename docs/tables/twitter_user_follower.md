# Table: twitter_user_follower

Query users who are followers of the specified user ID.

Note: The `user_id` field must be provided.

## Examples

### List followers for a user

```sql
select
  *
from
  twitter_user_follower
where
  user_id = '1318177503995985921' -- @steampipeio
```

### List followers by username

Via subselect:
```sql
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
  )
```

Via join:
```sql
select
  uf.id,
  uf.username
from
  twitter_user_follower as uf,
  twitter_user as u
where
  uf.user_id = u.id
  and u.username = 'steampipeio'
```

### Find the top 10 followers for a user

Via join:
```sql
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
  10
```
