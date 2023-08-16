package config

import (
	"encoding/json"
	"github.com/spf13/viper"
)

func GetDatabaseConfig() DatabaseConfig {
	var (
		dbConfig struct {
			Name     string `json:"name"`
			UserName string `json:"username"`
			Password string `json:"password"`
		}
		dbHost   string
		dbName   string
		User     string
		Password string
	)
	dbHost = viper.GetString("DATABASE_HOST")
	if dbHost == "" {
		dbHost = viper.GetString("database.host")
	}
	_ = json.Unmarshal([]byte(viper.GetString("DATABASE_CONFIG")), &dbConfig)
	if dbName = dbConfig.Name; dbName == "" {
		dbName = viper.GetString("database.name")
	}
	if User = dbConfig.UserName; User == "" {
		User = viper.GetString("database.user")
	}
	if Password = dbConfig.Password; Password == "" {
		Password = viper.GetString("database.password")
	}
	return DatabaseConfig{
		DbHost:   dbHost,
		DbName:   dbName,
		User:     User,
		Password: Password,
	}

}
