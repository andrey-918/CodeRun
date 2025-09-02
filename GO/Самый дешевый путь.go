package main

import (
	"os"
	"bufio"
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

	
	matrix := make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, M)
	}
	

	for n := range N {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		for m := range M {
			num, _ := strconv.Atoi(parts[m])
			matrix[n][m] = num
		}
	}

	writer.WriteString(strconv.Itoa(minWay(N, M, matrix)))
}

func minWay(N, M int, matrix [][]int) int {
	const INF = 1e10 
	dp := make([][]int, N + 1)
	for i := range N + 1 {
		dp[i] = make([]int, M + 1)
		for j := range dp[i] {
			dp[i][j] = INF 
		}
	}
	dp[1][1] = matrix[0][0]
	for i := 1; i < N + 1; i++ {
		for j := 1; j < M + 1; j++ {
			if i == 1 && j == 1 {
				continue
			}
			dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + matrix[i - 1][j - 1]
		}
	}
	return dp[N][M]
}