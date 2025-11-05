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

	answers := make([]int, n)
	
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(strings.TrimSpace(line))
		k, _ := strconv.Atoi(fields[0])
		arr := make([]float64, k)
		for j := 1; j < k + 1; j++ {
			val, _ := strconv.ParseFloat(fields[j], 64)
        	arr[j-1] = float64(val)
		}
		if sortConv(arr) {
			answers[i] = 1
		}
	}

	for _, answer := range answers {
		writer.WriteString(strconv.Itoa(answer) + "\n")
	}
}

func sortConv(priorities []float64) bool {
	var sklad []float64
	var cehB []float64
	for _, priority := range priorities {
		for len(cehB) > 0 && cehB[len(cehB)-1] < priority {
			if len(sklad) == 0 || sklad[len(sklad) - 1] <= cehB[len(cehB)-1] {
				sklad = append(sklad, cehB[len(cehB)-1])
				cehB = cehB[:len(cehB)-1]
			} else {
				return false
			}
		}
		cehB = append(cehB, priority)
	}
	return len(cehB) == 0 || len(sklad) == 0 || cehB[len(cehB)-1] >= sklad[len(sklad)-1]
}