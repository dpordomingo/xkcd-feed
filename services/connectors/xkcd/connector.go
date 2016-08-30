package xkcd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dpordomingo/learning-exercises/xkcd.com/interfaces"
	"github.com/dpordomingo/learning-exercises/xkcd.com/services"
)

type Connector struct{}

func (x *Connector) Retrieve(id interface{}) (string, error) {
	url := getPostJsonUrl(id.(int))

	return services.HTTPget(url)
}

func (x *Connector) Hydrate(input string) (interfaces.Postable, error) {
	if len(input) == 0 {
		return nil, errors.New("No string passed")
	}
	var post Post
	bytes := []byte(input)
	if err := json.Unmarshal(bytes, &post); err != nil {
		return nil, errors.New("Wrong JSON format")
	}
	return &post, nil
}

func getPostJsonUrl(id int) string {
	return fmt.Sprintf("http://xkcd.com/%d/info.0.json", id)
}
