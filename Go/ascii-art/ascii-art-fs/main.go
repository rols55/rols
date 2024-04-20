package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	banner := "" // banner name to specify style
	switch len(os.Args) {
	case 2: // case unspecified banner
		banner = "standard.txt"
		asciiFormat(os.Args[1], banner)
	case 3: //case for specified
		switch string(os.Args[2]) {
		case "shadow", "standard", "thinkertoy":
			banner = os.Args[2] + ".txt"
		case "shadow.txt", "standard.txt", "thinkertoy.txt":
			banner = os.Args[2]
		default:
			fmt.Println("Wrong format name")
			os.Exit(0)
		}
		asciiFormat(os.Args[1], banner)
	default:
		fmt.Println("\nPlease read the README once again: number of args or banner is wrong!")
	}
}

// prints inputText according to the specified banner format
func asciiFormat(inputText string, banner string) {

	switch { //check for newlines and empty strings
	case inputText == `\n`:
		fmt.Println()
	case inputText == "":
		os.Exit(0)
	default:
		inputText := strings.Split(strings.ReplaceAll(inputText, "\\n", "\n"), "\n") // splits inputText into substrings by newline
		banner, err := os.ReadFile(banner)
		if err != nil {
			fmt.Printf("Could not open the formatfile %s\n", err)
		}
		bannerSeparated := strings.Split(string(banner), "\n") // splits banner into substrings by newlines
		for i := 0; i < len(inputText); i++ {                  // loops through each element and subelement of inputText to make up rows in the style specified by banner and display them as output
			if inputText[i] == "" {
				fmt.Println()
			} else {
				for rowCount := 1; rowCount < 9; rowCount++ { // banner's character height is 8 rows
					var row string // variable for storing row prinout info
					for j := 0; j < len(inputText[i]); j++ {
						if inputText[i][j] >= 32 && inputText[i][j] <= 126 { // controll of ascii
							row = row + bannerSeparated[((int(inputText[i][j])-32)*9)+rowCount] // populates a row with slices from banner
						}
					}
					fmt.Println(row)
				}
			}
		}
	}
}
