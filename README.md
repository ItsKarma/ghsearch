# GitHub Search

Command line tool to search GitHub.

Search string is case insensitive.

## Installation

`git clone https://github.com/ItsKarma/ghsearch`

You can run directly without building:

`go run . <searchString> <flags>`

If you build a binary, you can put it in your bin directory so you can use it anywhere:

`go build .`

Make sure `/usr/local/bin` is in your path.

`mv ghsearch /usr/local/bin`

*Tip:* Make sure `/usr/local/bin` is in your `$PATH`.

Now you can use `ghsearch` from anywhere.

`ghsearch <searchString> <flags>`

## Usage

Requires user to set `GH_USER` and `GH_TOKEN` environment variables prior to use.

`command <searchString> <flags>`

## Examples

`ghsearch us-west-2 --path="/terraform/" --org="MyOrg" --output=csv`

`ghsearch "fmt" --repo="ItsKarma/ghsearch" --output=text`

## Inputs

| Flag | Description | Type | Optional | Default |
| ---- | ----------- | ---- | -------- | ------- |
| path | Path within repository(ies) to search. | string | True | n/a |
| org | Organization to search within. (mutually exclusive with repo) | string | True | n/a |
| repo | Repository to search within. (mutually exclusive with org) | string | True | n/a |
| output | Format of the output. | string["json", "text", "csv"] | True | json |
| v | Verbose - Outputs fields for debugging. | bool | True | n/a |
| vv | Very verbose - Outputs user and token !! logs sensitive information !! | bool | True | n/a |
| vvv | Very very verbose - Also outpus raw unfiltered GitHub response. | bool | True | n/a |