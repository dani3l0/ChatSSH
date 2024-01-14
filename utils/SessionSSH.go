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
		t.SetPrompt("> ")
		input, _ := t.ReadLine()

		// Handle exit
		if input == "exit" {
			s.Write([]byte("Goodbye!\n"))
			return
		}

		words := len(strings.Fields(input))
		if words < 3 {
			s.Write([]byte("Hey, at least 3 words please.\n"))
			continue
		}
		// ChatGPT Query
		s.Write([]byte("Okay, gimme a sec...\n"))
		answer, _ := Chat(input)
		s.Write([]byte(answer + "\n"))
	}
}
