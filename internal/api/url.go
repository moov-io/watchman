package api

import (
	"net/url"
	"slices"
	"strings"
)

type QueryParams struct {
	url.Values

	valuesRead []string
}

func NewQueryParams(u *url.URL) *QueryParams {
	return &QueryParams{
		Values: u.Query(),
	}
}

func (q *QueryParams) WithPrefix(prefix string) []string {
	var out []string

	for key := range q.Values {
		if strings.HasPrefix(key, prefix) {
			out = append(out, key)
		}
	}
	slices.Sort(out)

	return out
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
