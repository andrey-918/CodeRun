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
	nodes := make([][]int, N + 1)

	for i := 1; i < N + 1; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)

		for j := 1; j < N + 1; j++ {
			if x, _ := strconv.Atoi(parts[j - 1]); x == 1 {
				nodes[i] = append(nodes[i], j)
			}
		}
	}

	line, _ = reader.ReadString('\n')
	parts := strings.Fields(line)
	start, _ := strconv.Atoi(parts[0])
	finish, _ := strconv.Atoi(parts[1])

	dist, prev := minRoad(N, nodes, start, finish)
	writer.WriteString(strconv.Itoa(dist) + "\n")

	if dist > 0 {
		var answerRoad []int
		answerRoad = append(answerRoad, finish)
		curNode := finish
		for prev[curNode] != 0 {
			curNode = prev[curNode]
			answerRoad = append(answerRoad, curNode)
		}
		for i := len(answerRoad) - 1; i > -1; i-- {
			writer.WriteString(strconv.Itoa(answerRoad[i]) + " ")
		}
	}


}

func minRoad(N int, nodes [][]int, start int, finish int) (int, []int) {
	// Если начальная и конечная вершины совпадают
	if start == finish {
		return 0, []int{}
	}

	// Используем очередь для BFS
	queue := make([]int, 0)
	queue = append(queue, start)
	
	// Массив для хранения расстояний от start до каждой вершины
	dist := make([]int, N+1)
	for i := range dist {
		dist[i] = -1 // -1 означает, что вершина еще не посещена
	}
	dist[start] = 0

	// Массив для хранения информации о предке
	prev := make([]int, N+1)
	
	// BFS обход
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // Удаляем первый элемент
		
		// Обходим всех соседей текущей вершины
		for _, neighbor := range nodes[current] {
			// Если сосед еще не посещен
			if dist[neighbor] == -1 {
				dist[neighbor] = dist[current] + 1
				prev[neighbor] = current
				
				// Если достигли конечной вершины, возвращаем расстояние
				if neighbor == finish {
					return dist[neighbor], prev
				}
				
				queue = append(queue, neighbor)
			}
		}
	}
	
	// Если путь не найден
	return dist[finish], []int{}
}