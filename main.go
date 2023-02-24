package main

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == j && i == 0) || (j == x-1 && i == 0) || (i == y-1 && j == 0) || (j == x-1 && i == y-1) {
				z01.PrintRune('o')
			} else if (i == 0 && (j > 0 && j < x-1)) || (i == y-1 && (j > 0 && j < x-1)) {
				z01.PrintRune('-')
			} else if (j == 0 && (i > 0 && i < y-1)) || (j == x-1 && (i > 0 && i < y-1)) {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadB(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == j && i == 0 {
				z01.PrintRune('/')
			} else if j == x-1 && i == 0 {
				z01.PrintRune(rune(92))
			} else if (i == y-1 || i == 0) && (j > 0 && j < x-1) {
				z01.PrintRune('*')
			} else if (j == x-1 || j == 0) && (i > 0 && i < y-1) {
				z01.PrintRune('*')
			} else if j == 0 && i == y-1 {
				z01.PrintRune(rune(92))
			} else if i == y-1 && j == x-1 {
				z01.PrintRune('/')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	QuadB(1, 5)
}
