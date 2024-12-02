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

// Структура узла AVL дерева
type Node struct {
	value  []int
	left   *Node
	right  *Node
	height int
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Безопасное получение глубины дерева
func (n *Node) heightGet() int {
	if n == nil {
		return 0
	}
	return n.height
}

// Левый поворот
func (x *Node) leftRotate() *Node {
	y := x.right
	x.right = y.left
	y.left = x
	x.height = max(x.left.heightGet(), x.right.heightGet()) + 1
	y.height = max(y.left.heightGet(), y.right.heightGet()) + 1
	return y
}

// Правый поворот
func (y *Node) rightRotate() *Node {
	x := y.left
	y.left = x.right
	x.right = y
	y.height = max(y.left.heightGet(), y.right.heightGet()) + 1
	x.height = max(x.left.heightGet(), x.right.heightGet()) + 1
	return x
}

// Проверка разныцы глубин левого и равого дерева
func (n *Node) balanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.heightGet() - n.right.heightGet()
}

// Юалансировка дерева
func (n *Node) balance() *Node {
	// Баласировка
	bf := n.balanceFactor()

	// Левый-Левый поворот
	if bf > 1 && n.left.balanceFactor() >= 0 {
		return n.rightRotate()
	}
	// Левый-Правый поворот
	if bf > 1 && n.left.balanceFactor() < 0 {
		if n.left.right == nil {
			print()
		}
		n.left = n.left.leftRotate()
		return n.rightRotate()
	}
	// Правый-Правый поворот
	if bf < -1 && n.left.balanceFactor() <= 0 {
		return n.leftRotate()
	}
	// Левый-Правый поворот
	if bf < -1 && n.left.balanceFactor() > 0 {
		n.right = n.right.rightRotate()
		return n.leftRotate()
	}

	return n
}

func insert(n *Node, value int) *Node {
	// Если наше дерево было пустое до нового элемента
	if n == nil {
		newNote := &Node{height: 1}
		newNote.value = append(newNote.value, value)
		return newNote
	}

	if value < n.value[0] {
		// Добавление если число больше
		n.left = insert(n.left, value)
	} else if value > n.value[0] {
		// Добавление если число меньше
		n.right = insert(n.right, value)
	} else {
		// Добавление если число равно
		n.value = append(n.value, value)
		return n
	}

	// Изменение высоты родителя
	n.height = max(n.left.heightGet(), n.right.heightGet()) + 1

	return n.balance()
}

func getAndDeleteMinItem(n *Node) (*Node, int) {
	valueData := 0
	if n == nil {
		return nil, valueData
	}
	if n.left == nil {
		if len(n.value) == 1 {
			return n.right, n.value[0]
		} else {
			valueData = n.value[0]
			n.value = append(n.value[1:], n.value[1:]...)
			return n.left, valueData
		}
	}
	n.left, valueData = getAndDeleteMinItem(n.left)
	return n, valueData
}

// Функция для красивого вывода дерева
func printTree(node *Node, prefix string, isLeft bool) {
	if node != nil {
		// Вывод текущего узла
		if isLeft {
			fmt.Printf("%s├── %d\n", prefix, node.value)
		} else {
			fmt.Printf("%s└── %d\n", prefix, node.value)
		}

		// Создаем новый префикс для следующего уровня
		newPrefix := prefix
		if isLeft {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}

		// Рекурсивно вызываем для левого и правого поддерева
		printTree(node.left, newPrefix, true)
		printTree(node.right, newPrefix, false)
	}
}

func test() {
	var avl *Node

	testData := []int{1, 2, 3, 4, 5, 6, 7}
	for _, v := range testData {
		avl = insert(avl, v)
	}

	printTree(avl, "", false)
}

// Функция для подсчета количества узлов в дереве
func countNodes(node *Node) int {
	if node == nil {
		return 0
	}
	// Рекурсивно считаем узлы в левом и правом поддереве
	return 1 + countNodes(node.left) + countNodes(node.right)
}

// Не работает
func answerAVL(pathFile string) {
	// Открытия файла
	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var avlFirst *Node
	var avlSecond *Node

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
		avlFirst = insert(avlFirst, dataNum1)
		avlSecond = insert(avlSecond, dataNum2)
		// fmt.Println(dataNum1, dataNum2) // Выводим строку
	}

	var valueFirst, valueSecond, answer int
	var counterData int
	for avlFirst != nil {
		avlFirst, valueFirst = getAndDeleteMinItem(avlFirst)
		avlSecond, valueSecond = getAndDeleteMinItem(avlSecond)
		answer += absInt(valueFirst - valueSecond)
		counterData += 1
	}
	fmt.Println(answer)
}

func answerSimple(pathFile string) {
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
		answer += absInt(firstSlice[i] - secondSlice[i])
	}
	fmt.Println(answer)
}

func main() {
	// Получение абстрактного пути
	absPath, _ := os.Getwd()
	inputPath := "2024/first_day/input.txt"
	fullPath := filepath.Join(absPath, inputPath)

	answerAVL(fullPath)
	answerSimple(fullPath)
}
