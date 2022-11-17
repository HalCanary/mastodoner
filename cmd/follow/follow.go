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

func follow(args []string) error {
	mastodonInfo, err := mammut.GetMastodonInfo()
	if err != nil {
		return err
	}
	auth := "Bearer " + mastodonInfo.AccessToken
	for _, arg := range args {
		if err = mammut.Follow(auth, mastodonInfo.Host, arg); err != nil {
			return err
		}
	}
	return nil
}
