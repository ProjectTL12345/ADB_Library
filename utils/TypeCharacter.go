package utils

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

func ClickShift(getCharacter string, number bool) {
	log.Println("[ADB_Library] Click Shift button")

	cmd := exec.Command("adb", "shell", "input", "keyevent", "59")
	cmd.Start()
	time.Sleep(time.Millisecond * 80)
	if !number {
		setLowerCase := strings.ToLower(getCharacter)

		DictionaryCharacter(setLowerCase)
		cmd.Wait()
	} else {
		DictionaryCharacter(getCharacter)
		cmd.Wait()
	}
}

func DictionaryCharacter(getCharacter string) {
	switch getCharacter {
	case "a":
		log.Println("[ADB_Library] Click button 'a'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "29")
		Action(cmd, 200)
	case "b":
		log.Println("[ADB_Library] Click button 'b'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "30")
		Action(cmd, 200)
	case "c":
		log.Println("[ADB_Library] Click button 'c'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "31")
		Action(cmd, 200)
	case "d":
		log.Println("[ADB_Library] Click button 'd'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "32")
		Action(cmd, 200)
	case "e":
		log.Println("[ADB_Library] Click button 'e'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "33")
		Action(cmd, 200)
	case "f":
		log.Println("[ADB_Library] Click button 'f'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "34")
		Action(cmd, 200)
	case "g":
		log.Println("[ADB_Library] Click button 'g'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "35")
		Action(cmd, 200)
	case "h":
		log.Println("[ADB_Library] Click button 'h'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "36")
		Action(cmd, 200)
	case "i":
		log.Println("[ADB_Library] Click button 'i'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "37")
		Action(cmd, 200)
	case "j":
		log.Println("[ADB_Library] Click button 'j'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "38")
		Action(cmd, 200)
	case "k":
		log.Println("[ADB_Library] Click button 'k'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "39")
		Action(cmd, 200)
	case "l":
		log.Println("[ADB_Library] Click button 'l'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "40")
		Action(cmd, 200)
	case "m":
		log.Println("[ADB_Library] Click button 'm'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "41")
		Action(cmd, 200)
	case "n":
		log.Println("[ADB_Library] Click button 'n'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "42")
		Action(cmd, 200)
	case "o":
		log.Println("[ADB_Library] Click button 'o'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "43")
		Action(cmd, 200)
	case "p":
		log.Println("[ADB_Library] Click button 'p'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "44")
		Action(cmd, 200)
	case "q":
		log.Println("[ADB_Library] Click button 'q'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "45")
		Action(cmd, 200)
	case "r":
		log.Println("[ADB_Library] Click button 'r'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "46")
		Action(cmd, 200)
	case "s":
		log.Println("[ADB_Library] Click button 's'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "47")
		Action(cmd, 200)
	case "t":
		log.Println("[ADB_Library] Click button 't'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "48")
		Action(cmd, 200)
	case "u":
		log.Println("[ADB_Library] Click button 'u'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "49")
		Action(cmd, 200)
	case "v":
		log.Println("[ADB_Library] Click button 'v'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "50")
		Action(cmd, 200)
	case "w":
		log.Println("[ADB_Library] Click button 'w'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "51")
		Action(cmd, 200)
	case "x":
		log.Println("[ADB_Library] Click button 'x'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "52")
		Action(cmd, 200)
	case "y":
		log.Println("[ADB_Library] Click button 'y'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "53")
		Action(cmd, 200)
	case "z":
		log.Println("[ADB_Library] Click button 'z'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "54")
		Action(cmd, 200)

	case "A":
		ClickShift(getCharacter, false)
	case "B":
		ClickShift(getCharacter, false)
	case "C":
		ClickShift(getCharacter, false)
	case "D":
		ClickShift(getCharacter, false)
	case "E":
		ClickShift(getCharacter, false)
	case "F":
		ClickShift(getCharacter, false)
	case "G":
		ClickShift(getCharacter, false)
	case "H":
		ClickShift(getCharacter, false)
	case "I":
		ClickShift(getCharacter, false)
	case "J":
		ClickShift(getCharacter, false)
	case "K":
		ClickShift(getCharacter, false)
	case "L":
		ClickShift(getCharacter, false)
	case "M":
		ClickShift(getCharacter, false)
	case "N":
		ClickShift(getCharacter, false)
	case "O":
		ClickShift(getCharacter, false)
	case "P":
		ClickShift(getCharacter, false)
	case "Q":
		ClickShift(getCharacter, false)
	case "R":
		ClickShift(getCharacter, false)
	case "S":
		ClickShift(getCharacter, false)
	case "T":
		ClickShift(getCharacter, false)
	case "U":
		ClickShift(getCharacter, false)
	case "V":
		ClickShift(getCharacter, false)
	case "W":
		ClickShift(getCharacter, false)
	case "X":
		ClickShift(getCharacter, false)
	case "Y":
		ClickShift(getCharacter, false)
	case "Z":
		ClickShift(getCharacter, false)

	case "0":
		log.Println("[ADB_Library] Click button '0'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "7")
		Action(cmd, 200)
	case "1":
		log.Println("[ADB_Library] Click button '1'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "8")
		Action(cmd, 200)
	case "2":
		log.Println("[ADB_Library] Click button '2'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "9")
		Action(cmd, 200)
	case "3":
		log.Println("[ADB_Library] Click button '3'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "10")
		Action(cmd, 200)
	case "4":
		log.Println("[ADB_Library] Click button '4'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "11")
		Action(cmd, 200)
	case "5":
		log.Println("[ADB_Library] Click button '5'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "12")
		Action(cmd, 200)
	case "6":
		log.Println("[ADB_Library] Click button '6'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "13")
		Action(cmd, 200)
	case "7":
		log.Println("[ADB_Library] Click button '7'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "14")
		Action(cmd, 200)
	case "8":
		log.Println("[ADB_Library] Click button '8'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "15")
		Action(cmd, 200)
	case "9":
		log.Println("[ADB_Library] Click button '9'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "16")
		Action(cmd, 200)
	case "!":
		ClickShift(getCharacter, true)
	case "@":
		ClickShift(getCharacter, true)
	case "#":
		ClickShift(getCharacter, true)
	case "$":
		ClickShift(getCharacter, true)
	case "%":
		ClickShift(getCharacter, true)
	case "^":
		ClickShift(getCharacter, true)
	case "&":
		ClickShift(getCharacter, true)
	case "*":
		ClickShift(getCharacter, true)
	case "(":
		ClickShift(getCharacter, true)
	case ")":
		ClickShift(getCharacter, true)
	case "-":
		log.Println("[ADB_Library] Click button '69'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "69")
		Action(cmd, 200)
	case "_":
		ClickShift(getCharacter, true)
	case "=":
		log.Println("[ADB_Library] Click button '70'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "70")
		Action(cmd, 200)
	case "+":
		ClickShift(getCharacter, true)
	case "\\":
		log.Println("[ADB_Library] Click button '76'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "76")
		Action(cmd, 200)
	case "|":
		ClickShift(getCharacter, true)
	case ";":
		log.Println("[ADB_Library] Click button '74'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "74")
		Action(cmd, 200)
	case ":":
		ClickShift(getCharacter, true)
	case "'":
		log.Println("[ADB_Library] Click button '75'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "75")
		Action(cmd, 200)
	case "\"":
		ClickShift(getCharacter, true)
	case ",":
		log.Println("[ADB_Library] Click button '55'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "55")
		Action(cmd, 200)
	case ".":
		ClickShift(getCharacter, true)
	case "<":
		ClickShift(getCharacter, true)
	case ">":
		ClickShift(getCharacter, true)
	case "/":
		log.Println("[ADB_Library] Click button '73'")
		cmd := exec.Command("adb", "shell", "input", "keyevent", "73")
		Action(cmd, 200)
	case "?":
		ClickShift(getCharacter, true)
	case "`":
	case "~":
	}
}
