package main

import (
	"os"

	"github.com/wgui1/lmsensors/cmd"

	_ "github.com/wgui1/lmsensors/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
