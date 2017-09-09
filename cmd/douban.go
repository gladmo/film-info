package cmd

import (
	"github.com/gladmo/film-info/app/douban"
	"github.com/urfave/cli"
)

var Douban = cli.Command{
	Name:        "douban",
	Aliases:     []string{"d"},
	Usage:       "Film scrapy Command",
	Description: "Film scrapy frame",
	Action:      runScrapy,
}

func runScrapy(ctx *cli.Context) error {

	douban.Run()

	return nil
}
