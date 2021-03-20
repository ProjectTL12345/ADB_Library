package tasks

import (
	"github.com/ProjectTL12345/ADB_Library/utils"
	"log"
	"os/exec"
)

func OpenApp(packageName string) {
	log.Println("[ADB_Library] Open app with package name for:", packageName)

	cmd := exec.Command("adb", "shell", "am", "start", "-n", packageName)
	utils.Action(cmd, 200)
}