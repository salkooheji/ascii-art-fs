package main

import (
	functions "Ascii/functions"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	//extracting the file name from the args
	var shapeFileContent string
	if len(os.Args) == 3 || len(os.Args) == 2 {
		if len(os.Args) == 2 { // standard is the default
			shapeFileContent = functions.ReadShapeFile("standard.txt")
		} else if len(os.Args) == 3 {
			if functions.ValidateFileName(os.Args[2]) {
				shapeFileContent = functions.ReadShapeFile(os.Args[2])
			} else {
				log.Fatal("The file name received is invalid \nExpected files (standard.txt,shadow.txt, thinkertoy.txt)")
			}
		}
	} else {
		log.Fatal("Invalid Number of arguments")
	}

	charactersMap := functions.MapFileContentToRunes(os.Args[1], shapeFileContent)
	textToDraw := os.Args[1]
	functions.CheckAlphbets(textToDraw)
	if textToDraw == "" {
		return
	}
	// This will help in handling \n each index in textToDrawSplit will be considered a line
	textToDrawSplit := strings.Split(textToDraw, `\n`)

	if IsEmpty(textToDrawSplit) {
		textToDrawSplit = textToDrawSplit[1:]
	}

	for i, line := range textToDrawSplit {
		if line == "" {
			fmt.Println()
			continue
		}
		if line != "" {
			generateShape := functions.GenerateShape(line, charactersMap)
			fmt.Println(generateShape)

		}

		if i < len(textToDrawSplit)-1 {
			// fmt.Println()
		}
	}

}

func IsEmpty(arr []string) bool {
	for _, v := range arr {
		if v != "" {
			return false
		}
	}
	return true
}
