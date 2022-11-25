package model

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Title        string
	Content      string
	UserID       string
	User         User
	TopicOptions []TopicOption
}

type TopicOption struct {
	gorm.Model
	TopicID int
	Detail  string
	Topic   Topic
}

type Vote struct {
	gorm.Model
	VoteID int
	UserID int
	User   User
}
