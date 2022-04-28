## v0.1.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#15](https://github.com/turbot/steampipe-plugin-twitter/pull/15))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#14](https://github.com/turbot/steampipe-plugin-twitter/pull/14))

## v0.0.3 [2021-11-23]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) and Go version 1.17 ([#8](https://github.com/turbot/steampipe-plugin-twitter/pull/8))
- Updated the example in the document of `twitter_user` table ([#6](https://github.com/turbot/steampipe-plugin-twitter/pull/6))

## v0.0.2 [2021-09-22]

_What's new?_

- Clarified table documentation with where clause field requirements

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v161--2021-09-21) ([#4](https://github.com/turbot/steampipe-plugin-twitter/pull/4))
- Changed plugin license to Apache 2.0 per [turbot/steampipe](https://github.com/turbot/steampipe/issues/488) ([#2](https://github.com/turbot/steampipe-plugin-twitter/pull/2))

## v0.0.1 [2021-04-03]

_What's new?_

- New tables added
  - [twitter_search](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_search)
  - [twitter_tweet](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_tweet)
  - [twitter_user](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_user)
  - [twitter_user_follower](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_user_follower)
  - [twitter_user_following](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_user_following)
  - [twitter_user_mention](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_user_mention)
  - [twitter_user_tweet](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_user_tweet)
