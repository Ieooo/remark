package conf

import "gopkg.in/ini.v1"

var Conf = new(AppConfig)

type AppConfig struct {
	Port         int `ini:"port"`
	*MysqlConfig `ini:"mysql"`
}

type MysqlConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database string `ini:"database"'`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
