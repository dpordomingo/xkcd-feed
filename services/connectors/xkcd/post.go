package xkcd

import (
	"fmt"
	"strconv"

	"github.com/dpordomingo/learning-exercises/xkcd.com/models"
)

type Post struct {
	Title  string
	ID     int    `json:"num"`
	ImgSrc string `json:"img"`
	Year   string
	Month  string
	Day    string
}

func (p *Post) GetTitle() string {
	return p.Title
}

func (p *Post) GetID() string {
	return fmt.Sprintf("xkcd_%d", p.ID)
}

func (p *Post) GetImgSrc() string {
	return p.ImgSrc
}

func (p *Post) GetDate() *models.Date {
	year, _ := strconv.Atoi(p.Year)
	month, _ := strconv.Atoi(p.Month)
	day, _ := strconv.Atoi(p.Day)
	return &models.Date{year, month, day}
}

func (p *Post) GetLink() string {
	return fmt.Sprintf("http://xkcd.com/%d", p.ID)
}
