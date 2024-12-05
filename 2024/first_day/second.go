package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func countValueInSlice(array []int, target int) int {
	count := 0
	for _, value := range array {
		if value == target {
			count++
		}
	}
	return count
}

func answerSecond(pathFile string) {
	// Открытия файла
	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var firstSlice, secondSlice []int

	scanner := bufio.NewScanner(file)
	// Читаем файл построчно
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "   ")
		// Convert string to int
		dataNum1, err1 := strconv.Atoi(data[0])
		dataNum2, err2 := strconv.Atoi(data[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting:", err1, " , ", err2)
			return
		}
		firstSlice = append(firstSlice, dataNum1)
		secondSlice = append(secondSlice, dataNum2)
	}
	sort.Ints(firstSlice)
	sort.Ints(secondSlice)
	var answer int
	for i := 0; i < len(firstSlice); i++ {
		answer += firstSlice[i] * countValueInSlice(secondSlice, firstSlice[i])
	}
	fmt.Println(answer)
}

func main() {
	// Получение абстрактного пути
	absPath, _ := os.Getwd()
	inputPath := "2024/first_day/inputSecond.txt"
	fullPath := filepath.Join(absPath, inputPath)

	answerSecond(fullPath)
}
