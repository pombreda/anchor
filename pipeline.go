package anchor

import (
	"time"
)

type Pipeline struct {
	UUID     string    `json:"uuid"`
	Name     string    `json:"name"`
	Stages   []string  `json:"stages"`
	Requires []string  `json:"requires"`
	Created  time.Time `json:"created"`
}

type CreatePipelineOptions struct {
	Name     *string  `json:"name" binding:"Required"`
	Stages   []string `json:"stages"`
	Requires []string `json:"requires"`
}

func (opts CreatePipelineOptions) PrintName(name string) string {
	return map[string]string{
		"Name":     "Pipeline name",
		"Stages":   "Stages",
		"Requires": "Prerequisites",
	}[name]
}
