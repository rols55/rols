package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// add flag functionality to program
func flags() *string {
	outputFileName := flag.String("output", "", "Provides a name for file where program's output will be saved") // flag that saves program's output into new file specified by flag's value
	flag.Parse()
	return outputFileName
}

func main() {
	outputFileName := flags() //parses/scans flags in arguments and saves outputfilename into pointer
	inputText := ""           // text to format by banner style
	banner := ""              // banner name to specify style
	err := "Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --help to see all the possible flags"
	switch len(flag.Args()) {
	case 1: // case unspecified banner
		inputText = string(flag.Arg(0))
		banner = "standard.txt"
		printout(asciiFormat(inputText, banner), outputFileName)
	case 2: //case for specified
		switch string(flag.Arg(1)) {
		case "shadow", "standard", "thinkertoy":
			banner = flag.Arg(1) + ".txt"
		case "shadow.txt", "standard.txt", "thinkertoy.txt":
			banner = flag.Arg(1)
		default:
			fmt.Println(err)
			os.Exit(0)
		}
		inputText = string(flag.Arg(0))
		printout(asciiFormat(inputText, banner), outputFileName)
	default:
		fmt.Println(err)
	}
}

// prints inputText according to the specified banner format
func asciiFormat(inputText string, banner string) map[int]string {
	modInput := make(map[int]string) //makes map where keys are character height and values are rows of printable characters
	rowCount := 1                    // initializes row count at this step in order to display newline or exit from the program depending on circumstances
	switch {
	case inputText == `\n`:
		modInput[rowCount] = "\n"
	case inputText == "":
		os.Exit(0)
	default:
		inputText := strings.Split(string(inputText), `\n`) // splits inputText into substrings by newline
		banner, err := os.ReadFile(banner)
		if err != nil {
			fmt.Printf("Could not open the formatfile %s\n", err)
		}
		bannerSeparated := strings.Split(string(banner), "\n") // splits banner into substrings by newlines
		for i := 0; i < len(inputText); i++ {                  // loops through each element and subelement of inputText to make up rows in the style specified by banner and display them as output
			if inputText[i] == "" {
				modInput[rowCount] = "\n"
			} else {
				for rowCount = 1; rowCount < 9; rowCount++ { // banner's character height is 8 rows
					var row string // variable for storing row prinout info
					for j := 0; j < len(inputText[i]); j++ {
						if inputText[i][j] >= 32 && inputText[i][j] <= 126 { // controll of ascii
							row = row + bannerSeparated[((int(inputText[i][j])-32)*9)+rowCount] // populates a row with slices from banner
						}
					}
					modInput[rowCount] = row + "\n"
				}
			}
		}
	}
	return modInput
}

func printout(modInput map[int]string, outputFileName *string) {
	fileName := *outputFileName
	if fileName != "" {
		var data []byte
		for i := 0; i < 9; i++ {
			for _, v := range modInput[i] {
				data = append(data, byte(v))
			}
		}
		os.WriteFile(fileName, data, 0644)
	} else {
		for i := 1; i < 9; i++ {
			fmt.Print(modInput[i])
		}
	}
}
