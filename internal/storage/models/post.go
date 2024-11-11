package models

type Post struct {
	Base
	Title   string `gorm:"not null"`
	Content string
	Upvote  int
}
