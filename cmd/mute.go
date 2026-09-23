package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var muteCmd = &cobra.Command{
	Use:   "mute [on|off|toggle]",
	Short: "Get or set audio mute (V2.00+)",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		defer c.Close()

		if len(args) == 0 {
			on, err := c.Mute()
			if err != nil {
				return err
			}
			fmt.Println(map[bool]string{true: "on", false: "off"}[on])
			return nil
		}
		switch strings.ToLower(args[0]) {
		case "on":
			return c.SetMute(true)
		case "off":
			return c.SetMute(false)
		case "toggle":
			on, err := c.Mute()
			if err != nil {
				return err
			}
			return c.SetMute(!on)
		default:
			return fmt.Errorf("bad arg %q (want on|off|toggle)", args[0])
		}
	},
}

func init() { rootCmd.AddCommand(muteCmd) }
