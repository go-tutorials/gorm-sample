package app

import (
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	sv "github.com/core-go/service"
	"github.com/core-go/sql"
)

type Config struct {
	Server     sv.ServerConf    `mapstructure:"server"`
	Sql        sql.Config       `mapstructure:"sql"`
	Log        log.Config       `mapstructure:"log"`
	MiddleWare mid.LogConfig    `mapstructure:"middleware"`
}
