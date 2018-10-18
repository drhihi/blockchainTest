package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()
	sA := flag.Args()

	var iA []int64
	if err := getArrayInt64(&iA, &sA); err != nil {
		panic(err)
	}
	for _, v := range iA {
		fmt.Println(v)
	}

	fmt.Printf("%v => %d\n", iA, GetWaterAmount(&iA))
}

func getArrayInt64(iA *[]int64, sA *[]string) error {

	for _, v := range *sA {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*iA = append(*iA, int64(i))
	}

	return nil
}

func GetWaterAmount(iA *[]int64) (result int64) {

	var sum, maxH int64

	for _, v := range *iA {
		if v >= maxH {
			result += sum
			sum = 0
			maxH = v
		} else {
			sum += maxH - v
		}
	}

	sum, maxH = 0, 0

	for i := len(*iA) - 1; i >= 0; i-- {
		v := (*iA)[i]
		if v > maxH {
			result += sum
			sum = 0
			maxH = v
		} else {
			sum += maxH - v
		}
	}

	return
}
