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

	charactersMap := functions.MapFileContentToRunes(shapeFileContent)
	textToDraw := os.Args[1]
	// This will help in handling \n each index in textToDrawSplit will be considered a line  
	textToDrawSplit  := strings.Split(textToDraw,`\n`)

	for i, line := range textToDrawSplit {
		if line != ""{
			generateShape := functions.GenerateShape(line,charactersMap)
			fmt.Println(generateShape)

		}

		if(i<len(textToDrawSplit)-1){
			fmt.Println()
		}
	}

}
