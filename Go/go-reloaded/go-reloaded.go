package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func keyword(index int) { //Finds the given keyword and runs corresponding function
	switch str_array[index] {
	case "(hex)":
		encode(index-1, 16)
	case "(bin)":
		encode(index-1, 2)
	case "(up)":
		single_conv(index, strings.ToUpper)
	case "(up,":
		multi_conv(index, strings.ToUpper)
	case "(low)":
		single_conv(index, strings.ToLower)
	case "(low,":
		multi_conv(index, strings.ToLower)
	case "(cap)":
		single_conv(index, strings.Title) //Deprecated from standard library but still works
	case "(cap,":
		multi_conv(index, strings.Title) //Used to not import cases

	}
}

func encode(index int, base int) { //Converts string to int and to given base
	conv, conv_err := strconv.ParseInt(str_array[index], base, 64)
	if conv_err != nil {
		log.Fatal(conv_err)
	}
	str_array[index] = strconv.Itoa(int(conv)) //Back to string conv
	cleanup(index+1, 1)
}

func single_conv(index int, f func(string) string) { //Converts single previous index
	str_array[index-1] = f(str_array[index-1])
	cleanup(index, 1)
}

func multi_conv(index int, f func(string) string) { //Converts n of previous indexies
	n, _ := strconv.Atoi(strings.Trim(str_array[index+1], ")"))
	for i := index - 1; i >= index-n; i-- {
		str_array[i] = f(str_array[i])
	}
	cleanup(index, 2)
}

func is_keyword(s string) bool { //Compares string to keywords
	keywords := []string{"(hex)", "(bin)", "(up)", "(up,", "(low)", "(low,", "(cap)", "(cap,"}
	for _, e := range keywords {
		if s == e {
			return true
		}
	}
	return false
}

func cleanup(index int, count int) { //Cleans up str_array of keywords
	temp := append(str_array[:index], str_array[index+count:]...)
	str_array = temp
}

func has_punc(index int) bool { //Compares string to punctuations
	punctuation := []string{".", ",", "!", "?", ":", ";"}
	for _, e := range punctuation {
		if string(str_array[index][0]) == e {
			return true
		}
	}
	return false
}

func fix_punc(index int) { //Moves punctuation to previous index and rewrites current index without punctuation
	if len(str_array[index]) > 1 {
		str_array[index-1] = str_array[index-1] + string(str_array[index][0])
		str_array[index] = str_array[index][1:]
	} else {
		str_array[index-1] = str_array[index-1] + str_array[index]
		cleanup(index, 1)
	}
}

func vowel(index int) { //Adds "n" if next index was vowel
	str_array[index] = str_array[index] + "n"
}

func quotes(index int) { //Checks quotes position and corrects them
	if index == len(str_array)-1 {
		str_array[index-1] = str_array[index-1] + str_array[index]
		cleanup(index, 1)
	} else {
		str_array[index] = str_array[index] + str_array[index+1]
		cleanup(index+1, 1)
	}
}

var str_array []string //Global variable of string slice

func main() {
	file, read_err := os.ReadFile(os.Args[1]) //Reads whole file
	if read_err != nil {
		log.Fatal(read_err)
	}

	str_array = strings.Fields(string(file)) //Assigns to global variable as slice of strings

	for i := 0; i < len(str_array); i++ { //Main loop
		switch {
		case is_keyword(str_array[i]):
			keyword(i)
			i = 0
		case has_punc(i):
			fix_punc(i)
			i = 0
		case (str_array[i] == "a" || str_array[i] == "A") && strings.ContainsAny(string(str_array[i+1][0]), "aeiouh"):
			vowel(i)
		case str_array[i] == "'" || str_array[i] == "`":
			quotes(i)
		}
	}

	write_err := os.WriteFile(os.Args[2], []byte(strings.Join(str_array, " ")), 0644) // Writes(creates) to file
	if write_err != nil {
		log.Fatal(write_err)
	} else {
		fmt.Println("Success!")
	}

}
