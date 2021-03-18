package unlockphone

import "github.com/ProjectTL12345/ADB_Library/utils"

func UnlockPassword(password string) {
	getCharsetLength := len(password)
	var count int

	for count = 0; count < getCharsetLength; count++ {
		utils.DictionaryCharacter(string(password[count]))
	}
}