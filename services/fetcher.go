package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dpordomingo/learning-exercises/xkcd.com/models"
)

func RetrievePosts(limit int, offset int) []*models.Post {
	var result []*models.Post
	for i := offset; i < offset+limit; i++ {
		post, err := getPost(i)
		if err != nil {
			fmt.Println(fmt.Errorf("#%d :: %s", i, err.Error()))
			continue
		}
		result = append(result, post)
	}
	return result
}

func getPost(id int) (*models.Post, error) {
	postString, err := xkcdPostRetriever(id)
	if err != nil {
		return nil, err
	}
	return stringToXkcdPost(postString)
}

func xkcdPostRetriever(id int) (string, error) {
	var retrievedPost []byte
	url := getPostJsonUrl(id)

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil || resp.StatusCode != http.StatusOK {
		return "", errors.New("No post found")
	}
	retrievedPost, err = ioutil.ReadAll(resp.Body)
	return string(retrievedPost), err
}

func stringToXkcdPost(input string) (*models.Post, error) {
	if len(input) == 0 {
		return nil, errors.New("No string passed")
	}
	var post models.Post
	bytes := []byte(input)
	if err := json.Unmarshal(bytes, &post); err != nil {
		return nil, errors.New("Wrong JSON format")
	}
	return &post, nil
}

func getPostJsonUrl(id int) string {
	return fmt.Sprintf("http://xkcd.com/%d/info.0.json", id)
}
