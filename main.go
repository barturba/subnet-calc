package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "subnet-calc",
		Usage: "calculated subnets",
		Action: func(*cli.Context) error {
			fmt.Println("here's the subnet!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
