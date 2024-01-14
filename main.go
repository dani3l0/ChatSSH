package main

import (
	"chatssh/utils"
	"fmt"
	"time"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func main() {
	conf := &utils.Config
	utils.LoadConfig()

	ssh.Handle(func(s ssh.Session) {
		var pass string
		t := term.NewTerminal(s, "")

		if conf.RequireSecret {
			for pass != conf.Secret {
				pass, _ = t.ReadPassword("")
				if pass == "exit" {
					s.Close()
					return
				}
				time.Sleep(time.Second)
			}
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

	fmt.Println("SSH server is listening on " + conf.ListenAddr)
	e := ssh.ListenAndServe(conf.ListenAddr, nil)
	if e != nil {
		fmt.Println(e.Error())
	}
}
