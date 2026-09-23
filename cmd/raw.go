package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rawCmd = &cobra.Command{
	Use:   "raw BYTE [BYTE ...]",
	Short: "Send raw data bytes; print reply data",
	Long: `Send arbitrary data bytes (decimal or 0xHH). The first byte is the
SICP command (Data[0]); remaining bytes are its parameters. Reply data bytes
are printed in hex. With --monitor 0 (broadcast) no reply is read.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		defer c.Close()

		data := make([]byte, len(args))
		for i, a := range args {
			v, err := parseByte(a)
			if err != nil {
				return fmt.Errorf("arg %d: %w", i+1, err)
			}
			data[i] = v
		}
		rep, err := c.Do(data...)
		if err != nil {
			return err
		}
		if rep == nil {
			fmt.Println("(broadcast: no reply)")
			return nil
		}
		fmt.Printf("reply: % X\n", rep)
		return nil
	},
}

func init() { rootCmd.AddCommand(rawCmd) }
