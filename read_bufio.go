package main

import (
	"bufio"
	"fmt"
	"os"
	"unsafe"

	"github.com/davecheney/profile"
)

func run() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	bufferedFile := bufio.NewReaderSize(file, 8*1024*1024)
	buf := make([]byte, 8)
	sum := int64(0)
	total := 0

	for {
		_, err := bufferedFile.Read(buf)
		if err != nil {
			break
		}
		cptr := unsafe.Pointer(&buf[0])
		converted := *(*int64)(cptr)
		sum += converted
		total += 1
	}

	fmt.Printf("out: sum=%d count=%d\n", sum, total)
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()

	for i := 0; i < 10; i++ {
		run()
	}
}
