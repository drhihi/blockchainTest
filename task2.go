package main

import (
	"flag"
	"fmt"
	"strconv"
)

const (
	val3 string = "zip"
	val5 string = "zap"
)

func main() {
	flag.Parse()
	s := flag.Arg(0)
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(ZipZap(i))
}

func ZipZap(i uint64) string {
	result := ""
	for ii := uint64(1); ii <= i; ii++ {
		if ii%3 == 0 {
			result += val3
		} else if ii%5 == 0 {
			result += val5
		} else {
			result += strconv.FormatUint(ii, 10)
		}
	}
	return result
}
