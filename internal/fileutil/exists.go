package fileutil

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func Exists(ctx context.Context, baseURL, path string) (bool, error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	endpoint := fmt.Sprintf("%s/file/exists?path=%s", baseURL, url.QueryEscape(path))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return false, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true, nil
	} else if resp.StatusCode == 404 {
		return false, nil
	}

	return false, fmt.Errorf("service returned %d", resp.StatusCode)
}
