// Package cmd holds the cobra command tree for sicpctl. The root command and
// the persistent flags live here; each subcommand is in its own file.
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sicp/sicp"
	"github.com/spf13/cobra"
)

var (
	flagHost    string
	flagMonitor uint
	flagGroup   uint
	flagTimeout time.Duration
)

var rootCmd = &cobra.Command{
	Use:   "sicpctl",
	Short: "Command-line client for the Philips SICP digital-signage protocol",
	Long: `sicpctl speaks SICP over TCP (default port 5000) to Philips professional
and digital-signage displays. See "The SICP Commands Document V2.05" for
protocol details.`,
	SilenceUsage: true, // don't print full usage on a RunE error
}

// Execute runs the root cobra command. Called from main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	pf := rootCmd.PersistentFlags()
	pf.StringVar(&flagHost, "host", "", "display host (host or host:port; port defaults to 5000)")
	pf.UintVar(&flagMonitor, "monitor", 1, "monitor ID (1..255 unicast, 0 broadcast)")
	pf.UintVar(&flagGroup, "group", 0, "group ID (0..254, 0 = address by monitor ID)")
	pf.DurationVar(&flagTimeout, "timeout", 3*time.Second, "TCP I/O timeout")
	_ = rootCmd.MarkPersistentFlagRequired("host")
}

// newClient validates the persistent flags and dials the display. The caller
// must Close() the returned client.
func newClient() (sicp.Client, error) {
	if flagMonitor > 255 {
		return nil, fmt.Errorf("--monitor must be 0..255")
	}
	if flagGroup > 254 {
		return nil, fmt.Errorf("--group must be 0..254")
	}
	t, err := sicp.DialTCP(flagHost, flagTimeout)
	if err != nil {
		return nil, err
	}
	return sicp.NewClient(t, byte(flagMonitor), byte(flagGroup)), nil
}
