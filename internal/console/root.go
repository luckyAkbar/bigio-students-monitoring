package console

import (
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/luckyAkbar/bigio-students-monitoring/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command{
	Use: "BIGIO Students Monitoring System",
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		logrus.Error(err);
		os.Exit(1)
	}
}

func setupLogger() {
	formatter := runtime.Formatter{
		ChildFormatter: &logrus.JSONFormatter{},
		Line:           true,
		Package:        true,
		File:           true,
	}

	if config.Env() == "development" {
		formatter = runtime.Formatter{
			ChildFormatter: &logrus.TextFormatter{
				ForceColors:   true,
				FullTimestamp: true,
			},
			Line:    true,
			Package: true,
			File:    true,
		}
	}

	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)

	logLevel, err := logrus.ParseLevel(config.LogLevel())
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	logrus.SetLevel(logLevel)
}