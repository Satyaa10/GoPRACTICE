package main

import (
	"fmt"
)

var b = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}

func show() {
	for _, r := range b {
		fmt.Println(r)
	}
}

func win(p string) bool {
	for i := 0; i < 3; i++ {
		if b[i][0] == p && b[i][1] == p && b[i][2] == p {
			return true
		}
		if b[0][i] == p && b[1][i] == p && b[2][i] == p {
			return true
		}
	}
	if b[0][0] == p && b[1][1] == p && b[2][2] == p {
		return true
	}
	if b[0][2] == p && b[1][1] == p && b[2][0] == p {
		return true
	}
	return false
}

func main() {
	p := "X"
	for {
		show()
		var r, c int
		fmt.Printf("Player %s (row col): ", p)
		fmt.Scan(&r, &c)

		if b[r][c] != " " {
			fmt.Println("Spot taken!")
			continue
		}

		b[r][c] = p

		if win(p) {
			show()
			fmt.Println("Winner:", p)
			break
		}

		if p == "X" {
			p = "O"
		} else {
			p = "X"
		}
	}
}
