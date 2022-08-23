package main

import (
	"fmt"
	"github.com/janrockdev/eth-wallet/cmd"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			cmd.ServerCommand,
			cmd.NodeCommand,
			cmd.NodewalletCommand,
			cmd.WalletCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Errorf("%s", err)
	}
}
