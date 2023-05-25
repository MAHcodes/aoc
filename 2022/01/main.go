package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
  var arr []int

  last := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
      arr = append(arr, last)
      last = 0
      continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		last += calories
	}

  highest := 0
  for _, num := range arr {
    if num > highest {
      highest = num
    }
  }

  fmt.Println(highest)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
