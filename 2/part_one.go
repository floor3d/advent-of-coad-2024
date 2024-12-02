package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	amt_gud := 0

	for scanner.Scan() {
		str := scanner.Text()
		arr := strings.Split(str, " ")

		nums_arr := make([]int, 0, 65535)

		for _, value := range arr {
			val, _ := strconv.Atoi(value)
			nums_arr = append(nums_arr, val)
		}

		past := 0
		up := true
		good := true
		for index, value := range nums_arr {

			if index == 1 {
				if past > value {
					up = false
				}
			}

			if index > 0 {
				if past == value || (up && past > value) || (!up && past < value) || math.Abs(float64(past-value)) > 3.0 {
					good = false
					break
				}
			}

			past = value

		}

		if good {
			amt_gud += 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(amt_gud)
}
