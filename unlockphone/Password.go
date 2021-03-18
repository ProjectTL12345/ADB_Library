package unlockphone

import "ADB_Library/src/utils"

func UnlockPassword(password string) {
	getCharsetLength := len(password)
	var count int

	for count = 0; count < getCharsetLength; count++ {
		utils.DictionaryCharacter(string(password[count]))
	}
}