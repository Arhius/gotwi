package types

import "github.com/michimani/gotwi/resources"

type ListOutput struct {
	Data   []resources.Trend        `json:"data"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
