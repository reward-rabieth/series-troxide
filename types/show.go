package types

import "time"

type Show struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Genres  []string `json:"genres"`
	Summary string   `json:"summary"`
	Status  string   `json:"status"`
	Embeds  struct {
		Episodes []Episode
	} `json:"_embedded"`
}

type Episode struct {
	ID       int       `json:"id"`
	URL      string    `json:"url"`
	Name     string    `json:"name"`
	Season   int       `json:"season"`
	Number   int       `json:"number"`
	Type     string    `json:"type"`
	Airdate  string    `json:"airdate"`
	Airtime  string    `json:"airtime"`
	Airstamp time.Time `json:"airstamp"`
	Runtime  int       `json:"runtime"`
	Rating   struct {
		Average float64 `json:"average"`
	} `json:"rating"`
	Summary string `json:"summary"`
	Links   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Show struct {
			Href string `json:"href"`
		} `json:"show"`
	} `json:"_links"`
}
type EpisodeSchedule struct {
	Episode  []Episode `json:"episode_description"`
	Time     string    `json:"time"`
	Days     []string  `json:"days"`
	Name     string    `json:"name"`
	Code     string    `json:"code"`
	Timezone string    `json:"timezone"`
}
