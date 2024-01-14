package utils

import (
	"strings"
	"time"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func SessionSSH(s ssh.Session) {
	var secret string
	t := term.NewTerminal(s, "")

	// Check secret
	if Config.RequireSecret {
		for secret != Config.Secret {
			secret, _ = t.ReadPassword("")
			if secret == "exit" {
				return
			}
			time.Sleep(time.Second)
		}
	}

	// Chat with GPT over SSH!
	for {
		t.SetPrompt("👦: ")
		input, _ := t.ReadLine()

		// Handle exit
		if input == "exit" {
			sendMessage(&s, "Goodbye!")
			return
		}

		words := len(strings.Fields(input))
		if words < 3 {
			sendMessage(&s, "Hey, at least 3 words please.")
			continue
		}
		// ChatGPT Query
		sendMessage(&s, "Gimme sec...")
		answer, _ := Chat(input)
		sendMessage(&s, answer)
	}
}
