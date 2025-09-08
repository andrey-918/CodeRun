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

	for i := 0; i < M; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		if !inArr(graph[num1], num2) {
			graph[num1] = append(graph[num1], num2)
		}
	}
	hasCycle := false
	visited := make([]bool, N + 1)
	inStack := make([]bool, N + 1)
	for i := 1; i < N + 1; i++ {
		if !visited[i] {
			hasCycle = hasCycleDFS(graph, i, visited, inStack)
			if hasCycle {
				break
			}
		}
	}
	if hasCycle {
		writer.WriteString("-1")
	} else {
		var stack []int
		visited := make([]bool, N + 1)
		for i := 1; i < N + 1; i++ {
			if !visited[i]  {
				topsort(graph, i, visited, &stack)
			}
		}

		for i := len(stack) - 1; i >= 0; i-- {
			writer.WriteString(strconv.Itoa(stack[i]) + " ")
		}
	}
}

func inArr(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}

func topsort(graph [][]int, index int, visited []bool, stack *[]int) {
	visited[index] = true
	for _, i := range graph[index] {
		if !visited[i] {
			topsort(graph, i, visited, stack)
		}
	}
	*stack = append(*stack, index)
}

func hasCycleDFS(graph [][]int, index int, visited []bool, inStack []bool) bool {
    if inStack[index] {
        return true 
    }
    if visited[index] {
        return false 
    }
    
    visited[index] = true
    inStack[index] = true
    
    for _, neighbor := range graph[index] {
        if hasCycleDFS(graph, neighbor, visited, inStack) {
            return true
        }
    }
    
    inStack[index] = false 
    return false
}