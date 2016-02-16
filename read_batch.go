package main

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"

	"github.com/davecheney/profile"
)

func run() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 64*1024*8)
	sum := int64(0)
	total := 0

	for {
		bytesRead, err := file.Read(buf)
		if err != nil {
			break
		}
		slicedToCapacity := buf[:bytesRead]

		newSlice := *(*reflect.SliceHeader)(unsafe.Pointer(&slicedToCapacity))
		newSlice.Len /= 8
		newSlice.Cap = newSlice.Len

		intSlice := *(*[]int64)(unsafe.Pointer(&newSlice))
		for _, val := range intSlice {
			sum += val
			total += 1
		}
	}

	fmt.Printf("out: sum=%d count=%d\n", sum, total)
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()

	for i := 0; i < 10; i++ {
		run()
	}
}
