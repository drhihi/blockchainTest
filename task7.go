package main

import (
	"flag"
	"fmt"
	"strconv"
)

const SizeBoard int = 8

func main()  {
	flag.Parse()
	f1 := flag.Arg(0)
	f2 := flag.Arg(1)
	x, err := strconv.Atoi(f1)
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(f2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", GetQueenPossiblePoints(x, y))
}

func GetQueenPossiblePoints(x, y int) [][2]int {
	var steps [][2]int
	//top; top-right
	for i := 1; i <= SizeBoard; i++ {
		//top
		if iy := y - i; iy > 0 {
			steps = append(steps, [2]int{x, iy})
		}
		//top-right
		if ix, iy := x + i, y - i; ix <= SizeBoard && iy > 0 {
			steps = append(steps, [2]int{ix, iy})
		}
		//right
		if ix := x + i; ix <= SizeBoard {
			steps = append(steps, [2]int{ix, y})
		}
		//bottom-right
		if ix, iy := x + i, y + i; ix <= SizeBoard && iy <= SizeBoard {
			steps = append(steps, [2]int{ix, iy})
		}
		//bottom
		if iy := y + i; iy <= SizeBoard {
			steps = append(steps, [2]int{x, iy})
		}
		//bottom-left
		if ix, iy := x - i, y + i; ix > 0 && iy <= SizeBoard {
			steps = append(steps, [2]int{ix, iy})
		}
		//left
		if ix := x - i; ix > 0 {
			steps = append(steps, [2]int{ix, y})
		}
		//top-left
		if ix, iy := x - i, y - i; ix > 0 && iy > 0 {
			steps = append(steps, [2]int{ix, iy})
		}
	}
	return steps
}
