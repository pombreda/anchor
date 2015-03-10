package anchor

import (
	"time"
)

type Stage struct {
	UUID    string    `json:"uuid"`
	Name    string    `json:"name"`
	Job     string    `json:"job"`
	Created time.Time `json:"created"`
}

type CreateStageOptions struct {
	Name *string `json:"name" binding:"Required"`
	Job  *string `json:"job"`
}

func (opts CreateStageOptions) PrintName(name string) string {
	return map[string]string{
		"Name": "Stage name",
	}[name]
}
