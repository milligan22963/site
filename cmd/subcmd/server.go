package subcmd

import (
	"github.com/spf13/cobra"
)

// ServerCmd is the main server command
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Server hosts the web site",
	Long:  `A base server interface running on a device`,
	Run: func(cmd *cobra.Command, args []string) {
		serverInstance := ServerInstance{}
	},
}
