package server

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/reward-rabieth/series-troxide/db"
	"github.com/reward-rabieth/series-troxide/types"
	"net/http"
)

var (
	showList    []types.Show
	episodeList []types.Episode
)

type ShowHandler struct {
	SeriesRepo db.ShowHandler
	Client     Client
}

func NewShowHandler(SeriesRepo db.ShowHandler, Client Client) *ShowHandler {
	return &ShowHandler{
		SeriesRepo: SeriesRepo,
		Client:     Client,
	}
}

func (sh *ShowHandler) fetchShowInformation() ([]types.Show, error) {
	url := BaseUrlWithPath("/shows")
	status, err := sh.Client.Get(url, &showList)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch  data from API: status %d", status)

	}

	return showList, nil
}
func (sh *ShowHandler) FetchEpisodeForShow(c *fiber.Ctx) ([]types.Episode, error) {
	id := c.Path("id")
	url := BaseUrlWithPath(fmt.Sprintf("/shows/%s/episodes", id))
	status, err := sh.Client.Get(url, &episodeList)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data from APi: status %d", status)
	}
	return episodeList, nil

}

func (sh *ShowHandler) ShowInformationHandler(c *fiber.Ctx) error {
	showList, err := sh.fetchShowInformation()
	if err != nil {
		return err
	}
	err = sh.SeriesRepo.StoreShows(context.TODO(), showList)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"shows": showList})
}
