package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dimon2027/proxx/game"
)

func main() {
	fmt.Println("Specify board dimensions and the number of black holes")
	w, err := ReadUserInputInt(fmt.Sprintf("Board Width (max value - %d):", game.MaxBoardW))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	h, err := ReadUserInputInt(fmt.Sprintf("Board Height (max value - %d):", game.MaxBoardH))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	bhCnt, err := ReadUserInputInt("Black holes number:")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	g := game.New()
	err = g.Init(w, h, bhCnt, game.NewRandomLayout())
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for g.GetState() == game.Ongoing {
		DrawBoard(g.GetBoard(), false)
		fmt.Println("Make next turn")

		x, err := ReadUserInputInt("Enter cell x coordinate (left-to-right, 0-based):")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		y, err := ReadUserInputInt("Enter cell y coordinate (top-to-bottom, 0-based):")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		_, err = g.MakeTurn(x, y)
		if err != nil {
			fmt.Println("Error! ", err)
			continue
		}
	}

	DrawBoard(g.GetBoard(), true)

	state := g.GetState()

	if state == game.Lost {
		fmt.Println("You've lost!")
	} else if state == game.Won {
		fmt.Println("You've won!")
	} else {
		fmt.Println("Error: Invalid game state")
	}
}

func ReadUserInputInt(prompt string) (int, error) {
	fmt.Print(prompt, " ")
	reader := bufio.NewReader(os.Stdin)
	reader.Reset(os.Stdin)
	vstr, err := reader.ReadString('\r')
	if err != nil {
		return 0, err
	}

	vstr = strings.TrimSuffix(vstr, "\r")
	v, err := strconv.ParseInt(vstr, 10, 32)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

func DrawBoard(board [][]game.Cell, showBombs bool) {
	fmt.Println()
	fmt.Print(" ")
	for range board {
		fmt.Print("_")
	}

	fmt.Println()

	for y := range board {
		fmt.Print("|")
		for x := range board[y] {
			if board[x][y].GetType() == game.BlackHole && showBombs {
				fmt.Print("*")
			} else {
				if board[x][y].GetState() == game.Open {
					fmt.Print(board[x][y].GetAdjBHolesCnt())
				} else {
					fmt.Print("#")
				}
			}
		}
		fmt.Print("|")
		fmt.Println()
	}

	fmt.Print(" ")
	for range board {
		fmt.Print("-")
	}

	fmt.Println()
}
