package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendWebhook(text, username, avatar, webhook_url string) error {
    message := map[string]string{
	"content":   text,
	"username":  username,
	"avatar_url": avatar,
    }
	
    jsonData, err := json.Marshal(message)
	
    if err != nil {
        return fmt.Errorf("[error]: %v", err)
	}
	resp, err := http.Post(webhook_url, "application/json", bytes.NewBuffer(jsonData))
	
    if err != nil {
		return fmt.Errorf("#[error]: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("![error]: %d", resp.StatusCode)
	}

	fmt.Println("[ok]")
	return nil
}

