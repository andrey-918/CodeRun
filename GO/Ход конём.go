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

	writer.WriteString(strconv.Itoa(horse(N, M)))
}

func horse( N, M int ) int {
	dp := make([][]int, N + 2)
	for i := range N + 2 {
		dp[i] = make([]int, M + 2)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}
	dp[2][2] = 1
	for i := 2; i < N + 2; i++ {
		for j := 2; j < M + 2; j++ {
			if i == 2 && j == 2 {
				continue
			}
			dp[i][j] = dp[i-1][j-2] + dp[i-2][j-1]
		}
	}

	return dp[N + 1][M + 1]
}