package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const baseURLEnv = "DEVCTL_BASE_URL"

func LoadDefault() error {
	return LoadDotEnv(".env")
}

func LoadDotEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "export ") {
			line = strings.TrimSpace(strings.TrimPrefix(line, "export "))
		}
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			return errors.New("invalid .env line: " + line)
		}
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		if key == "" {
			return errors.New("invalid .env line: " + line)
		}
		if len(val) >= 2 {
			if (val[0] == '"' && val[len(val)-1] == '"') || (val[0] == '\'' && val[len(val)-1] == '\'') {
				val = val[1 : len(val)-1]
			}
		}
		if os.Getenv(key) == "" {
			if err := os.Setenv(key, val); err != nil {
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func BaseURL() (string, error) {
	baseURL := os.Getenv(baseURLEnv)
	if baseURL == "" {
		return "", errors.New("DEVCTL_BASE_URL not set")
	}
	return baseURL, nil
}
