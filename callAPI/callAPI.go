package callapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Area struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func GetArea(n int) ([]Area, error) {
	fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", n)
	req, err := http.Get(fullURL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer req.Body.Close()

	reader, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(reader))
	areas := make([]Area, 20)

	for areaIndex := 0; areaIndex < 20; areaIndex++ {
		var area Area
		err = json.Unmarshal(reader, &area)
		if err != nil {
			return nil, err
		}
		areas[areaIndex] = area
	}
	return areas, nil
}
