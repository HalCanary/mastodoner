# mastodoner

programs for talking to a mastodon server.

See [LICENSE.md](LICENSE.md).

* * *

## To Use

1.  Sync this repo:

        git clone https://github.com/HalCanary/mastodoner.git
        cd mastodoner/

2.  Install `go` version ≥ 1.19 (<https://go.dev/dl/>).

3.  Execute:

        go build -o . ./...

4.  Create an access token by visiting this url:

    https://YOUR_HOSTNAME_HERE/settings/applications

    Grant the following access:

    *   `read:accounts`
    *   `read:search`
    *   `write:follows`
    *   `write:statuses`

5.  Create a config file `~/mastodon.json` that looks like:

        {
            "AccessToken": "YOUR_ACCESS_TOKEN HERE",
            "Language": "en",
            "Host": "YOUR_HOSTNAME_HERE",
            "MaximumStatusLength": 500
        }

6.  Write your status in a text file, e.g. `STATUS.txt`

7.  Update your status:

        ./poststatus < STATUS.txt

    or

        ./poststatus -spoiler 'Content Warning: US politics' < STATUS.txt

