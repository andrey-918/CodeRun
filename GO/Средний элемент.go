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
    
    a, _ := strconv.Atoi(parts[0])
    b, _ := strconv.Atoi(parts[1])
    c, _ := strconv.Atoi(parts[2])
    
    numbers := []int{a, b, c}
    sort(numbers)
    
    writer.WriteString(strconv.Itoa(numbers[1]))
	writer.WriteByte('\n')
}

func sort(nums []int) {
    if nums[0] > nums[1] {
        nums[0], nums[1] = nums[1], nums[0]
    }
    if nums[1] > nums[2] {
        nums[1], nums[2] = nums[2], nums[1]
    }
    if nums[0] > nums[1] {
        nums[0], nums[1] = nums[1], nums[0]
    }
}