package Ascii

func ValidateFileName(fileName string) bool {
	if fileName == "standard.txt" || fileName == "shadow.txt" || fileName == "thinkertoy.txt" {
		return true
	}
	return false
}
