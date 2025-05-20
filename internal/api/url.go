package api

import (
	"net/url"
	"slices"
)

type QueryParams struct {
	url.Values

	valuesRead []string
}

func (q *QueryParams) Get(name string) string {
	q.valuesRead = append(q.valuesRead, name)

	return q.Values.Get(name)
}

func (q *QueryParams) GetAll(name string) []string {
	q.Get(name)

	return q.Values[name]
}

func (q *QueryParams) UnusedQueryParams() []string {
	var extra []string

	for key := range q.Values {
		if !slices.Contains(q.valuesRead, key) {
			extra = append(extra, key)
		}
	}

	slices.Sort(extra)
	return extra
}
