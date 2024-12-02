package main

import (
	"bufio"
	"fmt"
	"log"
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

	m := make(map[int]int)

	for scanner.Scan() {
		str := scanner.Text()
		arr := strings.Split(str, "   ")

		num1, _ := strconv.Atoi(arr[0])
		num2, _ := strconv.Atoi(arr[1])

		left = append(left, num1)
		_, ok := m[num2]
		if !ok {
			m[num2] = 1
		} else {
			m[num2] += 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)

	total := 0
	for i, _ := range left {
		total += (left[i] * m[left[i]])
	}

	fmt.Println(int(total))
}
