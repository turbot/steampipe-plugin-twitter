# Table: twitter_user

Get information about a user.

Note: The `id` (priority) or `username` field must be provided.

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
  id = 'steampipeio'
```
