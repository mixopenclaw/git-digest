package main

import "fmt"

func Authenticate() (string, error) {
	// stub: return token from env/config
	cfg, _ := LoadConfig("")
	if cfg.Token == "" {
		return "", fmt.Errorf("no token")
	}
	return cfg.Token, nil
}
