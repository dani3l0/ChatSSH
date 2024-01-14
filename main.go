package main

import (
	"time"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		var pass string
		t := term.NewTerminal(s, "")

		for pass != "pass" {
			pass, _ = t.ReadPassword("")
			if pass == "exit" {
				s.Close()
				return
			}
			time.Sleep(time.Second)
		}

		for {
			t.SetPrompt("> ")
			input, _ := t.ReadLine()
			s.Write([]byte(input))

			if input == "exit" {
				s.Close()
				return
			}
		}
	})

	ssh.ListenAndServe(":2222", nil)
}
