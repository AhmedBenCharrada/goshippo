package shippo

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type ObjectFields struct {
	ObjectCreated *time.Time `json:"object_created,omitempty"`
	ObjectUpdated *time.Time `json:"object_updated,omitempty"`
	ObjectID      string     `json:"object_id,omitempty"`
	ObjectOwner   string     `json:"object_owner,omitempty"`
	ObjectState   string     `json:"object_state,omitempty"`
}

type OutputMessage struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Text    string `json:"text,omitempty"`
	Source  string `json:"source,omitempty"`
}

type List[T any] struct {
	Count           int     `json:"count"`
	NextPageURL     *string `json:"next"`
	PreviousPageURL *string `json:"previous"`
	Results         []T     `json:"results"`
}

type DateFilter struct {
	Gt  *time.Time `json:"object_created_gt"`
	Gte *time.Time `json:"object_created_gte"`
	Lt  *time.Time `json:"object_created_lt"`
	Lte *time.Time `json:"object_created_lte"`
}

func (d DateFilter) ToPathParam() string {
	if d.Gt == nil && d.Gte == nil && d.Lt == nil && d.Lte == nil {
		return ""
	}

	p := ""
	if d.Gt != nil {
		p += "object_created_gt=" + d.Gt.Format("2006-01-02T03:30:30.5") + "&"
	}
	if d.Gte != nil {
		p += "object_created_gte=" + d.Gte.Format("2006-01-02T03:30:30.5") + "&"
	}
	if d.Lt != nil {
		p += "object_created_lt=" + d.Lt.Format("2006-01-02T03:30:30.5") + "&"
	}
	if d.Lte != nil {
		p += "object_created_lte=" + d.Lte.Format("2006-01-02T03:30:30.5") + "&"
	}

	return p
}

func create[K any, T any](ctx context.Context, client Client, url string, in *K) (*T, error) {
	if in == nil {
		return nil, fmt.Errorf("nil input")
	}

	// Todo: validate required fields

	req, err := client.NewRequest(ctx, http.MethodPost, url, in)
	if err != nil {
		return nil, err
	}

	var item T
	if _, err := client.Do(req, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func list[T any](ctx context.Context, client Client, url string, params *string, page, size uint) ([]T, error) {
	if size == 0 {
		size = 25
	}
	if page == 0 {
		page = 1
	}

	url = fmt.Sprintf("%s?page=%v&results=%v", url, page, size)
	if params != nil {
		url = fmt.Sprintf("%s?%s&page=%v&results=%v", url, *params, page, size)
	}

	req, err := client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var res List[T]
	if _, err := client.Do(req, &res); err != nil {
		return nil, err
	}

	return res.Results, nil
}

func get[T any](ctx context.Context, client Client, url, id string, r ...string) (*T, error) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, fmt.Errorf("empty ID")
	}
	req, err := client.NewRequest(ctx, http.MethodGet, url+id+strings.Join(r, "/"), nil)
	if err != nil {
		return nil, err
	}

	var item T
	if _, err := client.Do(req, &item); err != nil {
		return nil, err
	}

	return &item, nil
}
