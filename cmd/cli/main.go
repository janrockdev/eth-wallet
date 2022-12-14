package main

import (
	"fmt"
	"github.com/janrockdev/eth-wallet/cmd"
	"github.com/urfave/cli/v2"
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
