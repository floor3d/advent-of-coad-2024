package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	b, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	

	r := regexp.MustCompile(`mul\(\d+,\d+\)`)

	arr := r.FindAllString(str, -1)

	total := 0

	for _, mul := range arr {
		sanitized := mul
		sanitized = strings.ReplaceAll(sanitized, "(", " ")
		sanitized = strings.ReplaceAll(sanitized, ")", " ")
		sanitized = strings.ReplaceAll(sanitized, ",", " ")

		parts := strings.Fields(sanitized)

		num1, _ := strconv.Atoi(parts[1])
		num2, _ := strconv.Atoi(parts[2])

		total += num1 * num2

	}
	fmt.Println(total)
}

