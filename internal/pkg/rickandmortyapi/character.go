package rickandmortyapi

import (
	"encoding/json"
)

type Referance struct {
	Name string
	Url  string
}

type Character struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	Species  string    `json:"species"`
	Type     string    `json:"type"`
	Gender   string    `json:"gender"`
	Origin   Referance `json:"origin"`
	Location Referance `json:"location"`
	Image    string    `json:"image"`
	Episode  []string  `json:"episode"`
	Url      string    `json:"url"`
	Created  string    `json:"created"`
}

type Characters struct {
	Info    Info        `json:"info"`
	Results []Character `json:"results"`
}

func GetCharacters(params string) (*Characters, error) {
	b, err := req(baseUrl + characters + "?" + params)
	if err != nil {
		return nil, err
	}
	var c Characters
	err = json.Unmarshal(*b, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
