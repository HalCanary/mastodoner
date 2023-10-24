// Copyright 2022 Hal Canary.  See LICENSE.md.
package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// @param result will be JSON-decoded
func Get(auth, host, path string, query map[string]string, result any) error {
	q := url.Values{}
	for key, value := range query {
		q[key] = []string{value}
	}
	getUrl := url.URL{Scheme: "https", Host: host, Path: path, RawQuery: q.Encode()}
	req := http.Request{Method: "GET", URL: &getUrl}
	if auth != "" {
		req.Header = http.Header{"Authorization": []string{auth}}
	}
	return doHttp(&req, result)
}

// @param post will be JSON-encoded.
// @param result will be JSON-decoded
func Post(auth, host, path string, post, result any) error {
	postBytes, _ := json.Marshal(post)
	postUrl := url.URL{Scheme: "https", Host: host, Path: path}
	req := http.Request{
		Method:        "POST",
		URL:           &postUrl,
		Header:        http.Header{"Content-Type": []string{"application/json"}},
		Body:          io.NopCloser(bytes.NewReader(postBytes)),
		ContentLength: int64(len(postBytes)),
	}
	if auth != "" {
		req.Header["Authorization"] = []string{auth}
	}
	return doHttp(&req, result)
}

func doHttp(req *http.Request, result any) error {
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp == nil {
		return errors.New("Nil Reponse")
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Response: %q %q", resp.Status, string(body))
	}
	//prnt(body)
	if result != nil {
		return json.Unmarshal(body, result)
	}
	return nil
}

func prnt(src []byte) {
	var x interface{}
	json.Unmarshal(src, &x)
	y, _ := json.MarshalIndent(x, "", "    ")
	fmt.Println(string(y))
}
