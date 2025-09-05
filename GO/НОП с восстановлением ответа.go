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

	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	line3, _ := reader.ReadString('\n')
	line4, _ := reader.ReadString('\n')

	N, _ := strconv.Atoi(strings.TrimSpace(line1))
	n_line := strings.Fields(line2)
	M, _ := strconv.Atoi(strings.TrimSpace(line3))
	m_line := strings.Fields(line4)

	writer.WriteString(NOP(N, n_line, M, m_line))
}

func NOP(N int, n_line []string, M int, m_line []string) string {
	var answer []string
	dp := make([][]int, N + 1)
	for i := range N + 1 {
		dp[i] = make([]int, M + 1)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}

	for i := 1; i < N + 1; i++ {
		for j := 1; j < M + 1; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if n_line[i - 1] == m_line[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}
		}
	}
	n := N
	m := M
	for n != 0 && m != 0 {
		if dp[n][m] == dp[n][m-1] {
			m--
		} else if dp[n][m] == dp[n - 1][m] {
			n--
		} else {
			answer = append(answer, n_line[n - 1])
			n--
			m--
		}
	}
	return strings.Join(Reverse(answer), " ")
}


func Reverse(s []string) []string {
    str := s
    for i, j := 0, len(str) - 1; i < j; i, j = i + 1, j - 1 {
        str[i], str[j] = str[j], str[i]
    }
    return str
}
