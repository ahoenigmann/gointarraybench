package main

import (
	"encoding/binary"
	"os"
	"strconv"
)

func main() {
	itemCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	for i := 0; i < itemCount; i++ {
		binary.Write(os.Stdout, binary.LittleEndian, int64(i))
	}
}
