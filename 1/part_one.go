package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	left := make([]int, 0, 65535)
	right := make([]int, 0, 65535)

	for scanner.Scan() {
		str := scanner.Text()
		arr := strings.Split(str, "   ")

		num1, _ := strconv.Atoi(arr[0])
		num2, _ := strconv.Atoi(arr[1])

		left = append(left, num1)
		right = append(right, num2)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	total := 0.0
	for i, _ := range left {
		total += math.Abs(float64(right[i] - left[i]))
	}

	fmt.Println(int(total))
}
