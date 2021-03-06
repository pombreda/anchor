package anchor

import (
	"bytes"
	"encoding/json"
	"time"
)

type Flow struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Pipelines []string  `json:"pipelines"`
	Created   time.Time `json:"created"`
}

type CreateFlowOptions struct {
	Name      *string  `json:"name" binding:"Required"`
	Pipelines []string `json:"pipelines"`
}

func (opts CreateFlowOptions) PrintName(name string) string {
	return map[string]string{
		"Name":      "Flow name",
		"Pipelines": "Pipelines",
	}[name]
}

// CreateFlow creates a flow by given name.
func (c *Client) CreateFlow(opts CreateFlowOptions) (*Flow, error) {
	body, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}
	flow := new(Flow)
	return flow, c.getParsedResponse("POST", "/flows", jsonHeader, bytes.NewReader(body), flow)
}
