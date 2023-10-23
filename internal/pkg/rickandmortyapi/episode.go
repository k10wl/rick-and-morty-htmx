package rickandmortyapi

import (
	"encoding/json"
)

type Episode struct {
	Id         int
	Name       string
	Air_date   string
	Episode    string
	Characters []string
	Url        string
	Created    string
}

type Episodes struct {
	Info    Info
	Results []Episode
}

func GetEpisodes(query string) (*Episodes, error) {
	d, err := req(baseUrl + episode + "?" + query)
	if err != nil {
		return nil, err
	}

	var e Episodes
	if err := json.Unmarshal(*d, &e); err != nil {
		return nil, err
	}

	return &e, nil
}
