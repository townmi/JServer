package utils

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

// TomlConfig s
type TomlConfig struct {
	Title string
	Owner OwnerInfo
	DB    Database `toml:"database"`
}

// OwnerInfo s
type OwnerInfo struct {
	Name string
	Port string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

// Database s
type Database struct {
	Server  string
	Ports   string
	ConnMax int `toml:"connection_max"`
	Enabled bool
	Dirver  string
}

// PWD s
var PWD string
var conf TomlConfig

func init() {
	environ := os.Environ()
	for i := range environ {
		es := strings.Split(environ[i], "=")
		if es[0] == "PWD" {
			PWD = es[1]
			break
		}
	}
	if PWD != "" {
		if _, err := toml.DecodeFile(PWD+"/config.toml", &conf); err != nil {
			log.Fatal(err)
		}
	}
}

// GetDataBaseConnection s
func GetDataBaseConnection() (string, string) {
	var connection string
	if conf.DB.Dirver == "mysql" {
		connection = "root:abcd1234@tcp(" + conf.DB.Server + ":" + conf.DB.Ports + ")/JHome?charset=utf8&parseTime=True&loc=Local"
	}
	return conf.DB.Dirver, connection
}

// GetServerPost s
func GetServerPost() string {
	if conf.Owner.Port != "" {
		return conf.Owner.Port
	}
	return "3000"
}
