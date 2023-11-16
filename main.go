package main

import (
	"bufio"
	"fmt"
	"github.com/dustinbowers/sudoku-solver/sudokusolver"
	"os"
	"strconv"
	"time"
)

func main() {
	var board [9][9]int
	scanner := bufio.NewScanner(os.Stdin)

	// Collect user input
	for i := 0; i < 9; i++ {
		fmt.Printf("Input Sudoku row #%d: ", i + 1)
		scanner.Scan()
		line := scanner.Text()
		if len(line) != 9 {
			fmt.Println("Error: Input line must be exactly 9 numbers.")
			return
		}
		for j, c := range line {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				fmt.Println("Error: Bad Sudoku board input.")
				return
			}
			board[i][j] = n
		}
	}

	// Solve
	ss := sudokusolver.NewSudokuSolver(board)
	fmt.Printf("Solving...\n")

	start := time.Now()
	ss.Solve()
	elapsed := time.Since(start)

	fmt.Printf("\nSolution found in %s:\n\n", elapsed)
	ss.PrettyPrintBoard()
}
