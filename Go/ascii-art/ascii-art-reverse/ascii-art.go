package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// inits an object for using flags more conveniently
type programFlags struct {
	output  *string
	reverse *string
}

// inits flags, saves them to object for later use and parses flags
func flags() programFlags {
	initFlags := programFlags{
		output:  flag.String("output", "", "Provides a name for file where program's output will be saved"),
		reverse: flag.String("reverse", "", "Provides a file's name for the program from which ascii art will be converted to string"),
	}
	flag.Parse()
	return initFlags
}

func main() {
	printout(flags())
}

// function to display result of the program
func printout(f programFlags) {
	switch { //checks the flags first, if they are empty proceeds to programs basic function
	case *f.output != "": // outputs program's output into a file
		f.fileOutput(basic())
	case *f.reverse != "": // reads modified file and outputs its contents into standard output
		if flag.Arg(0) != "" {
			fmt.Println(f.reverseAscii(mapFromBanner(separateBanner(readBanner(flag.Arg(0) + ".txt")))))
		}
		fmt.Println(f.reverseAscii(mapFromBanner(separateBanner(readBanner("standard.txt")))))
	default: // executes program's basic function
		for i := range basic() {
			fmt.Print(basic()[i])
		}
	}
}

// basic functionality of the program; it returns inputtext specified by banner style in ascii format
func basic() []string {
	inputText := "" // text to format by banner style
	banner := ""    // banner name to specify style
	err := "Usage: go run . [OPTION]* [STRING] [BANNER]\n\nEX: go run . --help to see all the possible flags\n\n*Option is optional"
	var modInput []string
	switch len(flag.Args()) {
	case 1: // case unspecified banner
		inputText = string(flag.Arg(0))
		banner = "standard.txt"
		modInput = asciiFormat(inputText, banner)
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
		modInput = asciiFormat(inputText, banner)
	default:
		fmt.Println(err)
	}
	return modInput
}

// prints input text according to the specified banner format
func asciiFormat(inputText string, banner string) []string {
	var modInput []string // initialize functions's output, modInput is modified input
	switch {
	case inputText == `\n`:
		modInput = append(modInput, "\n")
	case inputText == "":
		os.Exit(0)
	default:
		inputText := strings.Split(string(inputText), `\n`)   // splits inputText into substrings by newline
		bannerSeparated := separateBanner(readBanner(banner)) // splits banner into substrings by newlines
		for i := 0; i < len(inputText); i++ {                 // loops through each element and subelement of inputText to make up rows in the style specified by banner and display them as output
			if inputText[i] == "" {
				modInput = append(modInput, "\n")
			} else {
				for rowCount := 1; rowCount < 9; rowCount++ { // banner's character height is 8 rows
					var row string // variable for storing row prinout info
					for j := 0; j < len(inputText[i]); j++ {
						if inputText[i][j] >= 32 && inputText[i][j] <= 126 { // controll of ascii
							row = row + bannerSeparated[((int(inputText[i][j])-32)*9)+rowCount] // populates a row with slices from banner
						}
					}
					modInput = append(modInput, row+"\n")
				}
			}
		}
	}
	return modInput
}

func readBanner(banner string) []byte {
	readBanner, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
	}
	return readBanner
}

// separates banner file by newline fo furhter use
func separateBanner(banner []byte) []string {
	sliceFromBanner := strings.Split(string(banner), "\n")
	return sliceFromBanner
}

// makes map out of banner in ascnending order of ascii table of characters
func mapFromBanner(bannerSeparated []string) map[rune]string {
	asciiMap := make(map[rune]string)
	key := rune(32)
	row := 1
	rowValue := ""
	for _, v := range bannerSeparated[0:] {
		if row != 9 {
			rowValue += v
			row++
		} else {
			asciiMap[key] = rowValue
			row = 1
			rowValue = ""
			key++
		}

	}
	return asciiMap
}

// method to put output of the program into a file
func (f programFlags) fileOutput(input []string) {
	var data []byte
	for i := range input {
		for _, v := range input[i] {
			data = append(data, byte(v))
		}
	}
	os.WriteFile(*f.output, data, 0644)
}

// reads file and outputs its ascii styled text to standard output as regular characters
func (f programFlags) reverseAscii(asciiMap map[rune]string) string {
	content, err := os.ReadFile(*f.reverse + ".txt") // reads file
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [BANNER]*\n\nEX: go run . --reverse=<fileName>\n\n*Specifying banner style is necessary if input file is in different style than standard")
	}
	contentRows := strings.Split(string(content), "\n") // separates content into rows
	var origin string                                   // fucntion's output
	var charLen int                                     // width of a character from the map
	var compare string                                  // median variable that' used to make comparison with the character from the map
	for l := 0; len(contentRows[0]) > 1; l++ {          // comapres character from the map with character from the file if they are same then stores key (which is characters runic value) in function's output
		for i, v := range asciiMap { // compares every character from map with a character form file
			charLen = len(v) / 7                          // gets width of the character form the map
			for rowCount := 0; rowCount < 7; rowCount++ { // saves a character from file into compare for comparison with the maps character value; file's character is is determined by the width of the map's character
				for j := 0; j < charLen; j++ {
					if charLen > len(contentRows[rowCount]) { // after finding a match, file is shortened by the found character width. In order for the program not to crash with an error when last file's character width is less than actual character's width from the map we set file's character index to width of remaining characters
						charLen = len(contentRows[rowCount]) - 1
					}
					compare += string(contentRows[rowCount][j])
				}
			}
			if v == compare { // compares map's character with file's character, if they are the same then shorten file's width by the character's with
				origin += string(i)
				for rowCount := 0; rowCount < 7; rowCount++ {
					contentRows[rowCount] = contentRows[rowCount][charLen:]
				}
			}
			compare = ""                  // resets varialbe
			i = ' '                       // resets variable
			if len(contentRows[0]) == 0 { // if file becomes and empty slice then exit the loop
				break
			}
		}
	}
	return origin
}