package rickandmortyapi

import (
	"fmt"
	"io"
	"net/http"
)

func req(url string) (*[]byte, error) {
	fmt.Printf("url: %v\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}
