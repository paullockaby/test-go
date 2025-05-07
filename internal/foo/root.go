package foo

import (
	"fmt"

	"github.com/paullockaby/test-go/internal/logging"
	"github.com/spf13/viper"
)

type Config struct {
	Verbose bool
	Name    string
	Options *viper.Viper
}

var (
	logger = logging.Log
)

func Run(config Config) error {
	logger.Info(fmt.Sprintf("gracefully exiting from %s", config.Name))
	return nil
}
