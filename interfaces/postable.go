package interfaces

import "github.com/dpordomingo/learning-exercises/xkcd.com/models"

type Postable interface {
	GetTitle() string
	GetID() string
	GetImgSrc() string
	GetDate() *models.Date
	GetLink() string
}
