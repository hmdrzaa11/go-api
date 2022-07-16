package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hmdrzaa11/got/pkg/config"
)

type Application struct {
	Server   *http.Server
	Router   *http.ServeMux
	Logger   *log.Logger
	Config   *config.Config
	Database *sql.DB
}

func (a *Application) Run() {
	a.Logger.Println("listening on port", a.Config.Port)
	err := a.Server.ListenAndServe()
	if err != nil {
		a.Logger.Fatal(err)
	}

}

func (a *Application) GracefullShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	a.Logger.Println("received intruption signals, gracefully shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	a.Server.Shutdown(ctx)
	os.Exit(0)
}

func (a *Application) SendErrorResponse(w http.ResponseWriter, status int, err error) {
	customeErr := struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(customeErr)
	if err != nil {
		panic(err)
	}
}

func (a *Application) SendResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
		return
	}
}

func Boot() *Application {
	config := config.NewConfig()
	mux := http.NewServeMux()
	logger := log.New(os.Stdout, fmt.Sprintf("%s: ", config.Name), log.LstdFlags)

	server := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 10,
	}
	return &Application{
		Server: server,
		Router: mux, //if you want to apply a middleware to entire application this is the place to put it
		Logger: logger,
		Config: config,
	}
}
