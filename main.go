package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const (
	b  = 1
	kB = 1024 * b
	mB = 1024 * kB
	gB = 1024 * mB
)

func main() {
	app := cli.NewApp()

	app.Name = "sample-cli"
	app.Usage = "This app echo system status"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		if context.Bool("disk") {
			disk := DiskUsage("/")
			fmt.Printf("All  : %.2f GB\n", float64(disk.All)/float64(gB))
			fmt.Printf("Used : %.2f GB\n", float64(disk.Used)/float64(gB))
			fmt.Printf("Free : %.2f GB\n", float64(disk.Free)/float64(gB))
			fmt.Printf("Avail: %.2f GB\n", float64(disk.Available)/float64(gB))
		} else if context.Bool("mem") {
			total := MemStatus()
			fmt.Printf("Total: %.2f GB\n", float64(total)/float64(gB))
		} else if context.Bool("cpu") {
			cpu, _ := CpuUsage()
			fmt.Printf("User  : %f %%\n", float64(cpu.User)/float64(cpu.Total)*100)
			fmt.Printf("System: %f %%\n", float64(cpu.System)/float64(cpu.Total)*100)
			fmt.Printf("Idle  : %f %%\n", float64(cpu.Idle)/float64(cpu.Total)*100)
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "disk, d",
			Usage: "Echo disk size",
		},
		cli.BoolFlag{
			Name:  "mem, m",
			Usage: "Echo memory size (darwin)",
		},
		cli.BoolFlag{
			Name:  "cpu, c",
			Usage: "Echo cpu usage (darwin)",
		},
	}

	app.Run(os.Args)
}
