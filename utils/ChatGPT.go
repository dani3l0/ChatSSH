package utils

import (
	"context"
	"fmt"

	"github.com/ayush6624/go-chatgpt"
)

func Chat(question string) (string, error) {
	client, err := chatgpt.NewClient(Config.ChatGPTKey)
	if err != nil {
		return "Holy shit, I crashed! " + err.Error(), err
	}
	resp, err := client.SimpleSend(context.Background(), question)
	fmt.Println(resp)
	return resp.Choices[0].Message.Content, err
}
