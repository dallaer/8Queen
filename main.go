package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

// This function creates a matrix of int specified length and return it.
func createMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}

// This function writes matrices with solutions to a file.
func printInFile(m [][]int) {
	s, err := os.OpenFile("file.txt", os.O_APPEND, 0666) // open file in append mode
	if err != nil {
		fmt.Println("Unexpected error while open file")
	}
	for i := 0; i < len(m); i++ {
		data, err1 := json.Marshal(m[i]) // use json.Marshal to convert []int to a []byte
		if err1 != nil {
			fmt.Println("Unexpected error")
		}
		s.Write(data)
		s.WriteString("\n")
	}
	s.WriteString("\n")

}

// This function, receiving the coordinates of the place of the queens in array format,
//creates a matrix, where 1 is the queen, and 0 is an empty place
func printM(pos [][]int, s []int) {
	matrix := createMatrix(len(pos) + 1) // create a matrix with 0
	for i := 0; i < len(pos); i++ {      // iterate over the matrix and set 1 according to the previously obtained coordinates
		matrix[pos[i][0]-1][pos[i][1]-1] = 1
	}
	matrix[s[0]-1][s[1]-1] = 1
	printInFile(matrix) // run a function that writes the matrix to a file
}

// This function is backtrack algorithm. Where n is a number of queens and board size. x is a column number.
// pos is a a two-dimensional array with the coordinates of the queens. And this function is return number of solutions(combs)
func get_que(n, x, combs int, pos [][]int) int {
	for y := 1; y <= n; y++ { //iterate over a column x
		can_put := true
		for i := range pos {
			X, Y := pos[i][0], pos[i][1]
			if X == x || Y == y || math.Abs(float64(X-x)) == math.Abs(float64(Y-y)) { // this is a condition under which the queen are not under attack by other queens
				can_put = false
				break
			}
		}
		if can_put {
			if x == n { // if we can put the queen and this is the last column
				printM(pos, []int{x, y})
				return (combs + 1) // add to the number of solutions 1
			} else { // if we can put the queen and this is not the last column
				pos_copy := pos
				pos_copy = append(pos_copy, []int{x, y}) // add queen coordinates to array
				combs = get_que(n, x+1, combs, pos_copy) // recursively run the algorithm, but for the next column
			}
		}
	}
	return combs
}

// This function designed to simply print every 5 seconds, since for values greater than 10, the time to find all solutions is significant
func W8() {
	for true {
		fmt.Println("Waiting...")
		time.Sleep(5 * time.Second)
	}
}

var n string

func main() {
	for true {
		fmt.Println("Enter number")
		_, err := os.Create("file.txt") // create a file in which matrices with solutions will be written
		if err != nil {
			fmt.Println("Unexpected error while creating file")
		}
		fmt.Scan(&n)
		if n == "quit" {
			break
		}
		x, _ := strconv.Atoi(n) //convert received string to int
		if x > 0 {
			go W8()
			fmt.Println("For ", x, " Queen ", get_que(x, 1, 0, [][]int{}), " difference combination\nYou can check it in the file") // print the number of solutions and run the algorithm
			break
		} else {
			fmt.Println("Incorrect input. Try again. Enter quit for break.")
		}
	}
}
