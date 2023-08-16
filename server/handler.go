package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reward-rabieth/series-troxide/db"
	"github.com/reward-rabieth/series-troxide/types"
	log "github.com/sirupsen/logrus"
)

type ApiHandler struct {
	show        types.Show
	app         *fiber.App
	client      Client
	seriesRepo  db.ShowHandler
	episodeRepo db.EpisodeRepo
}

func (apiHandler *ApiHandler) RunServer(addr string) {
	err := apiHandler.app.Listen(addr)
	if err != nil {
		log.WithError(err).Fatalln("failed to start server")
	}
}
func NewApiHandler(seriesRepo db.ShowHandler, episodeRepo db.EpisodeRepo, client Client) *ApiHandler {
	app := fiber.New()
	api := &ApiHandler{
		app:         app,
		seriesRepo:  seriesRepo,
		episodeRepo: episodeRepo,
		client:      client,
	}
	api.addShowsEndpoint()
	api.addSeasonsEndpoint()
	return api
}

func (apiHandler *ApiHandler) addShowsEndpoint() {
	seriesHandler := NewShowHandler(apiHandler.seriesRepo, apiHandler.client)

	series := apiHandler.app.Group("/shows")
	series.Get("/:id", seriesHandler.ShowInformationHandler)

}

func (apiHandler *ApiHandler) addSeasonsEndpoint() {
	episodeHandler := NewEpisodeHandler(apiHandler.episodeRepo, apiHandler.client, apiHandler.show)
	episode := apiHandler.app.Group("/seasons")
	episode.Get("/:id/episodes", episodeHandler.EpisodeHandler)

}
