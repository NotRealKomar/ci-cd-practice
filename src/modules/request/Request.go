package request

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

func create(title string) Request {
	return Request{
		Id:    uuid.NewString(),
		Title: title,
		Date:  time.Now().Format(time.Kitchen),
	}
}
