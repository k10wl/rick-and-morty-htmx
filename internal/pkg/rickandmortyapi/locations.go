package rickandmortyapi

import "encoding/json"

type Location struct {
	Id        int
	Name      string
	Type      string
	Dimension string
	Residents []string
	Url       string
	Created   string
}

type Locations struct {
	Info    Info
	Results []Location
}

func GetLocations(query string) (*Locations, error) {
	b, err := req(baseUrl + location + "?" + query)
	if err != nil {
		return nil, err
	}

	var res Locations
	if err := json.Unmarshal(*b, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
