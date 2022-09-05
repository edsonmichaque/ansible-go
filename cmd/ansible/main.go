package main

import (
	"github.com/edsonmichaque/ansible-go/cmd/ansible/cmd"
	"os"
)

func main() {
	ansible := &cmd.Ansible{}

	if err := ansible.Command().Execute(); err != nil {
		os.Exit(1)
	}
}
