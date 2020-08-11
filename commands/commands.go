package commands

import "fmt"

func Hello(senderUsername string) string {
	return fmt.Sprintf("Hello <@%s> :)", senderUsername)
}
