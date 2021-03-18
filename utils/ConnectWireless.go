package utils

import (
	"log"
	"os/exec"
	"strconv"
	"time"
)

func Connect(ipAddress string, port int) {
	bind := ipAddress + ":" + strconv.Itoa(port)
	log.Println("[ADB_Library] Connect to", bind)

	cmd := exec.Command("adb", "connect", bind)
	Action(cmd, time.Second * 1)
}