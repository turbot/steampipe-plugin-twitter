## v0.5.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#46](https://github.com/turbot/steampipe-plugin-twitter/pull/46))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#46](https://github.com/turbot/steampipe-plugin-twitter/pull/46))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-twitter/blob/main/docs/LICENSE). ([#46](https://github.com/turbot/steampipe-plugin-twitter/pull/46))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to  column `_ctx`, and fixing connection and potential divide-by-zero bugs. ([#45](https://github.com/turbot/steampipe-plugin-twitter/pull/45))

## v0.4.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#38](https://github.com/turbot/steampipe-plugin-twitter/pull/38))

## v0.4.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#35](https://github.com/turbot/steampipe-plugin-twitter/pull/35))
- Recompiled plugin with Go version `1.21`. ([#35](https://github.com/turbot/steampipe-plugin-twitter/pull/35))

## v0.3.0 [2023-04-12]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#26](https://github.com/turbot/steampipe-plugin-twitter/pull/26))

## v0.2.0 [2022-09-09]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v416-2022-09-02) which includes several caching and memory management improvements. ([#21](https://github.com/turbot/steampipe-plugin-twitter/pull/21))
- Recompiled plugin with Go version `1.19`. ([#21](https://github.com/turbot/steampipe-plugin-twitter/pull/21))

## v0.1.1 [2022-07-20]

_Enhancements_

- Added geo-related examples and to `twitter_search_recent` table document. ([#11](https://github.com/turbot/steampipe-plugin-twitter/pull/11))

_Bug fixes_

- Fixed incorrect table name in `twitter_search_recent` table document example queries. ([#18](https://github.com/turbot/steampipe-plugin-twitter/pull/18))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v332--2022-07-11) which includes several caching fixes.

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
  - [twitter_search_recent](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_search_recent)
  - [twitter_tweet](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_tweet)
  - [twitter_user](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_user)
  - [twitter_user_follower](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_user_follower)
  - [twitter_user_following](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_user_following)
  - [twitter_user_mention](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_user_mention)
  - [twitter_user_tweet](https://hub.steampipe.io/plugins/turbot/twitter/tables/twitter_user_tweet)
