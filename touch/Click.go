package touch

import (
	"github.com/ProjectTL12345/ADB_Library/utils"
	"log"
	"os/exec"
	"strconv"
)

func NormalClick(position1 int, position2 int) {
	log.Printf("[ADB_Library] Click Screen from Position [ X: %d, Y: %d ]", position1, position2)
	cmd := exec.Command("adb", "shell", "input", "touchscreen", "tap", strconv.Itoa(position1), strconv.Itoa(position2))
	utils.Action(cmd, 200)
}

func SwipeClick(firstPosition1 int, firstPosition2 int, secondPosition1 int, secondPosition2 int) {
	log.Printf("[ADB_Library] Click Screen from Position [ X1: %d, Y1: %d, X2: %d, Y2: %d ]", firstPosition1, firstPosition2, secondPosition1, secondPosition2)

	cmd := exec.Command("adb", "shell", "input", "touchscreen", "swipe",
		strconv.Itoa(firstPosition1),
		strconv.Itoa(firstPosition2),
		strconv.Itoa(secondPosition1),
		strconv.Itoa(secondPosition2))
	utils.Action(cmd, 300)
}
