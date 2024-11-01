package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"os"
	"path/filepath"
)

type User struct {
    Name  string `json:"name"`
    Avatar string `json:"avatar"`
    Webhook_url string `json:"Webhook_url"`
}

var (
    RED   = "\033[0;31m"
    GREEN = "\033[0;32m"
    NC    = "\033[0m"
)

func main() {
    args := os.Args[1:]
    text := strings.Join(args, " ")

	homeDir := os.Getenv("HOME")
	configPath := filepath.Join(homeDir, ".config", "nwsend", "config.json")

	jsonData, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(RED, "#[error]:", err, NC)
		return
	}

	var user User
	if err := json.Unmarshal(jsonData, &user); err != nil {
		fmt.Println(RED, "#[error]:", err, NC)
		return
	}
   
    if text == "" {
        fmt.Println(RED, "[error]:", GREEN, "usage ```nwsend hi```", NC)
        return
    }
    sendWebhook(text, user.Name, user.Avatar, user.Webhook_url)
}

