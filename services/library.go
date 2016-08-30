package services

import (
	"fmt"

	"github.com/dpordomingo/learning-exercises/xkcd.com/interfaces"
)

type Library struct {
	Connector interfaces.Connectable
}

func NewLibrary(c interfaces.Connectable) *Library {
	library := Library{}
	library.Connector = c
	return &library
}

func (l *Library) Get(limit int, offset int) []interfaces.Postable {
	var result []interfaces.Postable
	for i := offset; i < offset+limit; i++ {
		post, err := getPost(l.Connector, i)
		if err != nil {
			fmt.Println(fmt.Errorf("#%d :: %s", i, err.Error()))
			continue
		}
		result = append(result, post)
	}
	return result
}

func getPost(connectable interfaces.Connectable, id interface{}) (interfaces.Postable, error) {
	postString, err := connectable.Retrieve(id)
	if err != nil {
		return nil, err
	}
	return connectable.Hydrate(postString)
}

func (l *Library) convert() []interfaces.Postable {
	var res []interfaces.Postable
	return res
}
