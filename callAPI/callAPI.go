package callapi

import (
	"fmt"
	"io"
	"net/http"
)

func GetPokeInfo(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("Your argument was not found!")
	}

	reader, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	return reader, nil
}
