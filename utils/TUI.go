package utils

import "github.com/gliderlabs/ssh"

func sendMessage(_s *ssh.Session, text string) (int, error) {
	s := *_s
	str := "🤖: " + text + "\n\n"
	return s.Write([]byte(str))
}
