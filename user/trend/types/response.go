package types

import "github.com/Arhius/gotwi/resources"

type ListOutput struct {
	Data   []resources.PersonalizedTrend `json:"data"`
	Errors []resources.PartialError      `json:"errors"`
}

func (r *ListOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
