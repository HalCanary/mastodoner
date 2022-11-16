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
	if err := mainImpl(os.Args[1:], os.Stdin); err != nil {
		log.Fatal(err)
	}
}

func mainImpl(args []string, statusReader io.Reader) error {
	var flagSet flag.FlagSet
	spoiler := flagSet.String("spoiler", "", "spoiler text for content warning, if not empty")
	visibility := flagSet.String("visibility", "", "public or unlisted or private")
	inReplyToId := flagSet.String("replyTo", "", "in reply to id")
	flagSet.Parse(args)

	mastodonInfo, err := mammut.GetMastodonInfo()
	if err != nil {
		return err
	}

	statusData, err := io.ReadAll(statusReader)
	if err != nil {
		return err
	}

	post := mammut.Status{
		Status:      string(statusData),
		Language:    mastodonInfo.Language,
		SpoilerText: *spoiler,
		Visibility:  *visibility,
		InReplyToId: *inReplyToId,
	}

	if mastodonInfo.MaximumStatusLength > 0 {
		size := len(post.Status) + len(post.SpoilerText)
		if size > mastodonInfo.MaximumStatusLength {
			return fmt.Errorf("error: status data length is %d out of %d availible bytes.",
				size, mastodonInfo.MaximumStatusLength)
		}
	}
	auth := "Bearer " + mastodonInfo.AccessToken
	id, url, err := mammut.PostStatus(auth, mastodonInfo.Host, post)
	if err != nil {
		return err
	}
	fmt.Println(id, url)
	return nil
}
