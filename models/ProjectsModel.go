package models

type Project struct {
	Id          int    `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	createdAt   string `json:"createdAt" gorm:"autoCreateTime"`
	updatedAt   string `json:"updatedAt"`
	Title       string `json:"title" gorm:"unique"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Link        string `json:"link"`
	GithubLink  string `json:"githubLink"`
	image       string `json:"image"`
	Tags        []Tag  `gorm:"many2many:project_tags;"`
}
