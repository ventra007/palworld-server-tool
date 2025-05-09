package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/zaigie/palworld-server-tool/internal/logger"
)

type Config struct {
	Web struct {
		Password  string `mapstructure:"password"`
		Port      int    `mapstructure:"port"`
		Tls       bool   `mapstructure:"tls"`
		CertPath  string `mapstructure:"cert_path"`
		KeyPath   string `mapstructure:"key_path"`
		PublicUrl string `mapstructure:"public_url"`
	} `mapstructure:"web"`
	Task struct {
		SyncInterval        int    `mapstructure:"sync_interval"`
		PlayerLogging       bool   `mapstructure:"player_logging"`
		PlayerLoginMessage  string `mapstructure:"player_login_message"`
		PlayerLogoutMessage string `mapstructure:"player_logout_message"`
	} `mapstructure:"task"`
	Rcon struct {
		Address   string `mapstructure:"address"`
		Password  string `mapstructure:"password"`
		UseBase64 bool   `mapstructure:"use_base64"`
		Timeout   int    `mapstructure:"timeout"`
	} `mapstructure:"rcon"`
	Rest struct {
		Address  string `mapstructure:"address"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Timeout  int    `mapstructure:"timeout"`
	} `mapstructure:"rest"`
	Save struct {
		Path           string `mapstructure:"path"`
		DecodePath     string `mapstructure:"decode_path"`
		SyncInterval   int    `mapstructure:"sync_interval"`
		BackupInterval int    `mapstructure:"backup_interval"`
		BackupKeepDays int    `mapstructure:"backup_keep_days"`
	} `mapstructure:"save"`
	Manage struct {
		KickNonWhitelist bool `mapstructure:"kick_non_whitelist"`
	}
}

func Init(cfgFile string, conf *Config) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn("config file not found, try to read from env\n")
		} else {
			logger.Panic("config file was found but another error was produced\n")
		}
	}

	viper.SetDefault("web.port", 8080)

	viper.SetDefault("task.sync_interval", 60)

	viper.SetDefault("rcon.timeout", 5)
	viper.SetDefault("rcon.use_base64", false)

	viper.SetDefault("rest.username", "admin")
	viper.SetDefault("rest.timeout", 5)

	viper.SetDefault("save.sync_interval", 600)
	viper.SetDefault("save.backup_interval", 14400)
	viper.SetDefault("save.backup_keep_days", 7)

	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err = viper.Unmarshal(conf)
	if err != nil {
		logger.Panicf("Unable to decode config into struct, %s", err)
	}
}
