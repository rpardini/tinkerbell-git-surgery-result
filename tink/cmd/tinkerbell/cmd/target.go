package cmd

import (
	"fmt"

	"github.com/packethost/tinkerbell/cmd/tinkerbell/cmd/target"
	"github.com/spf13/cobra"
)

// templateCmd represents the template sub-command
var targetCmd = &cobra.Command{
	Use:     "target",
	Short:   "tinkerbell target client",
	Example: "tinkerbell target [command]",
	Args: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("%v requires arguments", c.UseLine())
		}
		return nil
	},
}

func init() {
	targetCmd.AddCommand(target.SubCommands...)
	rootCmd.AddCommand(targetCmd)
}
