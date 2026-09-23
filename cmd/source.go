package cmd

import (
	"fmt"

	"github.com/go-sicp/sicp"
	"github.com/spf13/cobra"
)

var sourceCmd = &cobra.Command{
	Use:   "source [NAME|0xHH]",
	Short: "Get or set the input source",
	Long: `Without an argument, prints the current input source (label or 0xHH if
unknown). With a label (hdmi, dvi-d, vga, displayport1, ...) or a byte (0x0D),
sets it and shows the source label on the OSD.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		defer c.Close()

		if len(args) == 0 {
			s, err := c.Source()
			if err != nil {
				return err
			}
			name := sicp.SourceName(s)
			if name == "" {
				name = fmt.Sprintf("0x%02X", s)
			}
			fmt.Println(name)
			return nil
		}
		src, ok := sicp.SourceFromName(args[0])
		if !ok {
			v, err := parseByte(args[0])
			if err != nil {
				return err
			}
			src = v
		}
		return c.SetSource(src, true)
	},
}

func init() { rootCmd.AddCommand(sourceCmd) }
