package main

import (
	"chatssh/utils"
	"fmt"

	"github.com/gliderlabs/ssh"
)

func main() {
	// Config
	utils.LoadConfig()
	conf := &utils.Config

	// SSH Handlers
	ssh.Handle(utils.SessionSSH)

	// Run SSH Server
	fmt.Println("SSH server is listening on " + conf.ListenAddr)
	e := ssh.ListenAndServe(conf.ListenAddr, nil)
	fmt.Println(e.Error())
}
