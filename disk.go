package main

import "syscall"

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
