// Package main solves the problem of calculating the sum of shortest paths for fleas on a chessboard
// using knight moves from a starting position. If any flea is unreachable, returns -1.
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Coords represents the coordinates on the board (row i, column j).
type Coords struct {
	i int
	j int
}

// moves defines the possible knight moves in chess (8 directions).
var moves = []Coords{{1, 2}, {2, 1}, {1, -2}, {-1, 2}, {-1, -2}, {-2, 1}, {-2, -1}, {2, -1}}

const inf = 1000000000000000

func main() {
	// Use buffered reader and writer for efficient I/O.
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// Read the first line: N (rows), M (columns), S (start row), T (start column), Q (number of fleas).
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0]) // Board rows (1 to N)
	M, _ := strconv.Atoi(parts[1]) // Board columns (1 to M)
	S, _ := strconv.Atoi(parts[2]) // Start row (1-based)
	T, _ := strconv.Atoi(parts[3]) // Start column (1-based)
	Q, _ := strconv.Atoi(parts[4]) // Number of fleas

	// Read Q lines, each with flea position (i, j).
	var fleas []Coords
	for i := 0; i < Q; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		fleaI, _ := strconv.Atoi(parts[0]) // Flea row (1-based)
		fleaJ, _ := strconv.Atoi(parts[1]) // Flea column (1-based)
		fleas = append(fleas, Coords{fleaI, fleaJ})
	}

	// Compute and output the sum of distances or -1 if any unreachable.
	writer.WriteString(strconv.Itoa(fleasDist(N, M, S, T, Q, fleas)))
}

// fleasDist calculates the sum of shortest distances each flea to finish.
// Returns -1 if any flea is unreachable.
func fleasDist(N, M, S, T, Q int, fleas []Coords) int {
	// Compute distance matrix from start position.
	desk := fleaDesk(N, M, S, T)
	sumDist := 0

	// Sum distances for each flea; return -1 if any is unreachable.
	for i := 0; i < Q; i++ {
		fleaI, fleaJ := fleas[i].i, fleas[i].j
		if desk[fleaI][fleaJ] == -1 {
			return -1
		}
		sumDist += desk[fleaI][fleaJ]
	}

	return sumDist
}

// fleaDesk computes the shortest distance from (S, T) to all reachable squares using BFS.
// Returns a 2D slice where desk[i][j] is the distance or -1 if unreachable.
func fleaDesk(N, M, S, T int) [][]int {
	// Initialize distance matrix (1-based indexing).
	desk := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		desk[i] = make([]int, M+1)
		for j := 1; j <= M; j++ {
			desk[i][j] = -1 // Unreachable initially
		}
	}

	// Start BFS from (S, T).
	desk[S][T] = 0
	queue := make([]Coords, 0)
	queue = append(queue, Coords{S, T})

	// BFS to compute distances.
	for len(queue) > 0 {
		curCoords := queue[0]
		curI, curJ := curCoords.i, curCoords.j
		queue = queue[1:] // Dequeue

		// Try all knight moves.
		for _, move := range moves {
			newI := curI + move.i
			newJ := curJ + move.j
			if lineCheck(N, M, newI, newJ) && desk[newI][newJ] == -1 {
				// Enqueue and update distance.
				queue = append(queue, Coords{newI, newJ})
				desk[newI][newJ] = desk[curI][curJ] + 1
			}
		}
	}
	return desk
}

// lineCheck verifies if the position (i, j) is within the board boundaries (1 to N, 1 to M).
func lineCheck(N, M, i, j int) bool {
	return i >= 1 && i <= N && j >= 1 && j <= M
}
