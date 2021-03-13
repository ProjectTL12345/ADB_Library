package utils

import (
	"os/exec"
	"time"
)

func Action(command *exec.Cmd, sleep time.Duration) {
	command.Start()
	command.Wait()

	time.Sleep(sleep)
}