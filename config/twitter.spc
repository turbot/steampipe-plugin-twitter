connection "twitter" {
  plugin = "twitter"

  # Recommended:
  # OAuth 2.0 Bearer Token allows a Twitter developer app to access information
  # publicly available on Twitter.
  # bearer_token = "YOUR_BEARER_TOKEN"

  # Alternate:
  # OAuth 1.0a allows an authorized Twitter developer app to access private
  # account information or perform a Twitter action on behalf of a Twitter
  # account.
  # consumer_key    = "YOUR_CONSUMER_KEY"
  # consumer_secret = "YOUR_CONSUMER_SECRET"
  # access_token    = "YOUR_ACCESS_TOKEN"
  # access_secret   = "YOUR_ACCESS_SECRET"

  # Cap the number of items retreived from the API as part of each query,
  # preventing over-consumption of Twitter API limits. Defaults to 1000. Set
  # to -1 to indicate no limit.
  # max_items_per_query = 1000
}
