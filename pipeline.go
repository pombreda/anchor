package anchor

import (
	"time"
)

type Pipeline struct {
	UUID     string    `json:"uuid"`
	Name     string    `json:"name"`
	Requires []string  `json:"requires"`
	Created  time.Time `json:"created"`
}

type CreatePipelineOptions struct {
	Name     *string  `json:"name" binding:"Required"`
	Requires []string `json:"requires"`
}

func (opts CreatePipelineOptions) PrintName(name string) string {
	return map[string]string{
		"Name":     "Pipeline name",
		"Requires": "Prerequisites",
	}[name]
}
