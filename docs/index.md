---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/twitter.svg"
brand_color: "#1DA1F2"
display_name: Twitter
name: twitter
description: Steampipe plugin to query tweets, users and followers from Twitter.
og_description: Query Twitter with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/twitter-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Twitter + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[Twitter](https://twitter.com) is an American microblogging and social networking service on which users post and interact with messages known as "tweets".

For example:

```sql
select
  id,
  text,
  mentions
from
  twitter_search_recent
where
  query = '#shiftleftjoin'
```

```
+---------------------+-----------------------------------------------------+-----------------+
| id                  | text                                                | mentions        |
+---------------------+-----------------------------------------------------+-----------------+
| 1378041446687768578 | New! Steampipe GCP plugin v0.5.0 🚀                 | ["googlecloud"] |
|                     |                                                     |                 |
|                     | Docs — https://t.co/s407cEdLAE                      |                 |
|                     |                                                     |                 |
|                     | Changelog — https://t.co/nFbLeh0KMu                 |                 |
|                     |                                                     |                 |
|                     | #shiftleftjoin @googlecloud https://t.co/0bGH747VRs |                 |
| 1377720348804853761 | New! Steampipe AWS plugin v0.10.0 🚀                | ["awscloud"]    |
|                     |                                                     |                 |
|                     | Docs - https://t.co/Y0vghMB1ub                      |                 |
|                     |                                                     |                 |
|                     | Changelog – https://t.co/8IETtyFBGZ                 |                 |
|                     |                                                     |                 |
|                     | #shiftleftjoin @awscloud https://t.co/o7ifZQlRMe    |                 |
+---------------------+-----------------------------------------------------+-----------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/twitter/tables)**

## Get started

### Install

Download and install the latest Twitter plugin:

```bash
steampipe plugin install twitter
```

### Credentials

| Item | Description |
| - | - |
| Credentials | Apply for a [free developer account](https://developer.twitter.com/en/apply-for-access). Create a project. Download the bearer token. |
| Permissions | `Read Only` is required for the app. |
| Radius | Each connection represents a single set of Twitter credentials. |
| Resolution |  1. `bearer_token` in Steampipe config.<br />2. `consumer_key`, `consumer_secret`, `access_token`, `access_secret` in Steampipe config.<br />3. `TWITTER_BEARER_TOKEN` environment variable.<br />4. `TWITTER_CONSUMER_KEY`, `TWITTER_CONSUMER_SECRET`, `TWITTER_ACCESS_TOKEN` and `TWITTER_ACCESS_SECRET` environment variables. |

### Configuration

Installing the latest twitter plugin will create a config file (`~/.steampipe/config/twitter.spc`) with a single connection named `twitter`:

```hcl
connection "twitter" {
  plugin       = "twitter"
  bearer_token = "AAAAAAAAAAAAAAAAAAAAAFL8NgEAAAAA2%2FyHFNeRK0CUoZ5ybpsHgnL91n0%3DEJPs4GsJVU8ZlrHYr1x0eyb4Br48WeqLtmM4aAKbIxpInTrrIu"
}
```


