# Table: twitter_user_following

Query which users the specified user follows.

Note: The `user_id` field must be set in the `where` clause.

## Examples

### List the follows for a user

```sql
select
  *
from
  twitter_user_following
where
  user_id = '1318177503995985921' -- @steampipeio
```

### List follows by username

Via subselect:
```sql
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

```sql
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

### Find users who follow you, but you don't follow them

```sql
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
