package unlockphone

import (
	"github.com/ProjectTL12345/ADB_Library/button"
	"github.com/ProjectTL12345/ADB_Library/utils"
)

func UnlockPin(code string) {
	codeLength := len(code)
	var count int

	for count = 0; count < codeLength; count++ {
		utils.DictionaryCharacter(string(code[count]))
	}

	button.EnterButton()
}