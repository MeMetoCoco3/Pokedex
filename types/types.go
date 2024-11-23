package types

import "github.com/MeMetoCoco3/Pokedex/internal"

type CliCommand struct {
	Name        string
	Description string
	Function    func(*Config) error
}

type Config struct {
	PointLocation int
	Cache         *pokecache.Cache
}

type Respose struct {
	Area []struct {
		Name string `json:"name"`
	} `json:"results"`
}
