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
	

	for i := range N {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		for j := range M {
			num, _ := strconv.Atoi(parts[j])
			matrix[i][j] = num
		}
	}

	count, steps := maxWay(N, M, matrix)

	writer.WriteString(strconv.Itoa(count))
	writer.WriteString("\n")
	writer.WriteString(steps)
}

func maxWay(N, M int, matrix [][]int) (int, string) {
	const INF = 1e10

	dp := make([][]int, N + 1)
	for i := range N + 1 {
		dp[i] = make([]int, M + 1)
		for j := range M + 1 {
			dp[i][j] = -INF
		}
	}

	dp[1][1] = matrix[0][0]
	for i := 1; i < N + 1; i++ {
		for j := 1; j < M + 1; j++ {
			if i == 1 && j == 1 {
				continue
			}
			dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]) + matrix[i - 1][j - 1]
		}
	}

	var steps string
	i := N
	j := M
	for i > 1 && j > 1 {
		if i != N || j != M {
			steps += " "
		}
		if dp[i - 1][j] > dp[i][j - 1] {
			steps += "D"
			i--
		} else {
			steps += "R"
			j--
		}
		
	}
	for k := i - 1; k > 0; k-- {
		steps += " "
		steps += "D"
	}

	for k := j - 1; k > 0; k-- {
		steps += " "
		steps += "R"
	}

	return dp[N][M], Reverse(steps)
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}