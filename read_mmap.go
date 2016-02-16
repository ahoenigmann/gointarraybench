package main

import (
	"fmt"
	"os"
	"reflect"
	"syscall"
	"unsafe"

	"github.com/davecheney/profile"
)

func run() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	mmaped, err := syscall.Mmap(int(file.Fd()), 0, int(stat.Size()), syscall.PROT_READ, syscall.MAP_SHARED|syscall.MAP_POPULATE)
	syscall.Madvise(mmaped, syscall.MADV_SEQUENTIAL)
	if err != nil {
		panic(err)
	}
	sum := int64(0)
	total := 0

	newSlice := *(*reflect.SliceHeader)(unsafe.Pointer(&mmaped))
	newSlice.Len /= 8
	newSlice.Cap = newSlice.Len

	intSlice := *(*[]int64)(unsafe.Pointer(&newSlice))
	for _, val := range intSlice {
		sum += val
		total += 1
	}

	fmt.Printf("out: sum=%d count=%d\n", sum, total)

	syscall.Munmap(mmaped)
	file.Close()
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()

	for i := 0; i < 10; i++ {
		run()
	}
}
