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
	fmt.Println("Specify board dimensions and the number of black holes:")
	w, err := ReadUserInputInt("Board Width:")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	h, err := ReadUserInputInt("Board Height:")
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
		drawBoard(g.GetBoard(), false)
		fmt.Println("Make next turn")

		x, err := ReadUserInputInt("Enter cell x coordinate: ")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		y, err := ReadUserInputInt("Enter cell y coordinate: ")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		state, err := g.MakeTurn(x, y)
		if err != nil {
			fmt.Println("Error! ", err)
			break
		}

		if state == game.Lost {
			fmt.Println("You've lost!")
			break
		} else if state == game.Won {
			fmt.Println("You've won")
			break
		}
	}

	drawBoard(g.GetBoard(), true)
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

func drawBoard(board [][]game.Cell, showBombs bool) {
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
