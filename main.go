package main

import (
	"context"
	"flag"
	"github.com/azkanurhuda/golang-absent/infrastructure/absentuser"
	"github.com/azkanurhuda/golang-absent/infrastructure/jwt"
	"github.com/azkanurhuda/golang-absent/infrastructure/user"
	"github.com/azkanurhuda/golang-absent/interfaces/handler"
	"github.com/azkanurhuda/golang-absent/interfaces/router"
	"github.com/azkanurhuda/golang-absent/pkg/conf"
	"github.com/azkanurhuda/golang-absent/pkg/postgres"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	conn, err := postgres.NewCon()
	if err != nil {
		panic(err)
	}

	userRepository := user.NewUserRepository(conn)
	jwtService := jwt.NewService()
	absentUserRepoitory := absentuser.NewAbsentUserRepository(conn)

	h := handler.NewHandler(conn, userRepository, jwtService, absentUserRepoitory)
	r := router.NewRouter(h)

	srv := &http.Server{
		Addr:         conf.Server.Addr(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		log.Infof(" ⇨ http server started on %s", conf.Server.Addr())
		log.Infof(" ⇨ graceful timeout: %s", wait)
		if err = srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Info("received stop signal")

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer func() {
		log.Info("cancel")
		cancel()
	}()

	_ = srv.Shutdown(ctx)
	log.Infof(" ⇨ shutting down")
	os.Exit(0)
}
