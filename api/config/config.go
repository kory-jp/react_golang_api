package config

import (
	"log"

	"github.com/kory-jp/react_golang_api/api/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	LogFile   string
	SQLDriver string
	UserName  string
	Password  string
	DBPort    string
	DBname    string
	Static    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		UserName:  cfg.Section("db").Key("user_name").String(),
		Password:  cfg.Section("db").Key("password").String(),
		DBPort:    cfg.Section("db").Key("port").String(),
		DBname:    cfg.Section("db").Key("db_name").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
