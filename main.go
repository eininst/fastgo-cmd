package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	s := "wewe"
	app := &cli.App{
		Action: func(context *cli.Context) error {
			log.Println(s)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:     "add",
				Aliases:  []string{"a"},
				Usage:    "calc 1+1",
				Category: "arithmetic",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "port, p",
						Value: 8000,
						Usage: "listening port",
					},
					&cli.StringFlag{
						Name:        "lang, l",
						Value:       "english",
						Usage:       "read from `FILE`",
						Destination: &s,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("1 + 1 = ", 1+1)

					return nil
				},
			},
		},
	}
	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:  "port, p",
			Value: 8000,
			Usage: "listening port",
		},
		&cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "read from `FILE`",
			Destination: &s,
		},
	}
	//app.Commands = []*cli.Command{
	//	{
	//		Name:     "add",
	//		Aliases:  []string{"a"},
	//		Usage:    "calc 1+1",
	//		Category: "arithmetic",
	//		Action: func(c *cli.Context) error {
	//			fmt.Println("1 + 1 = ", 1+1)
	//
	//			return nil
	//		},
	//	},
	//}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
