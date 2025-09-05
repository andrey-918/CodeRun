package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	arr := make([][]int, n + 1)

	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)

		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		
		if !IndexFunc(arr[num1], num2) && num1 != num2 {
			arr[num1] = append(arr[num1], num2)
			arr[num2] = append(arr[num2], num1)
		}
	}

	visited := make([]bool, n + 1)

	DFS(arr, 1, visited)

	var count int
	for _, num := range visited {
		if num {
			count++
		}
	}
	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')
	
	for i, num := range visited {
		if num == true {
			if i != 1 {
				writer.WriteByte(' ')
			}
			writer.WriteString(strconv.Itoa(i))
		}
	}
}

func IndexFunc(slice []int, target int) bool {
    for _, num := range slice {
        if num == target {
            return true
        }
    }
    return false
}

func DFS(graph [][]int, index int, visited []bool) {
	visited[index] = true
	for _, num := range graph[index] {
		if !visited[num] {
			DFS(graph, num, visited)
		}
	}
}