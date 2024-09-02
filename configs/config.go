package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

var Config struct {
	App struct {
		ENV                     string
		PrefixEndpoint          string
		Address                 string
		Port                    string
		TimeZone                string
		APIKey                  string
		CSRFTokenExpireDuration int
	}

	DB struct {
		Driver       string
		Host         string
		Port         string
		Name         string
		User         string
		Password     string
		Locale       string
		MaxOpenConns int
		Schema       string
	}

	JWT struct {
		SecretAccess  string
		SecretRefresh string
		TTLAccess     time.Duration
		TTLRefresh    time.Duration
	}
	Adsslot []struct {
		Position int
		Html     string
	}

	RSS struct {
		LinkRSS string
	}

	Redis struct {
		Addr     string
		Password string
		Db       int
	}
}

func init() {
	configFileLocation := getOSEnv("CONFIG_FILE_LOCATION", "LOCAL")
	viper.SetConfigType("yaml")

	// Get config file for local Development
	if configFileLocation == "LOCAL" {
		getConfigFromFile()
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		fmt.Println("Unmarshalling failed!")
	}
}

func getConfigFromFile() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

// For getting .env using default Value
// Will return all string
func getOSEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
