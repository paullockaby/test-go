package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"log/slog"

	"github.com/paullockaby/test-go/internal/bar"
	"github.com/paullockaby/test-go/internal/logging"
)

var (
	hideHealthAccessLogs bool
	hideNormalAccessLogs bool
	barHost              string
	barPort              int
	barCmd               = &cobra.Command{
		Use:   "bar",
		Short: "Listens for messages and sends the results back",
		RunE: func(cmd *cobra.Command, args []string) error {
			// NOTE:
			//    args is just an array of things that come after a "--" in the command line
			//    for example: testrepo bar -- these are args
			if verbose {
				logging.SetLevel(slog.LevelDebug)
				logger.Debug("verbose logging enabled")
			}

			// kick off the metrics server
			if metricsHost == "" {
				return errors.New("invalid metrics host")
			}
			if metricsPort < 0 || metricsPort > 65535 {
				return errors.New("invalid metrics port")
			}

			go startMetricsServer(enableMetrics, metricsHost, metricsPort)

			// validate this application's configuration values
			if barHost == "" {
				return errors.New("invalid listener host")
			}
			if barPort < 0 || barPort > 65535 {
				return errors.New("invalid listener port")
			}

			config := bar.Config{
				Verbose:              verbose,
				ListenerHost:         barHost,
				ListenerPort:         barPort,
				HideHealthAccessLogs: hideHealthAccessLogs,
				HideNormalAccessLogs: hideNormalAccessLogs,
				Options:              options,
			}

			if err := bar.Run(config); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(barCmd)

	barCmd.Flags().StringVar(&barHost, "listener-host", "localhost", "host on which to listen")
	barCmd.Flags().IntVar(&barPort, "listener-port", 8086, "port on which to listen")
	barCmd.Flags().BoolVar(&hideHealthAccessLogs, "hide-health-access-logs", false, "do not log normal requests to the health endpoint")
	barCmd.Flags().BoolVar(&hideNormalAccessLogs, "hide-normal-access-logs", false, "do not log normal requests to any endpoint")

	barCmd.Flags().BoolVar(&enableMetrics, "enable-metrics", true, "should the metrics server be enabled")
	barCmd.Flags().StringVar(&metricsHost, "metrics-host", "localhost", "interface on which to listen for metrics")
	barCmd.Flags().IntVar(&metricsPort, "metrics-port", 10000, "port on which to listen for metrics")
}
