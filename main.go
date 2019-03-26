package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/urfave/cli"
)

type diskStatus struct {
	All       uint64 `json:"all"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
	Free      uint64 `json:"free"`
}

func diskUsage(path string) (disk diskStatus) {
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

const (
	b  = 1
	kB = 1024 * b
	mB = 1024 * kB
	gB = 1024 * mB
)

func main() {
	app := cli.NewApp()

	app.Name = "cli sample"
	app.Usage = "This app echo system status"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		if context.Bool("disk") {
			disk := diskUsage("/")
			fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(gB))
			fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(gB))
			fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(gB))
			fmt.Printf("Available: %.2f GB\n", float64(disk.Available)/float64(gB))
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "disk, d",
			Usage: "Echo disk usage",
		},
	}

	app.Run(os.Args)
}
