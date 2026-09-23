package cmd

import (
	"fmt"

	"github.com/go-sicp/sicp"
	"github.com/spf13/cobra"
)

var videoCmd = &cobra.Command{
	Use:   "video [BRIGHTNESS COLOR CONTRAST SHARPNESS TINT BLACK GAMMA]",
	Short: "Get or set the seven video parameters",
	Long: `Without arguments, prints the current video parameters.

With seven values (decimal or 0xHH), sets them. Brightness/Color/Contrast/
Sharpness/Tint/BlackLevel are 0..100 (Tint may be -50..+50 sent as a signed
byte on Phoenix 2.0). Gamma: 0x01 native, 0x02 S, 0x03 2.2, 0x04 2.4,
0x05 D-image.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 && len(args) != 7 {
			return fmt.Errorf("expected 0 or 7 arguments, got %d", len(args))
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		defer c.Close()

		if len(args) == 0 {
			p, err := c.Video()
			if err != nil {
				return err
			}
			fmt.Printf("brightness=%d color=%d contrast=%d sharpness=%d tint=%d black=%d gamma=0x%02X\n",
				p.Brightness(), p.Color(), p.Contrast(), p.Sharpness(),
				p.Tint(), p.BlackLevel(), p.Gamma())
			return nil
		}
		vs := make([]byte, 7)
		for i, a := range args {
			v, err := parseByte(a)
			if err != nil {
				return fmt.Errorf("arg %d: %w", i+1, err)
			}
			vs[i] = v
		}
		return c.SetVideo(sicp.NewVideoParams(vs[0], vs[1], vs[2], vs[3], vs[4], vs[5], vs[6]))
	},
}

func init() { rootCmd.AddCommand(videoCmd) }
