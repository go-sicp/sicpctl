package cmd

import (
	"fmt"

	"github.com/go-sicp/sicp"
	"github.com/spf13/cobra"
)

var inputsCmd = &cobra.Command{
	Use:   "inputs",
	Short: "List the input sources available on the connected display (V2.05+)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		defer c.Close()

		srcs, err := c.AvailableSources()
		if err != nil {
			return err
		}
		for _, b := range srcs {
			name := sicp.SourceName(b)
			if name == "" {
				name = "?"
			}
			fmt.Printf("0x%02X  %s\n", b, name)
		}
		return nil
	},
}

func init() { rootCmd.AddCommand(inputsCmd) }
