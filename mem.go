package main

import (
	"encoding/binary"
	"fmt"

	"golang.org/x/sys/unix"
)

func MemStatus() uint64 {
	totalStr, err := unix.Sysctl("hw.memsize")
	if err != nil {
		panic(err)
	}

	totalStr += "\x00"
	total := uint64(binary.LittleEndian.Uint64([]byte(totalStr)))

	return total
}

func MemShow(mem uint64) {
	fmt.Printf("Total: %.2f GB\n", float64(mem)/float64(gB))
}
