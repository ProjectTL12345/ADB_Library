package utils

import (

	"log"
	"os/exec"
)

func ClickNumber(number int) {
	switch number {
	case 0:
		log.Println("[ADB_Library] Click button '0'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "7")
		Action(cmd, 200)
	case 1:
		log.Println("[ADB_Library] Click button '1'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "8")
		Action(cmd, 200)
	case 2:
		log.Println("[ADB_Library] Click button '2'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "9")
		Action(cmd, 200)
	case 3:
		log.Println("[ADB_Library] Click button '3'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "10")
		Action(cmd, 200)
	case 4:
		log.Println("[ADB_Library] Click button '4'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "11")
		Action(cmd, 200)
	case 5:
		log.Println("[ADB_Library] Click button '5'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "12")
		Action(cmd, 200)
	case 6:
		log.Println("[ADB_Library] Click button '6'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "13")
		Action(cmd, 200)
	case 7:
		log.Println("[ADB_Library] Click button '7'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "14")
		Action(cmd, 200)
	case 8:
		log.Println("[ADB_Library] Click button '8'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "15")
		Action(cmd, 200)
	case 9:
		log.Println("[ADB_Library] Click button '9'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "16")
		Action(cmd, 200)
	default:
		log.Fatalln("[ADB_Library] Invalid number, Please type 0 ~ 9. Typed Number:", number)
	}
}