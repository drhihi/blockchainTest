package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Int("dh", 0, "an int")
	flag.Int("db", 0, "an int")
	flag.String("hb", "", "a hex")
	flag.Int("bh", 0, "a bin")
	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "dh":
			s := f.Value.String()
			fmt.Printf("dec to hex: %s => %x\n", s, DecToHex(s))
		case "db":
			s := f.Value.String()
			fmt.Printf("dec to bin: %s => %b\n", s, DecToBin(s))
		case "hb":
			s := f.Value.String()
			fmt.Printf("hex to bin: %s => %b\n", s, HexToBin(s))
		case "bh":
			s := f.Value.String()
			fmt.Printf("bin to hex: %s => %x\n", s, BinToHex(s))
		}
	})
}

func DecToHex(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 16)
	return i
}

func DecToBin(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func HexToBin(s string) int64 {
	i, _ := strconv.ParseInt(s, 16, 64)
	return i
}

func BinToHex(s string) int64 {
	i, _ := strconv.ParseInt(s, 2, 16)
	return i
}
