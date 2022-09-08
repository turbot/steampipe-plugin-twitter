package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-twitter/twitter"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: twitter.Plugin})
}
