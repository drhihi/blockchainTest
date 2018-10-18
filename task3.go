package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

func main()  {
	var r, x, y float64
	var err error
	flag.Parse()
	s1 := flag.Arg(0)
	s2 := flag.Arg(1)
	s3 := flag.Arg(2)
	if r, err = strconv.ParseFloat(s1, 64); err != nil {
		panic(err)
	}
	if x, err = strconv.ParseFloat(s2, 64); err != nil {
		panic(err)
	}
	if y, err = strconv.ParseFloat(s3, 64); err != nil {
		panic(err)
	}
	fmt.Printf("(%g, %g, %g) => %t\n", r, x, y, IsInCircle(r, x, y))
}

func IsInCircle(r, x, y float64) bool {
	result := math.Pow(x, 2) + math.Pow(y, 2) <= math.Pow(r, 2)
	return result
}