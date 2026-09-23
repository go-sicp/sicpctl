package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var powerCmd = &cobra.Command{
	Use:   "power [on|off|toggle]",
	Short: "Get or set the power state",
	Long: `Without an argument, prints the current power state ("on" or "off").
With on / off / toggle, sets it.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		defer c.Close()

		if len(args) == 0 {
			on, err := c.Power()
			if err != nil {
				return err
			}
			fmt.Println(map[bool]string{true: "on", false: "off"}[on])
			return nil
		}
		switch strings.ToLower(args[0]) {
		case "on":
			return c.PowerOn()
		case "off":
			return c.PowerOff()
		case "toggle":
			on, err := c.Power()
			if err != nil {
				return err
			}
			if on {
				return c.PowerOff()
			}
			return c.PowerOn()
		default:
			return fmt.Errorf("bad arg %q (want on|off|toggle)", args[0])
		}
	},
}

func init() { rootCmd.AddCommand(powerCmd) }
