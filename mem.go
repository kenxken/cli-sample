package main

import (
	"encoding/binary"

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
