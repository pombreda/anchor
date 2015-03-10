package anchor

import (
	"time"
)

type Job struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

type CreateJobOptions struct {
	Name    *string `json:"name" binding:"Required"`
	Content *string `json:"content"`
}

func (opts CreateJobOptions) PrintName(name string) string {
	return map[string]string{
		"Name":    "Job name",
		"Content": "Job content",
	}[name]
}
