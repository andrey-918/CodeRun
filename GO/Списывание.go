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
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	graph := make([][]int, N + 1)
	group := make([]int, N + 1)

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

	answer := 1

	for i := 1; i < N + 1; i++ {
		if group[i] == 0 {
			groupSorting(graph, i, group, 1, &answer)
			if answer == 0 {
				break
			}
		}
	}
	if answer == 1 {
		writer.WriteString("YES")
	} else {
		writer.WriteString("NO")
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

func groupSorting(graph [][]int, index int, group []int, groupID int, answer *int) {
	group[index] = groupID
	for _, person := range graph[index] {
		if group[person] == 0 {
			groupSorting(graph, person, group, groupID % 2 + 1, answer)
		} else if group[person] != groupID % 2 + 1 {
			*answer = 0
		}
	}
}