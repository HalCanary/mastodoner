// Copyright 2022 Hal Canary.  See LICENSE.md.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/HalCanary/mastodoner/mammut"
)

func main() {
	if err := listfollows(os.Args[1:], os.Stdin); err != nil {
		log.Fatal(err)
	}
}

func listfollows(args []string, statusReader io.Reader) error {
	var flagSet flag.FlagSet
	accountQuery := flagSet.String("q", "", "account query")

	flagSet.Parse(args)

	mastodonInfo, err := mammut.GetMastodonInfo()
	if err != nil {
		return err
	}

	auth := "Bearer " + mastodonInfo.AccessToken
	follows, err := mammut.GetFollowing(auth, mastodonInfo.Host, *accountQuery)
	if err != nil {
		return err
	}
	for _, f := range follows {
		fmt.Println(f)
	}
	return nil
}