package main

import (
	"os"
	"sort"

	"github.com/gladmo/film-info/cmd"
	"github.com/urfave/cli"
)

const APP_VER = "1.0.0"

func main() {
	app := cli.NewApp()

	app.Name = "Film info spider"
	app.Usage = "Film info spider"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.Douban,
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}
