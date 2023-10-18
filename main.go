package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/zaenalarifin12/simplebank/api"
	db "github.com/zaenalarifin12/simplebank/db/sqlc"
	"github.com/zaenalarifin12/simplebank/gapi"
	"github.com/zaenalarifin12/simplebank/utils"
	"log"
)

func main() {

	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("can't load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect db: ", err)
	}

	store := db.NewStore(conn)
	go gapi.RunGatewayServer(config, store)
	gapi.RunGrpcServer(config, store)

}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("can't start server", err)
	}
}
