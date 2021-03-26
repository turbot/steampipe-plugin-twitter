---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/twitter.svg"
brand_color: "#1DA1F2"
display_name: Twitter
name: twitter
description: Steampipe plugin to query tweets, users and followers from Twitter.
---

# Twitter

Query tweets, users and followers from Twitter.


## Installation

To download and install the latest twitter plugin:

```bash
steampipe plugin install twitter
```

## Credentials

Twitter requires a developer account and API token for all requests, but offers a free tier. Apply on the [twitter developer website](https://developer.twitter.com/) to get your free token.


## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest twitter plugin will create a default connection named `twitter` in the `~/.steampipe/config/twitter.spc` file. You must edit this connection to include your API bearer token:

```hcl
connection "twitter" {
  plugin  = "twitter"
  # Bearer token from your Twitter project
  token = "AAAAAAAAAAAAAAAAAAAAAFL8NgEAAAAA2%2FyHFNeRK0CUoZ5ybpsHgnL91n0%3DEJPs4GsJVU8ZlrHYr1x0eyb4Br48WeqLtmM4aAKbIxpInTrrIu"
}
```

Alternatively, you may use a consumer key pair and access key pair:
```hcl
connection "twitter" {
  plugin          = "twitter"
  consumer_key    = "LhDEQ8kE1RXyHg2VwXPSlXhVh"
  consumer_secret = "mWZOzVuJccvRCUnVDcoQMyH1VeAph4EDvYN85KyhDtsZsAKFmz"
  access_token    = "8193452-AqA8J35qfGBZC6gS7dSJ7jQZnZmbVAQzutj7TVzhIH"
  access_secret   = "ZUBXaDIi3ouXP1RASz0VMZMykwVt2TiL3PVF5nRIwY72I"
}
```
