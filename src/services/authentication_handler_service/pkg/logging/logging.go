package logging

import (
	"bufio"
	"bytes"
	"flag"

	core_logging "github.com/BlackspaceInc/BlackspacePlatform/src/libraries/core/core-logging/json"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"k8s.io/klog/v2"
)

func ConfigureStructuredLogging() {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	logger := core_logging.NewJSONLogger(zapcore.AddSync(writer))

	klog.SetLogger(logger)
	klog.InitFlags(nil)
	flag.Set("logtostderr", viper.GetString("ENABLE_LOG_TO_STDERR"))
	flag.Set("alsologtostderr", viper.GetString("ENABLE_LOG_TO_STDERR_AND_FILES"))
	flag.Set("log_dir", viper.GetString("LOG_DIR"))
	flag.Set("log_file", viper.GetString("LOG_FILE"))
	flag.Set("v", string(viper.GetInt("LOG_LEVEL_VERBOSITY")))
	flag.Parse()
}
