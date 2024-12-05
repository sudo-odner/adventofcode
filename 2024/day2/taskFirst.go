package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func normalize(data int) int {
	if data < 0 {
		return -1
	} else if data == 0 {
		return 0
	}
	return 1
}

func taskFirst(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var dataReports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var dataReport []int
		line := scanner.Text()
		arrStr := strings.Split(line, " ")
		for _, value := range arrStr {
			numValue, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			dataReport = append(dataReport, numValue)
		}
		dataReports = append(dataReports, dataReport)
	}
	var answerCount int

	var statusFlag bool
	// Фдаг для определения возрастания или убывания чисел
	statusCheck := 0
	var distance int

	for _, report := range dataReports {
		statusFlag = true
		for idx := 0; idx < (len(report)-1) && statusFlag; idx++ {
			distance = report[idx] - report[idx+1]
			// Задаем убывание или возрастание
			if idx == 0 {
				statusCheck = normalize(distance)
			}
			// Условие возрастание/убывание, исключая одинаковые числа
			if statusCheck != normalize(distance) || normalize(distance) == 0 {
				statusFlag = false
				continue
			}
			// Условия что разница между уровнями от 1 до 3
			distance = distance * normalize(distance)
			if distance <= 0 || distance > 3 {
				statusFlag = false
				continue
			}
		}
		if statusFlag {
			answerCount++
		}
	}
	fmt.Print(answerCount)
}

func main() {
	absPath, _ := os.Getwd()
	dataFile := "/2024/day2/data.txt"
	fullPath := filepath.Join(absPath, dataFile)

	taskFirst(fullPath)
}
