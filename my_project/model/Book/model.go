package Book

import "my_project/model/author"

type Book struct {
	Id         uint          `json:"Id"`
	Name       string        `json:"Name"`
	PublicDate string        `json:"publicDate"`
	AboutWhat  string        `json:"aboutWhat"`
	Author     author.Author `gorm:"embedded"`
}
