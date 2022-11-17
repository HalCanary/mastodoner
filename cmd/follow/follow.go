// Copyright 2022 Hal Canary.  See LICENSE.md.
package main

import (
	"log"
	"os"

	"github.com/HalCanary/mastodoner/mammut"
)

func main() {
	if err := follow(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

// Needs `read:search` and `write:follows` authorization token.
func follow(args []string) error {
	mastodonInfo, err := mammut.GetMastodonInfo()
	if err != nil {
		return err
	}
	for _, arg := range args {
		if err = mammut.Follow(mastodonInfo.AccessToken, mastodonInfo.Host, arg); err != nil {
			return err
		}
	}
	return nil
}
