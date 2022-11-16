// Copyright 2022 Hal Canary.  See LICENSE.md.
package mammut

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/HalCanary/mastodoner/rest"
)

type MastodonInfo struct {
	AccessToken         string
	Language            string // ISO 639 language code
	Host                string
	MaximumStatusLength int
}

// Decode the JSON file `~/mastodon.json`.
func GetMastodonInfo() (MastodonInfo, error) {
	var result MastodonInfo
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return result, err
	}
	f, err := os.Open(filepath.Join(userHomeDir, "mastodon.json"))
	if err != nil {
		return result, err
	}
	err = json.NewDecoder(f).Decode(&result)
	return result, err
}

type Status struct {
	Status      string `json:"status,omitempty"`
	Visibility  string `json:"visibility,omitempty"`
	Language    string `json:"language,omitempty"`
	SpoilerText string `json:"spoiler_text,omitempty"`
	InReplyToId string `json:"in_reply_to_id,omitempty"`
	Sensitive   bool   `json:"sensitive,omitempty"`
}

func PostStatus(auth, host string, status Status) (id, url string, err error) {
	var result struct {
		Id  string `json:"id"`
		Url string `json:"url"`
	}
	err = rest.Post(auth, host, "/api/v1/statuses", &status, &result)
	return result.Id, result.Url, err
}

func GetAccountId(auth, host, accountQuery string) (string, error) {
	var value struct {
		Accounts []struct {
			Id   string `json:"id"`
			Acct string `json:"acct"`
		} `json:"accounts"`
	}
	err := rest.Get(auth, host, "/api/v2/search",
		map[string]string{"type": "accounts", "q": accountQuery}, &value)
	if err != nil {
		return "", err
	}
	if len(value.Accounts) != 1 {
		return "", fmt.Errorf("%d results for AccountId query", len(value.Accounts))
	}
	return value.Accounts[0].Id, nil
}