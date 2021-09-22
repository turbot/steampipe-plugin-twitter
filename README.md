![image](https://hub.steampipe.io/images/plugins/turbot/twitter-social-graphic.png)

# Twitter Plugin for Steampipe

Use SQL to query tweets, users and followers from Twitter.

* **[Get started →](https://hub.steampipe.io/plugins/turbot/twitter)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/twitter/tables)
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-twitter/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install twitter
```

Run a query:
```sql
select
  *
from
  twitter_search
where
  query = '#rowscoloredglasses'
```

## Developing

Prerequisites:
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-twitter.git
cd steampipe-plugin-twitter
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:
```
make
```

Configure the plugin:
```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/twitter.spc
```

Try it!
```
steampipe query
> .inspect twitter
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-twitter/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Twitter Plugin](https://github.com/turbot/steampipe-plugin-twitter/labels/help%20wanted)
