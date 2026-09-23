package cmd

import (
	"fmt"

	"github.com/go-sicp/sicp"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print model number, firmware, build date and platform",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := newClient()
		if err != nil {
			return err
		}
		defer c.Close()

		queries := []struct {
			label string
			fn    func() (string, error)
		}{
			{"model", func() (string, error) { return c.GetModelInfo(sicp.ModelInfoNumber) }},
			{"firmware", func() (string, error) { return c.GetModelInfo(sicp.ModelInfoFW) }},
			{"build_date", func() (string, error) { return c.GetModelInfo(sicp.ModelInfoBuildDate) }},
			{"platform", func() (string, error) { return c.GetPlatform(sicp.PlatformInfoLabel) }},
			{"platform_version", func() (string, error) { return c.GetPlatform(sicp.PlatformInfoVersion) }},
			{"sicp_version", func() (string, error) { return c.GetPlatform(sicp.PlatformInfoSICPVersion) }},
		}
		for _, q := range queries {
			v, err := q.fn()
			if err != nil {
				fmt.Printf("%s: ERROR %v\n", q.label, err)
				continue
			}
			fmt.Printf("%s: %s\n", q.label, v)
		}
		return nil
	},
}

func init() { rootCmd.AddCommand(infoCmd) }
