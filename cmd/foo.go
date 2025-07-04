package cmd

import (
	"github.com/spf13/cobra"
	"log/slog"

	"github.com/paullockaby/test-go/internal/foo"
	"github.com/paullockaby/test-go/internal/logging"
)

var (
	fooName string
	fooCmd  = &cobra.Command{
		Use:   "foo",
		Short: "Does some foo stuff",
		RunE: func(cmd *cobra.Command, args []string) error {
			// NOTE:
			//    args is just an array of things that come after a "--" in the command line
			//    for example: testrepo foo -- these are args
			if verbose {
				logging.SetLevel(slog.LevelDebug)
				logger.Debug("verbose logging enabled")
			}

			config := foo.Config{
				Verbose: verbose,
				Name:    fooName,
				Options: options,
			}

			if err := foo.Run(config); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(fooCmd)

	fooCmd.Flags().StringVar(&fooName, "name", "", "name of foo to run")
	if err := fooCmd.MarkFlagRequired("name"); err != nil {
		logger.Error("failed to mark foo name flag as required", "error", err)
	}
}
