package main

import (
	"fmt"

	"github.com/dpordomingo/learning-exercises/xkcd.com/services"
	"github.com/dpordomingo/learning-exercises/xkcd.com/services/connectors/xkcd"
)

func main() {
	xkcdConnector := xkcd.Connector{}
	library := services.NewLibrary(&xkcdConnector)
	posts := library.Get(2, 1722)
	for _, post := range posts {
		fmt.Println(post)
	}
}
