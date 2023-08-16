package main

import (
	"context"
	"fmt"
	"github.com/reward-rabieth/series-troxide/config"
	"github.com/reward-rabieth/series-troxide/db"
	"github.com/reward-rabieth/series-troxide/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

const (
	defaultConfigName   = "config.default"
	defaultConfigFolder = "."
)

func main() {
	err := readConfiguration()
	if err != nil {
		log.WithError(err).Fatalln("failed loading configuration")
	}

	cfg := config.GetDatabaseConfig()
	conn, err := db.Connect(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := conn.Client().Disconnect(context.Background()); err != nil {
			log.Error(err)
		}
	}()

	seriesRepo := db.NewMongoDBShowRepo(conn)
	episodeRepo := db.NewMongoDbEpisodeRepo(conn)
	client := server.NewClient()
	apiHandler := server.NewApiHandler(seriesRepo, episodeRepo, client)
	if err != nil {
		log.WithError(err).Fatalln("failed creating HTTP API handler")
	}
	addr := fmt.Sprintf(":%s", viper.GetString("port"))
	log.Infof("running http server on port: %s", addr)
	apiHandler.RunServer(addr)

}

func readConfiguration() error {
	viper.SetConfigName(defaultConfigName)
	viper.AddConfigPath(defaultConfigFolder)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()
	return nil
}
