package sudokusolver

import (
	"fmt"
	"sort"
)

type Cell struct {
	row int
	col int
	candidates []int
}

type SudokuSolver struct {
	board [9][9]int
}

func NewSudokuSolver(board [9][9]int) SudokuSolver {
	ss := SudokuSolver{}
	ss.board = board
	return ss
}

func (ss *SudokuSolver) Solve() [9][9]int {
	ss.backtrack()
	return ss.board
}

func (ss *SudokuSolver) PrettyPrintBoard() {
	fmt.Printf("+---+---+---+\n")
	for i := 0; i < 9; i++ {
		if i > 0 && i % 3 == 0 {
			fmt.Printf("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j % 3 == 0 {
				fmt.Printf("|")
			}
			fmt.Print(ss.board[i][j])
		}
		fmt.Println("|")

	}
	fmt.Printf("+---+---+---+\n")
}

func (ss *SudokuSolver) hasEmptyCells() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if ss.board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func (ss *SudokuSolver) getCandidates(row int, col int) []int {

	numbersTaken := map[int]struct{}{} // use map as a 'set' data structure

	// check rows and columns at (row, col)
	for i := 0; i < 9; i++ {
		numbersTaken[ss.board[row][i]] = struct{}{}
		numbersTaken[ss.board[i][col]] = struct{}{}
	}

	// check sub-grid containing (row,col)
	subRow := row / 3
	subCol := col / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r := subRow * 3 + i
			c := subCol * 3 + j
			numbersTaken[ss.board[r][c]] = struct{}{}
		}
	}

	candidates := make([]int, 0, 9)
	for i := 1; i <= 9; i++ {
		if _, ok := numbersTaken[i]; !ok {
			candidates = append(candidates, i)
		}
	}

	return candidates
}

func (ss *SudokuSolver) getAllCandidates() []Cell {
	cells := make([]Cell, 0, 81)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if ss.board[i][j] != 0 {
				continue
			}
			candidates := ss.getCandidates(i, j)
			if len(candidates) > 0 {
				cells = append(cells, Cell{
					row:        i,
					col:        j,
					candidates: candidates,
				})
			} else {
				return []Cell{}
			}
		}
	}
	// Sort so that paths with fewest candidates are explored first
	sort.Slice(cells, func(i int, j int) bool {
		return len(cells[i].candidates) < len(cells[j].candidates)
	})
	return cells
}

func (ss *SudokuSolver) backtrack() bool {
	if ss.hasEmptyCells() == false { // finish when no empty cells remain
		return true
	}

	cells := ss.getAllCandidates()
	for _, cell := range cells {
		candidates := cell.candidates
		for _, c := range candidates {
			ss.board[cell.row][cell.col] = c // try this path
			if ss.backtrack() {
				return true
			}
			ss.board[cell.row][cell.col] = 0 // reset if attempted path was unsuccessful
		}
		return false
	}
	return false
}
