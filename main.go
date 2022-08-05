package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func cmd(s string) (string, error) {
	//"go build -o run run_api_grace.go"
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()

	r := out.String()
	if len(r) > 0 {
		log.Println(r)
	}
	return r, err
}

func new(name string) {
	cmd(fmt.Sprintf("mkdir %s", name))
	cmd(fmt.Sprintf("curl -o %s/temp https://fab-jar.oss-cn-zhangjiakou.aliyuncs.com/t/templete.zip", name))
	cmd(fmt.Sprintf("unzip -d %s %s/temp && rm -rf %s/temp", name, name, name))

}
func main() {
	s := "wewe"
	app := &cli.App{
		Action: func(context *cli.Context) error {
			log.Println(s)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:     "new",
				Aliases:  []string{"n"},
				Usage:    "new project",
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
					fmt.Println(c.Args().Len())
					if c.Args().Len() > 0 {
						new(c.Args().Get(0))
					}
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
