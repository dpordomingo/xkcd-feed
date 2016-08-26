package models

import "fmt"

type Post struct {
	Title  string
	Year   string
	Month  string
	Day    string
	ImgSrc string `json:"img"`
	Id     int    `json:"num"`
}

func (p *Post) getLink() string {
	return fmt.Sprintf("http://xkcd.com/%d", p.Id)
}

func (p *Post) GetDate() string {
	return fmt.Sprintf("%s-%s-%s", p.Year, p.Month, p.Day)
}
