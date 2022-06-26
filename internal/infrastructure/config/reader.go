package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	CONFIG_PATH            = "./config/config.yaml"
	MYSQL_PASSWORD_ENVNAME = "MYSQL_PASSWORD"
)

func ReadConfig(dir string) *Bootstrap {
	if len(dir) == 0 {
		dir = CONFIG_PATH
	}
	c := new(Bootstrap)
	f, err := ioutil.ReadFile(dir)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = yaml.Unmarshal(f, c)
	if err != nil {
		log.Fatal(err.Error())
	}
	c.Data.Db.Password = readFromEnv(MYSQL_PASSWORD_ENVNAME)
	return c
}

// read configuration from env, if empty will call os.Exit
func readFromEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Fatalf("infrastructure.config.reader: env variable [%v] is empty and it must be provided.", key)
	}
	return value
}
