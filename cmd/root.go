package cmd

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/paullockaby/test-go/internal/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	logger  = logging.Log
	verbose = false

	enableMetrics bool
	metricsHost   string
	metricsPort   int

	configFile string
	options    *viper.Viper
	rootCmd    = &cobra.Command{
		Use:   "testrepo",
		Short: "A test program in the test-go repository",
	}
)

func Execute(args []string) error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $PWD/testrepo.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "increase verbosity of logging")
}

func initConfig() {
	options = viper.New()

	if configFile != "" {
		options.SetConfigFile(configFile)
	} else {
		exec, err := os.Executable()
		cobra.CheckErr(err)
		path := filepath.Dir(exec)

		// search for config in the current directory with the given name
		options.AddConfigPath(path)
		options.AddConfigPath(".")
		options.SetConfigName("testrepo.yaml")

		// supported types: "json", "toml", "yaml", "properties"
		options.SetConfigType("yaml")
	}

	options.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	options.SetEnvPrefix("TEST")
	options.AutomaticEnv()

	if err := options.ReadInConfig(); err == nil {
		logger.Info(fmt.Sprintf("using configuration file %s", options.ConfigFileUsed()))
	} else {
		logger.Warn(fmt.Sprintf("unable to load configuration file %s: %s", options.ConfigFileUsed(), err))
	}
}

func startMetricsServer(enableMetrics bool, metricsHost string, metricsPort int) {
	if !enableMetrics {
		logger.Info("metrics server is disabled")
		return
	}

	logger.Info(fmt.Sprintf("starting metrics server on %s:%d", metricsHost, metricsPort))
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", metricsHost, metricsPort), nil); err != nil {
		logger.Error(fmt.Sprintf("unable to start metrics server: %s", err))
		os.Exit(1)
	}
}
