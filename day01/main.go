package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		digits := ""
		for _, value := range scanner.Text() {
			if unicode.IsDigit(value) {
				digits += string(value)
			}
		} 
		cal_val, _ := strconv.Atoi(string(digits[0]) + string(digits[len(digits) - 1]))
		total += cal_val
	}
	fmt.Println(total)
}
