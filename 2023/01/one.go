package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Helper for handling errors
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Extract the numbers from a string and return it
func getNumbers(s string) (n []string) {
	r := regexp.MustCompile("[0-9]")
	return r.FindAllString(s, -1)
}

func main() {
	var res int

	// open the input file
	f, err := os.Open("input.txt")
	check(err)

	// close the input file
	defer f.Close()

	// create a scanner for the file
	scanner := bufio.NewScanner(f)

	// read the file line by line
	for scanner.Scan() {
		numbers := getNumbers(scanner.Text())

    // get the number
    first := numbers[0]
    last := numbers[len(numbers)-1]
		number := first + last

		// parse the number to int type
		value, err := strconv.Atoi(number)
		check(err)

		// add it to the result
		res += value
	}

	check(scanner.Err())

	fmt.Println(res)
}
