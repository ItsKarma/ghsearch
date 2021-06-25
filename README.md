# GitHub Search

Command line tool to search GitHub.

Search string is case insensitive.

## Usage

Requires user to set `GH_USER` and `GH_TOKEN` environment variables prior to use.

`command <searchString> <flags>`

|| Flag || Type || Optional || Default ||
| path | string | True | n/a |
| org | string | True | n/a |
| output | string | True | json |

## Examples

`ghsearch us-west-2 --path="/terraform/" --org="MyOrg" --output=csv`

`ghsearch "fmt" --repo="ItsKarma/ghsearch" --output=text`
