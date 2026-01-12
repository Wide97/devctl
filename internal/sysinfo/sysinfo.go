package sysinfo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Info struct {
	OS      string `json:"os"`
	Arch    string `json:"arch"`
	Runtime string `json:"runtime"`
}

const timeout = 5 * time.Second

func GetInfo(ctx context.Context) (Info, error) {
	baseURL := os.Getenv("DEVCTL_BASE_URL")
	if baseURL == "" {
		return Info{}, errors.New("DEVCTL_BASE_URL not set")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		baseURL+"/sys/info",
		nil,
	)
	if err != nil {
		return Info{}, err
	}

	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return Info{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Info{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var info Info
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return Info{}, err
	}

	if info.OS == "" || info.Arch == "" || info.Runtime == "" {
		return Info{}, errors.New("incomplete sys info received")
	}

	return info, nil
}
