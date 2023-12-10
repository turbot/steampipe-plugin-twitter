---
title: "Steampipe Table: twitter_tweet - Query Twitter Tweets using SQL"
description: "Allows users to query Twitter Tweets, specifically the content, user, location, and other metadata associated with individual tweets."
---

# Table: twitter_tweet - Query Twitter Tweets using SQL

Twitter is a widely used social media platform that allows users to post and interact with messages known as "tweets". The platform provides a robust API, enabling the extraction of vast amounts of public data for analysis. This includes data about tweets, users, hashtags, and more.

## Table Usage Guide

The `twitter_tweet` table gives insights into individual tweets on the Twitter platform. As a data analyst or social media manager, you can delve into tweet-specific details through this table, including content, user, location, and associated metadata. Use it to analyze tweet patterns, user interactions, and trending topics for strategic decision-making and targeted marketing campaigns.

## Examples

### Get tweet by ID
Discover the specific content and details of a particular tweet using its unique identification number. This can be useful for understanding the context, engagement, or user behavior related to that particular tweet.

```sql+postgres
select
  *
from
  twitter_tweet
where
  id = '1373134228620214275';
```

```sql+sqlite
select
  *
from
  twitter_tweet
where
  id = '1373134228620214275';
```