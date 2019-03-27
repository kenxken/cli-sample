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

	app.Name = "cli sample"
	app.Usage = "This app echo system status"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		if context.Bool("disk") {
			disk := DiskUsage("/")
			fmt.Printf("All  : %.2f GB\n", float64(disk.All)/float64(gB))
			fmt.Printf("Used : %.2f GB\n", float64(disk.Used)/float64(gB))
			fmt.Printf("Free : %.2f GB\n", float64(disk.Free)/float64(gB))
			fmt.Printf("Avail: %.2f GB\n", float64(disk.Available)/float64(gB))
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
