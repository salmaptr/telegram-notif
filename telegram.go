package service

import (

  "bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
  "time"

)

type TelegramRequestBody struct {
  Text string 'json:"text"'
  UserID string 'json:"chat_id"'
}

func NotifyTelegram(telegramBotURL, msg, id string) error {
    TelegramBody, _ := json.Marshal(TelegramRequestBody{Text: msg, UserID:id})
    req, err := http.NewRequest(http.MethodPost, telegramBotURL, bytes.NewBuffer(TelegramBody))
    if err != nil {
        return err
    }

    req.Header.Add("Content-Type", "application/json")

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    if buf.String() != "ok" {
        return errors.New("Error")
    }
    return nil
}
