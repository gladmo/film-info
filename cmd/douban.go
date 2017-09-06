package cmd

import (
	"github.com/gladmo/film-info/app/douban"
	"github.com/urfave/cli"
)

var Douban = cli.Command{
	Name:        "douban",
	Aliases:     []string{"d"},
	Usage:       "Film info Command",
	Description: "Film info frame",
	Action:      runSpider,
}

func runSpider(ctx *cli.Context) error {

	douban.Spider()

	return nil
}
