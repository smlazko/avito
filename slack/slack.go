package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type SlackClient struct {
	token string
}

type slackResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func NewClient(token string) SlackClient {
	return SlackClient{
		token: token,
	}
}

func (s SlackClient) SendMessage(text, channel string) error {
	form := url.Values{}
	form.Set("channel", channel)
	form.Set("text", text)
	req, err := http.NewRequest(
		http.MethodPost,
		"https://slack.com/api/chat.postMessage",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+s.token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := slackResponse{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return fmt.Errorf("error %w on unmarshal resonse %s", err, data)
	}
	if !result.Ok {
		return fmt.Errorf("error %s on send message ", result.Error)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response from slack [%s]: %s", resp.Status, data)
	}
	return nil
}
