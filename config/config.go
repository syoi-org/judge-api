package config

import (
	"bytes"
	_ "embed"
	"path"
	"strings"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var packageName = "judge-api"

//go:embed default.yaml
var defaultConfig []byte

func InitConfig(configFile string) {
	zap.S().Infow("Loading config for package", "package", packageName)
	viper.SetConfigType("yaml")
	zap.S().Infow("Loading default config")
	viper.MergeConfig(bytes.NewBuffer(defaultConfig))

	if configFile != "" {
		viper.SetConfigFile(configFile)
		zap.S().Info("Loading config", "file", configFile)
	} else {
		viper.SetConfigName("config")

		viper.AddConfigPath(path.Join("/etc", packageName))
		for _, configDir := range xdg.ConfigDirs {
			viper.AddConfigPath(path.Join(configDir, packageName))
		}
		viper.AddConfigPath(path.Join(xdg.ConfigHome, packageName))
		viper.AddConfigPath("./config")
		zap.S().Info("Searching config from default paths")
	}

	// support reading from environmental variables
	// all env variables are capitalized, dot (levels) and dashes are replaced with underscores
	viper.SetOptions(viper.ExperimentalBindStruct())
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	err := viper.MergeInConfig()

	if err != nil {
		zap.S().Warnw("Error in reading config", "err", err)
	}
}
