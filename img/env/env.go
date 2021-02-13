package env

import (
	"fmt"
	"os"
	"reflect"
)

// ENV and DEFAULT ...
const (
	ENV     = "env"
	DEFAULT = "default"
)

// Configuration variable
type Configuration struct {
	DBPort           string `env:"port" default:"5342"`
	DBHost           string `env:"db_host" default:"localhost"`
	PostgresPass     string `env:"postgres_pass" default:"postgres"`
	PostgresUserName string `env:"postgres_username" default:"postgres"`
	PostgresDB       string `env:"postgres_db" default:"moviwiki"`

	ServerAPIPort string `env:"server_api_port" default:":3000"`

	SALT string `env:"salt" default:"MOVIWIKIPASS"`

	SessionKey string `env:"session_key" default:"moviwiki"`

	TokenKey string `env:"session_key" default:"moviwiki"`

	Domain string `env:"domain" default:"localhost"`
}

var serviceConfig Configuration

func setConfig() {

	v := reflect.ValueOf(serviceConfig)
	for i := 0; i < v.NumField(); i++ {

		tag := v.Type().Field(i).Tag.Get(ENV)
		defaultTag := v.Type().Field(i).Tag.Get(DEFAULT)

		if tag == "" || tag == "-" {
			continue
		}
		// a := reflect.Indirect(reflect.ValueOf(serviceConfig))
		EnvVar, _ := loadFromEnv(tag, defaultTag)
		// if Info != "" {
		// 	fmt.Println("Missing environment configuration for '" + a.Type().Field(i).Name + "', Loading default setting!")
		// }

		reflect.ValueOf(&serviceConfig).Elem().Field(i).SetString(EnvVar)
	}
}

func loadFromEnv(tag string, defaultTag string) (string, string) {

	envVar := os.Getenv(tag)
	if envVar == "" {
		envVar = defaultTag
		/* '1' is used to indicate that default value is being loaded */
		return envVar, "1"
	}
	return envVar, ""
}

// GetConfiguration used to export ENV
func GetConfiguration() Configuration {
	return serviceConfig
}

func init() {
	setConfig()
	fmt.Println("Init ENV ...")

}
