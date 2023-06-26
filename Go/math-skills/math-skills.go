package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func average(array []float64) int { //Sums up every value and divides by len(array)
	var sum float64 = 0
	for i := range array {
		sum += array[i]
	}
	return int(math.Round(sum / float64(len(array))))
}

func median(array []float64) int { //Finds middle value or average of two middle values
	half := array[len(array)/2]
	second_half := array[len(array)/2-1]
	if len(array)%2 == 0 { //If len(array) is even
		return int(math.Round((half + second_half) / 2))
	} else { //If len(array) is odd
		return int(math.Round(half))
	}
}

func variance(array []float64, average int) int {
	var sum float64 = 0
	for _, i := range array { //Takes value subtracts average and goes to the power of 2
		sum += math.Pow(i-float64(average), 2)
	}
	return int(math.Round(sum / float64(len(array)))) //Divides sum by len(array)
}

func standard_deviation(variance int) int { //Sqrt of variance
	return int(math.Round(math.Sqrt(float64(variance))))
}

func main() {
	var array []float64
	input, input_err := os.Open(os.Args[1]) //Opens file of argument [1]
	if input_err != nil {
		log.Fatal(input_err)
	}
	scanned := bufio.NewScanner(input) //Buffers strings
	scanned.Split(bufio.ScanLines)
	for scanned.Scan() {
		number, conv_err := strconv.ParseFloat(scanned.Text(), 64) //Converts string to float64
		if conv_err != nil {
			log.Fatal(conv_err)
		}
		array = append(array, number) //Adds float to array
	}
	sort.Float64s(array) //Sorting to increasing order for median

	//Final prints
	fmt.Println("Average:", average(array))
	fmt.Println("Median:", median(array))
	fmt.Println("Variance:", variance(array, average(array)))
	fmt.Println("Standard deviation:", standard_deviation(variance(array, average(array))))
}
