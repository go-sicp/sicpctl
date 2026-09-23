package cmd

import (
	"fmt"
	"strings"

	"github.com/go-sicp/sicp"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart [android|scalar]",
	Short: "Restart the monitor's Android or scalar subsystem (V2.02+)",
	Long: `Sends a monitor-restart command. The default target is "android".
The display ACKs and immediately reboots the targeted subsystem; further
commands will fail until it comes back online.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		target := sicp.RestartAndroid
		if len(args) == 1 {
			switch strings.ToLower(args[0]) {
			case "android":
				target = sicp.RestartAndroid
			case "scalar", "scaler":
				target = sicp.RestartScalar
			default:
				return fmt.Errorf("bad target %q (want android|scalar)", args[0])
			}
		}
		c, err := newClient()
		if err != nil {
			return err
		}
		defer c.Close()
		return c.Restart(target)
	},
}

func init() { rootCmd.AddCommand(restartCmd) }
