package entities

import "fmt"

type News struct {
	Id      int64  `json:"id"`
	Author  string `json:"author"`
	Body    string `json:"body"`
	Created string `json:"created"`
}

func (news News) ToString() string {
	return fmt.Sprintf("id: %d\nauthor: %s\nbody: %s\ncreated: %s",
		news.Id, news.Author, news.Body, news.Created)
}
