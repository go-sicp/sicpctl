package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var volumeCmd = &cobra.Command{
	Use:   "volume [N]",
	Short: "Get or set the volume (0..100)",
	Long: `Without an argument, prints the current speaker and audio-out volumes.
With a percentage 0..100, sets both speaker and audio-out to that value.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		defer c.Close()

		if len(args) == 0 {
			spk, ao, err := c.VolumePercent()
			if err != nil {
				return err
			}
			fmt.Printf("speaker=%d audio_out=%d\n", spk, ao)
			return nil
		}
		v, err := strconv.Atoi(args[0])
		if err != nil || v < 0 || v > 100 {
			return fmt.Errorf("want integer 0..100, got %q", args[0])
		}
		return c.SetVolumePercent(byte(v), byte(v))
	},
}

func init() { rootCmd.AddCommand(volumeCmd) }
