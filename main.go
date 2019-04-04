package main

import (
	"os"

	"github.com/urfave/cli"
)

type statusFunc func()

func diskStatus() {
	disk := DiskUsage("/")
	DiskShow(disk)
}

func cpuStatus() {
	cpu, _ := CpuUsage()
	CpuShow(*cpu)
}

func memStatus() {
	mem := MemStatus()
	MemShow(mem)
}

func main() {
	app := cli.NewApp()

	app.Name = "sample-cli"
	app.Usage = "This app echo system status"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		var sf statusFunc
		if context.Bool("disk") {
			sf = diskStatus
		} else if context.Bool("mem") {
			sf = memStatus
		} else if context.Bool("cpu") {
			sf = cpuStatus
		} else {
			return nil
		}
		sf()
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
