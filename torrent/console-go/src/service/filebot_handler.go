package service

import (
	"log"
	"os/exec"
)

func ProcessFilebot(c chan *exec.Cmd) error {
	for {
		var cmd *exec.Cmd
		cmd = <-c
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Print(err)
		}
		log.Print(string(output))
	}
}
