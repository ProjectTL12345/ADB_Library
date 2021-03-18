package touch

import (
	"ADB_Library/src/utils"
	"log"
	"os/exec"
	"strconv"
)

func NormalClick(position1 int, position2 int) {
	log.Printf("[ADB_Library] Click Screen from Position[ X: %d, Y: %d ]", position1, position2)
	cmd := exec.Command("adb", "shell", "input", "touchscreen", "tap", strconv.Itoa(position1), strconv.Itoa(position2))
	utils.Action(cmd, 200)
}

func SwipeClick(firstPosition1 int, firstPosition2 int, secondPosition1 int, secondPosition2 int) {
	cmd := exec.Command("adb", "shell", "input", "touchscreen", "swipe",
		strconv.Itoa(firstPosition1),
		strconv.Itoa(firstPosition2),
		strconv.Itoa(secondPosition1),
		strconv.Itoa(secondPosition2))
	utils.Action(cmd, 300)
}