package main

import (
	"bufio"
	"fmt"
	"os"
)

type Board [][]int

func main() {
	board := Board{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	fmt.Println(board[0][0])
	printBoard(board)
	/*
		text1, text2 := readInput()
		fmt.Print(text1)
		fmt.Print(text2)
		for !verifyInput(text1, text2) {
			fmt.Println("ERROR INPUT")
			text1, text2 = readInput()
			fmt.Print(text1)
			fmt.Print(text2)
		}
	*/
}
func readInput() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter y-coordinate: ")
	text1, _ := reader.ReadString('\n')
	fmt.Print("Enter x-coordinate: ")
	text2, _ := reader.ReadString('\n')
	return text1, text2
}
func verifyInput(text1 string, text2 string) bool {
	if checkCharInput(text1) && checkCharInput(text2) {
		return true
	}
	return false
}
func checkCharInput(text string) bool {
	if text == "0\n" || text == "1\n" || text == "2\n" {
		return true
	}
	return false
}
func printBoard(board Board) {
	/* output each array element's value */
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			fmt.Printf("[%d][%d] = %d\n", x, y, board[x][y])
		}
	}
}
