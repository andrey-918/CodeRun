package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	graph := make([][]int, N + 1)

	for i := 0; i < M; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		if !inArr(graph[num1], num2) && num1 != num2 {
			graph[num1] = append(graph[num1], num2)
			graph[num2] = append(graph[num2], num1)
		}
	}

	var answer [][]int
	visited := make([]bool, N + 1)
	for i := 1; i < N + 1; i++ {
		if !visited[i] {
			var order []int
			DFS(graph, i, visited, &order)
			answer = append(answer, order)
		}
	}
	writer.WriteString(strconv.Itoa(len(answer)) + "\n")
	for _, order := range answer {
		sort.Ints(order)
		writer.WriteString(strconv.Itoa(len(order)) + "\n")
		for _, num := range order {
			writer.WriteString(strconv.Itoa(num) + " ")
		}
		writer.WriteString("\n")
	}
}

func inArr(arr []int, target int) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}

func DFS(graph [][]int, index int, visited []bool, order *[]int) {
	visited[index] = true
	*order = append(*order, index)
	for _, num := range graph[index] {
		if !visited[num] {
			DFS(graph, num, visited, order)
		}
	}
}