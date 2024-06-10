package Ascii
import "log"
func CheckAlphbets (text string) {
	for _,char := range text{
		if char<32 || char>126{
			log.Fatal("The characters are not valid you can use only English bits")
		}
	}
}