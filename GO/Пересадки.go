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
	N, _ := strconv.Atoi(strings.TrimSpace(line))
	line, _ = reader.ReadString('\n')
	M, _ := strconv.Atoi(strings.TrimSpace(line))

	matrix := make([][]bool, M)
	for i := 0; i < M; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(strings.TrimSpace(line))
		P, _ := strconv.Atoi(parts[0])
		matrix[i] = make([]bool, N + 1)
		for j := 1; j < P + 1; j++ {
			station, _ := strconv.Atoi(parts[j])
			matrix[i][station] = true
		}
	}
	line, _ = reader.ReadString('\n')
	parts := strings.Fields(strings.TrimSpace(line))
	A, _ := strconv.Atoi(parts[0])
	B, _ := strconv.Atoi(parts[1])

	writer.WriteString(strconv.Itoa(transferCount(N, M, matrix, A, B)))
}

func transferCount(N, M int, matrix [][]bool, A, B int) int {
	if A == B {
		return 0
	}
	arr := make([]int, N + 1)
	for index := range arr {
		arr[index] = -1
	}
	queue := []int{A}
	for len(queue) > 0 {
		station := queue[0]
		queue = queue[1:]
		for _, line := range matrix {
			if !line[station] {
				continue
			}
			for neigbourStation, check := range line {
				if check && neigbourStation != station{
					if arr[neigbourStation] == -1 {
						arr[neigbourStation] = arr[station] + 1
						queue = append(queue, neigbourStation)
					}
					if neigbourStation == B {
						return arr[neigbourStation]
					}
				}
			}
		}
	}
	return arr[B]
}