---
title: "Steampipe Table: twitter_user - Query Twitter Users using SQL"
description: "Allows users to query Twitter Users, specifically user profile details, providing insights into user behavior, interests, and associations."
---

# Table: twitter_user - Query Twitter Users using SQL

Twitter is a social networking service where users post and interact with messages known as 'tweets'. It allows users to follow others and be followed, retweet posts, like posts, and more. Twitter Users are the individual accounts that participate in this network, each with a unique username, profile details, and tweet history.

## Table Usage Guide

The `twitter_user` table provides insights into Twitter Users within the Twitter social networking service. As a social media analyst or digital marketer, explore user-specific details through this table, including profile information, follower counts, and tweet history. Utilize it to uncover information about users, such as their interests, associations, and influence within the Twitter community.

## Examples

### Get user by ID
Determine the specific user details on Twitter by using their unique ID. This can be particularly useful in understanding the user's activity and profile information without having to manually search for them.

```sql+postgres
select
  *
from
  twitter_user
where
  id = '1318177503995985921'; -- @steampipeio
```

```sql+sqlite
select
  *
from
  twitter_user
where
  id = '1318177503995985921'; -- @steampipeio
```

### Get user by username
Explore which user details are associated with a specific username on Twitter. This can be beneficial in scenarios where you need to understand the profile details of a particular user for research or analysis purposes.

```sql+postgres
select
  *
from
  twitter_user
where
  username = 'steampipeio';
```

```sql+sqlite
select
  *
from
  twitter_user
where
  username = 'steampipeio';
```