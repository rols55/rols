package asciiart

import (
	"fmt"
	"os"
	"strings"
)

// parses input to exclude special non ascii table characters
func ParseInput(inputText string) string {
	invalid := false
	var err string
	for _, v := range inputText {
		if v < 10 || v > 10 && v < 13 || v > 13 && v < 32 || v > 126 {
			invalid = true
		}
		continue
	}
	if invalid {
		err = "Invalid characters detected"
	}
	return err
}

func AsciiFormat(inputText1 string, banner1 string) string {
	var modInput string // initialize functions's output, modInput is modified input

	//sisendist võib vist \r\n ka tulla?
	inputText := strings.Split(strings.ReplaceAll((inputText1), "\r\n", "\n"), "\n") // splits inputText into substrings by newline
	banner, err := os.ReadFile("./banners/" + banner1 + ".txt")
	if err != nil {
		fmt.Printf("Could not open the formatfile %s\n", err)
	}

	bannerSeparated := strings.Split(strings.ReplaceAll(string(banner), "\r\n", "\n"), "\n") // splits banner into substrings by newlines
	for i := 0; i < len(inputText); i++ {                                                    // loops through each element and subelement of inputText to make up rows in the style specified by banner and display them as output
		if inputText[i] == "" {
			modInput = modInput + "\n"
		} else {
			for rowCount := 0; rowCount < 9; rowCount++ { // banner's character height is 8 rows

				for j := 0; j < len(inputText[i]); j++ {

					modInput = modInput + bannerSeparated[((int(inputText[i][j])-32)*9)+rowCount] // populates a row with slices from banner
				}

				modInput += "\n"

			}
		}
	}
	Printout(modInput)
	return modInput
}

// creates a file and puts ascii art into it for users to download onto their machine
func Printout(modInput string) {

	var data []byte // since os.WriteFile writes bytes
	for _, v := range modInput {
		data = append(data, byte(v))
	}
	os.WriteFile("./download/download.txt", data, 0644)
}
