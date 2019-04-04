package main

import (
	"fmt"
	"syscall"
)

const (
	b  = 1
	kB = 1024 * b
	mB = 1024 * kB
	gB = 1024 * mB
)

type DiskStatus struct {
	All       uint64 `json:"all"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
	Free      uint64 `json:"free"`
}

func DiskUsage(path string) (disk DiskStatus) {
	stat := syscall.Statfs_t{}
	err := syscall.Statfs(path, &stat)
	if err != nil {
		panic(err)
	}

	disk.All = stat.Blocks * uint64(stat.Bsize)
	disk.Free = stat.Bfree * uint64(stat.Bsize)
	disk.Available = stat.Bavail * uint64(stat.Bsize)
	disk.Used = disk.All - disk.Free

	return
}

func DiskShow(disk DiskStatus) {
	fmt.Printf("All  : %.2f GB\n", float64(disk.All)/float64(gB))
	fmt.Printf("Used : %.2f GB\n", float64(disk.Used)/float64(gB))
	fmt.Printf("Free : %.2f GB\n", float64(disk.Free)/float64(gB))
	fmt.Printf("Avail: %.2f GB\n", float64(disk.Available)/float64(gB))
}
