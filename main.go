package main

import (
	"context"
	"fmt"
	"github.com/core-go/config"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	sv "github.com/core-go/service"
	"github.com/gorilla/mux"

	"go-service/internal/app"
)

func main() {
	var conf app.Config
	er1 := config.Load(&conf, "configs/config")
	if er1 != nil {
		panic(er1)
	}
	r := mux.NewRouter()

	log.Initialize(conf.Log)
	r.Use(mid.BuildContext)
	logger := mid.NewLogger()
	if log.IsInfoEnable() {
		r.Use(mid.Logger(conf.MiddleWare, log.InfoFields, logger))
	}
	r.Use(mid.Recover(log.PanicMsg))

	er2 := app.Route(r, context.Background(), conf)
	if er2 != nil {
		panic(er2)
	}
	fmt.Println(sv.ServerInfo(conf.Server))
	server := sv.CreateServer(conf.Server, r)
	if er3 := server.ListenAndServe(); er3 != nil {
		fmt.Println(er3.Error())
	}
}
