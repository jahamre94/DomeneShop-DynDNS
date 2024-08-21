package config

import (
	"flag"
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Token   string
	Secret  string
	Domains []string
}

func Load() *Config {
	once.Do(func() {
		// Define command-line flags
		tokenFlag := flag.String("token", "", "Token for authentication")
		secretFlag := flag.String("secret", "", "Secret for authentication")
		domainsFlag := flag.String("domains", "", "Comma-separated list of domains")

		// Parse flags
		flag.Parse()

		// Check if all required flags are provided
		if *tokenFlag == "" || *secretFlag == "" || *domainsFlag == "" {
			// Set up Viper for config file handling
			viper.AddConfigPath("config")
			viper.AddConfigPath(".")
			viper.AddConfigPath("$HOME/.config/dyndns")
			viper.AddConfigPath("$HOME")
			viper.SetConfigName("app")
			viper.SetConfigType("yaml")

			if err := viper.ReadInConfig(); err != nil {
				fmt.Printf("Warning: %s - Using default configuration values.\n", err)
			}
		}

		// Overwrite config file values with flags if provided
		if *tokenFlag != "" {
			viper.Set("token", *tokenFlag)
		}
		if *secretFlag != "" {
			viper.Set("secret", *secretFlag)
		}
		if *domainsFlag != "" {
			viper.Set("domains", *domainsFlag)
		}

		instance = &Config{
			Token:   viper.GetString("token"),
			Secret:  viper.GetString("secret"),
			Domains: viper.GetStringSlice("domains"),
		}
	})
	return instance
}
