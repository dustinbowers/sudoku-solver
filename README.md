# Sudoku Solver
Solve a user-provided Sudoku Board using a recursive backtracking algorithm

## Usage
```go
import "fmt"
import "github.com/dustinbowers/sudku-solver/sudokusolver"

func main() {
  board := [9][9]int{ /* ... */ }
  
  ss := sudokusolver.NewSudokuSolver(board)
  _ = ss.Solve() // Solution as [9][9]int

  ss.PrettyPrintBoard()
}
```
## CLI Example 
Using ["the world's hardest sudoku"](https://abcnews.go.com/blogs/headlines/2012/06/can-you-solve-the-hardest-ever-sudoku) as input:
```
$ go run main.go
Input Sudoku row #1: 800000000
Input Sudoku row #2: 003600000
Input Sudoku row #3: 070090200
Input Sudoku row #4: 050007000
Input Sudoku row #5: 000045700
Input Sudoku row #6: 000100030
Input Sudoku row #7: 001000068
Input Sudoku row #8: 008500010
Input Sudoku row #9: 090000400
Solving...

Solution found in 168.791117ms:

+---+---+---+
|812|753|649|
|943|682|175|
|675|491|283|
+---+---+---+
|154|237|896|
|369|845|721|
|287|169|534|
+---+---+---+
|521|974|368|
|438|526|917|
|796|318|452|
+---+---+---+
```

## More info
- https://en.wikipedia.org/wiki/Backtracking
