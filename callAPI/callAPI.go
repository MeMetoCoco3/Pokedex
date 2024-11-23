package callapi

import (
	_ "fmt"
	"io"
	"net/http"
)

func GetPokeInfo(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	reader, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	return reader, nil
}
