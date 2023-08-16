package server

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/reward-rabieth/series-troxide/db"
	"github.com/reward-rabieth/series-troxide/types"
	"net/http"
)

const (
	episodeUrl = "https://api.tvmaze.com/shows/1/episodes"
)

type EpisodeHandler struct {
	EpisodeRepo db.EpisodeRepo
	client      Client
	show        types.Show
}

func NewEpisodeHandler(EpisodeRepo db.EpisodeRepo, client Client, show types.Show) *EpisodeHandler {
	return &EpisodeHandler{
		EpisodeRepo: EpisodeRepo,
		client:      client,
		show:        show,
	}
}

func (sh EpisodeHandler) fetchSeasonEpisodeList(c *fiber.Ctx) ([]types.Episode, error) {
	id := c.Params("id")
	var (
		episodeList []types.Episode
		url         = BaseUrlWithPath(fmt.Sprintf("/seasons/%s/episodes", id))
	)

	status, err := sh.client.Get(url, &episodeList)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data from API: status %d", status)
	}

	return episodeList, nil
}

// fetching episode that air in a given country on a given date
func fetchEpisodeSchedule(c *fiber.Ctx) ([]types.EpisodeSchedule, error) {

	return nil, nil

}

func (sh EpisodeHandler) EpisodeHandler(c *fiber.Ctx) error {
	episodeList, err := sh.fetchSeasonEpisodeList(c)
	if err != nil {
		return err
	}

	err = sh.EpisodeRepo.StoreEpisode(context.TODO(), episodeList)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"seasonEpisodeResponse": episodeList})
}
