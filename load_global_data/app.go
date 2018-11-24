package main

import (
	"fmt"
	"log"
	"ui-mockup-backend"
	"ui-mockup-backend/config"
	"ui-mockup-backend/server"
	"ui-mockup-backend/mongo"
)

type App struct {
	server  *server.Server
	session *mongo.Session
	config  *root.Config
}

func (a *App) Initialize() {
	a.config = config.GetConfig()
	var err error
	a.session, err = mongo.NewSession(a.config.Mongo)
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}

	userService := mongo.NewUserService(a.session.Copy(), a.config.Mongo)
	stdService := mongo.NewStandardsService(a.session.Copy(), a.config.Mongo)
	a.server = server.NewServer(stdService, userService, a.config)
}

func (a *App) Run() {
	fmt.Println("Run")
	defer a.session.Close()
	a.server.Start()
	LoadStandards()
}
