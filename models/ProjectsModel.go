package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title       string `json:"title" gorm:"unique"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Link        string `json:"link"`
	GithubLink  string `json:"githubLink"`
	image       string `json:"image"`
	Tags        []Tag  `gorm:"many2many:project_tags;"`
}
