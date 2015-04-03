package anchor

import (
	"net/http"
)

var jsonHeader = http.Header{"content-type": []string{"application/json"}}

func NewString(s string) *string {
	return &s
}
