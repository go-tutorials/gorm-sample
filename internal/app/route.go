package app

import (
	"context"
	. "github.com/core-go/service"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, ctx context.Context, conf Config) error {
	app, err := NewApp(ctx, conf)
	if err != nil {
		return err
	}
	r.HandleFunc("/health", app.Health.Check).Methods(GET)

	user := "/users"
	r.HandleFunc(user, app.User.All).Methods(GET)
	r.HandleFunc(user+"/search", app.User.Search).Methods(GET, POST)
	r.HandleFunc(user+"/{id}", app.User.Load).Methods(GET)
	r.HandleFunc(user, app.User.Create).Methods(POST)
	r.HandleFunc(user+"/{id}", app.User.Update).Methods(PUT)
	r.HandleFunc(user+"/{id}", app.User.Patch).Methods(PATCH)
	r.HandleFunc(user+"/{id}", app.User.Delete).Methods(DELETE)

	return nil
}
