package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	line = line[:len(line)-1] 

	if correctSeq(line) {
		writer.WriteString("yes")
	} else {
		writer.WriteString("no")
	}
}

func correctSeq(sequence string) bool {
    dict := map[rune]rune{'(':')', '[':']', '{':'}'}
    var stack []rune
    
    for _, char := range sequence {
        if char == '(' || char == '[' || char == '{' {
            stack = append(stack, char)
        } else {
            if len(stack) == 0 || dict[stack[len(stack)-1]] != char {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    
    return len(stack) == 0
}