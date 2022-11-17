// Copyright 2022 Hal Canary.  See LICENSE.md.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/HalCanary/mastodoner/mammut"
)

func main() {
	if err := listfollows(os.Args[1:], os.Stdin); err != nil {
		log.Fatal(err)
	}
}

// Needs `read:search` and `read:accounts` authorization token.
func listfollows(args []string, statusReader io.Reader) error {
	var flagSet flag.FlagSet
	accountQuery := flagSet.String("q", "", "account query")
	tags := flagSet.Bool("tags", false, "tags")
	flagSet.Parse(args)

	mastodonInfo, err := mammut.GetMastodonInfo()
	if err != nil {
		return err
	}

	var follows []string

	if *tags {
		follows, err = mammut.GetFollowingTags(mastodonInfo.AccessToken, mastodonInfo.Host)
	} else {
		follows, err = mammut.GetFollowing(
			mastodonInfo.AccessToken, mastodonInfo.Host, *accountQuery)
	}
	if err != nil {
		return err
	}
	sort.Strings(follows)
	for _, f := range follows {
		fmt.Println(f)
	}
	return nil
}
