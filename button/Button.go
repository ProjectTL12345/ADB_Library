package button

import (
	"ADB_Library/src/utils"
	"log"
	"os/exec"
)

func HomeButton() {
	log.Println("[ADB_Library] Click Home button")
	cmd := exec.Command("adb", "shell", "input", "keyevent", "3")
	utils.Action(cmd, 200)
}

func BackButton() {
	log.Println("[ADB_Library] Click Back button")
}

func PowerButton() {
	log.Println("[ADB_Library] Click Power button")
	cmd := exec.Command("adb", "shell", "input", "keyevent", "26")
	utils.Action(cmd, 200)
}

func EnterButton() {
	log.Println("[ADB_Library] Click Enter button")
	cmd := exec.Command("adb", "shell", "input", "keyevent", "66")
	utils.Action(cmd, 200)
}