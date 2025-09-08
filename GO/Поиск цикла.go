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
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(strings.TrimSpace(line))

		row := make([]int, len(fields))
		for j, field := range fields {
			num, _ := strconv.Atoi(field)
			row[j] = num
		}
		graph[i] = row
	}
	var path []int
	visited := make([]bool, n)
	answer := false
	for i := 0; i < n; i++ {
		if !visited[i] {
			answer = findCycleDFS(graph, i, visited, -1, &path)
			if answer {
				break
			}
		} 
	}

	if answer {
		writer.WriteString("YES\n")
		writer.WriteString(strconv.Itoa(len(path)) + "\n")
		for _, index := range path {
			writer.WriteString(strconv.Itoa(index + 1) + " ")
		}
	} else {
		writer.WriteString("NO")
		for _, index := range path {
			writer.WriteString(strconv.Itoa(index + 1) + " ")
		}
	}
}

func findCycleDFS(graph [][]int, index int, visited []bool, parent int, path *[]int) bool {
	visited[index] = true
	*path = append(*path, index)

	for ind, value := range graph[index] {
		if value == 1 {
			if !visited[ind] {
				if findCycleDFS(graph, ind, visited, index, path) {
					return true
				}
			} else if ind != parent {
				cycleStart := -1
                for i, v := range *path {
                    if v == ind {
                        cycleStart = i
                        break
                    }
                }
                if cycleStart != -1 {
                    *path = (*path)[cycleStart:]
                    return true
                }
            }
        }
    }
    
    *path = (*path)[:len(*path)-1]
    return false
}