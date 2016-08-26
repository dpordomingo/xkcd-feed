package main

import (
	"fmt"

	"github.com/dpordomingo/learning-exercises/xkcd.com/services"
)

func main() {
	posts := services.RetrievePosts(2, 1722)
	for _, post := range posts {
		fmt.Println(post)
	}

}
