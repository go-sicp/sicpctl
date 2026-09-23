package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var backlightCmd = &cobra.Command{
	Use:   "backlight [on|off|toggle]",
	Short: "Get or set the panel backlight (V2.02+)",
	Long: `Without an argument, prints the current backlight state. With on / off /
toggle, sets it. Backlight off blanks the screen while leaving the rest of
the system running (useful instead of full power-off).`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		defer c.Close()

		if len(args) == 0 {
			on, err := c.Backlight()
			if err != nil {
				return err
			}
			fmt.Println(map[bool]string{true: "on", false: "off"}[on])
			return nil
		}
		switch strings.ToLower(args[0]) {
		case "on":
			return c.SetBacklight(true)
		case "off":
			return c.SetBacklight(false)
		case "toggle":
			on, err := c.Backlight()
			if err != nil {
				return err
			}
			return c.SetBacklight(!on)
		default:
			return fmt.Errorf("bad arg %q (want on|off|toggle)", args[0])
		}
	},
}

func init() { rootCmd.AddCommand(backlightCmd) }
