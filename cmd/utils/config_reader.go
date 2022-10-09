package utils

import (
	"flag"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/uberate/mockers/pkg/files"
	"os"
	"strings"
)

// This file provide some function to help read the config. The level is ENV > ConfigFile. If is the command program,
// the level should be ENV > ConfigFile > Command. But, in most time, the program will control by user, so, the program
// cloud ignore the config.

const (
	ConfigPathDirDefaultEnvKey     = "CONFIG_DIR_PATH"
	ConfigPathFileDefaultFlag      = "config"
	ConfigPathFileDefaultShortFlag = "c"
	ConfigDefaultPath              = "~/config/config.yaml"
)

// ReadConfig will load the config from the env or file, it can receive the default value. And the first read env value.
// About the param: 'configPathEnvKey', 'configShortKey', 'configFlagKey' and 'enableFlag' see the function:
// Get GetConfigFilePath.
//
// The default value is from input param: 'config'. The config must a point. All the config can read from env
// properties. And the level is ENV > config-file > default. Because, for simple change any config in docker or k8s env.
//
// If you want to custom the flags, you can set enableFlag is false, the ReadConfig will ignore all command flag.
//
// About the configPathEnvKey, configShortKey, configFlagKey and enableFlag, see the function: GetConfigFilePath.
func ReadConfig(
	configPathEnvKey, configShorKey, configFlagKey string, enableFlag bool, // config file settings
	envKey string, config interface{},                                      // config read setting
) error {

	viperInstance := viper.NewWithOptions(
		viper.KeyDelimiter("."),
	)

	// set config key to help search env value.
	originConfigMap := map[string]interface{}{}
	if err := mapstructure.Decode(config, &originConfigMap); err != nil {
		return err
	}
	if err := viperInstance.MergeConfigMap(originConfigMap); err != nil {
		return err
	}

	// try to search the file of config, but if not found the file, skip
	filePath, ok := GetConfigFilePath(configPathEnvKey, configShorKey, configFlagKey, enableFlag)
	if ok {
		filePathSpilt := strings.Split(filePath, ".")
		viperInstance.SetConfigType(filePathSpilt[len(filePathSpilt)-1])
		fileReader, err := os.Open(filePath)
		if err != nil {
			// if config file read error, stop directly.
			return err
		}
		if err := viperInstance.MergeConfig(fileReader); err != nil {
			// if config file read error, stop directly.
			return err
		}
		if err := fileReader.Close(); err != nil {
			// if config file read error, stop directly.
			return err
		}

	}

	// set the env value
	viperInstance.SetEnvPrefix(envKey)
	viperInstance.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viperInstance.AutomaticEnv()

	// try to unmarshal to the config interface
	return viperInstance.Unmarshal(config)
}

// GetConfigFilePath will search and return a readable file. The search key is in the env and flag.
// If the input-param configEnvKey is emtpy, will set to ConfigPathDirDefaultEnvKey.
// If the input-param configFlagKey is empty, will set to ConfigPathFileDefaultFlag.
// If the input-param configFlagKey is empty, will set to ConfigPathFileDefaultShortFlag.
//
// If the enableFlag is empty, ignore the configFlagShortKey and configFlagKey. It can help the command program to limit
// the flags behavior and type.
//
// The search order is : env -> flag-short-key -> flag key
//
// If the function search all default value is empty, will try to search the file: ~/config/config.yaml
//
// The GetConfigFilePath can search some file type like yaml, json, toml and so on. If all file has no value, will
// return emtpy string and false. Else return the readable file path and true.
func GetConfigFilePath(configEnvKey, configFlagShortKey, configFlagKey string, enableFlag bool) (string, bool) {

	//--------------------------------------------------

	// set default values
	if len(configEnvKey) == 0 {
		configEnvKey = ConfigPathDirDefaultEnvKey
	}

	if len(configFlagShortKey) == 0 {
		configFlagShortKey = ConfigPathFileDefaultShortFlag
	}

	if len(configFlagKey) == 0 {
		configFlagKey = ConfigPathFileDefaultFlag
	}

	//--------------------------------------------------
	// try to search the config file.

	// search the env
	path := os.Getenv(configEnvKey)
	if files.IsFileExists(path) {
		return path, true
	}

	// try to search the flag.
	if enableFlag {
		var shortFlagKey, flagKey string
		flag.StringVar(&shortFlagKey, configFlagShortKey, "", "to set the config file path. "+
			"Equals 'config'")
		flag.StringVar(&flagKey, configFlagKey, "", "to set the config file path, short is -c. Default "+
			"is: "+ConfigDefaultPath)

		flag.Parse()

		if files.IsFileExists(shortFlagKey) {
			return shortFlagKey, true
		}
		if files.IsFileExists(flagKey) {
			return flagKey, true
		}
	}

	// try to search the default path
	if files.IsFileExists(ConfigDefaultPath) {
		return ConfigDefaultPath, true
	}

	// not found any readable file.
	return "", false
}
