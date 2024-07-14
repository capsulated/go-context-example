package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	resp, err := doRequest()
	if err != nil {
		panic(err)
	}
	println(resp.Status)
}

func doRequest() (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://logiq.one", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with ctx: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform http request: %w", err)
	}

	return res, nil
}
