package main

import (
	"github.com/milligan22963/site/cmd/subcmd"
	"github.com/spf13/cobra"
)

var cli struct {
	Port uint16 `help:"Default serving port."`
}

func main() {
	var rootCmd = &cobra.Command{Use: "Set port to listen on"}
	rootCmd.AddCommand(subcmd.ServerCmd)
	rootCmd.Execute()
}
