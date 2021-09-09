package main

import (
	"avito/slack"

	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type channel struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type botConfig struct {
	BotToken string    `json:"bot_token"`
	Channels []channel `json:"channels"`
}

type slackResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

var (
	fileName = flag.String("file", "", "file name") // id пользователя для диалога
)

func main() {
	flag.Parse()
	if *fileName == "" {
		fmt.Println("передайте имя файла -file")
		return
	}
	result, err := readBotConfig(*fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	slackClient := slack.NewClient(result.BotToken)
	for _, c := range result.Channels {
		err := slackClient.SendMessage(c.Text, c.Channel)
		if err != nil {
			fmt.Printf("%v \n", err)
			return
		}
	}
	fmt.Println("all messages sent")
}

func readBotConfig(fileName string) (*botConfig, error) {
	bb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error on read file %s", err)
	}
	result := botConfig{}
	err = json.Unmarshal(bb, &result)
	if err != nil {
		fmt.Printf("error %v on unmarshal file %s", err, string(bb))
	}
	return &result, nil
}
