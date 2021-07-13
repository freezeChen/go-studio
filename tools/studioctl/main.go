package main

import (
	"fmt"
	"github.com/freezeChen/go-studio/tools/studioctl/model"
	"github.com/urfave/cli"
	"os"
	"runtime"
)

var (
	buildVersion = "0.0.1"
	commands     = []cli.Command{
		{
			Name:  "model",
			Usage: "generate model code",
			Subcommands: []cli.Command{
				{
					Name:  "mysql",
					Usage: "generate mysql model",
					Subcommands: []cli.Command{
						{
							Name:  "datasource",
							Usage: "generate model from datasource",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "url",
									Usage: `the data source of database, like "root:password@tcp(127.0.0.1:3306)/database"`,
								},
								cli.StringFlag{
									Name:  "table,t",
									Usage: "the table or table globbing patterns in the database",
								},
								cli.StringFlag{
									Name:  "dir,d",
									Usage: "the target dir",
								},
							},
							Action: model.MysqlDataSource,
						},
					},
				},
			},

		},
	}
)

func main() {
	app := cli.NewApp()
	app.Version = fmt.Sprintf("%s %s/%s", buildVersion, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Println("error:", err)
		return
	}
}
