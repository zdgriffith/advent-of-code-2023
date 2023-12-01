package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)


func part_a() {
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


func part_b() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	digit_map := [][]string{
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"},
		{"zero", "0"},
	}
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		idx_bounds := []int{len(line), -1}
		digit_bounds := []string{"", ""}
		for i := 0; i < len(digit_map); i++ {
			for j := 0; j < len(digit_map[i]); j++ {
				for k := 0; k < len(idx_bounds); k++ {
					if k == 0 {
						idx := strings.Index(line, digit_map[i][j])
						if idx != -1 && idx < idx_bounds[k] {
							idx_bounds[k] = idx
							digit_bounds[k] = digit_map[i][1]
						}
					} else {
						idx := strings.LastIndex(line, digit_map[i][j])
						if idx != -1 && idx > idx_bounds[k] {
							idx_bounds[k] = idx
							digit_bounds[k] = digit_map[i][1]
						}
					}
				}
			}
		}
		cal_val, _ := strconv.Atoi(digit_bounds[0] + digit_bounds[1])
		total += cal_val
	}
	fmt.Println(total)
}


func main() {
	part_a()
	part_b()
}