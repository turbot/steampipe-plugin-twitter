# Table: twitter_user

Get information about a single user.

To query lists of users see
[twitter_user_follower](../twitter_user_follower) or
[twitter_user_following](../twitter_user_following) or
[twitter_user_mention](../twitter_user_mention) instead.

Note: The `id` (preferred) or `username` field must be set in the `where` clause.

## Examples

### Get user by ID

```sql
select
  *
from
  twitter_user
where
  id = '1318177503995985921' -- @steampipeio
```

### Get user by username

```sql
select
  *
from
  twitter_user
where
  username = 'steampipeio'
```
