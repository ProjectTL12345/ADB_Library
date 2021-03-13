package unlockphone

import "ADB_Library/src/utils"

func UnlockPin(code string) {
	codeLength := len(code)
	var count int

	for count = 0; count < codeLength; count++ {
		utils.DictionaryCharacter(string(code[count]))
	}
}