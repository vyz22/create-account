package otp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type Message struct {
	ID   int    `json:"id"`
	From string `json:"from"`
	Subj string `json:"subject"`
}

func GetMessages(login, domain string) ([]Message, error) {
	url := fmt.Sprintf("https://zacky.my.id/api/v1/?action=getMessages&login=%s&domain=%s", login, domain)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var msgs []Message
	json.Unmarshal(body, &msgs)

	return msgs, nil
}
func ReadMessage(login, domain string, id int) (string, error) {
	url := fmt.Sprintf("https://zacky.my.id/api/v1/?action=readMessage&login=%s&domain=%s&id=%d", login, domain, id)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}
func ExtractOTP(text string) string {
	re := regexp.MustCompile(`\d{4,6}`)
	return re.FindString(text)
}
