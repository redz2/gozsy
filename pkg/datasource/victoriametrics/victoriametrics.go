package datasource

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type response struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Labels map[string]string `json:"metric"`
			TV     [2]interface{}    `json:"value"`
		} `json:"result"`
	} `json:"data"`
	ErrorType string `json:"errorType"`
	Error     string `json:"error"`
}

type request struct {
	promql    string
	starttime time.Time
	endtime   time.Time
	step      time.Duration
}

type VMStorage struct {
	c   *http.Client
	url string
}

func NewVMStorage(c *http.Client, url string) *VMStorage {
	return &VMStorage{
		c:   c,
		url: url,
	}
}

func (v *VMStorage) Query(ctx context.Context, promql string) (*response, error) {

	req, err := http.NewRequest("POST", v.url+url.QueryEscape(promql), nil)
	if err != nil {
		return nil, err
	}
	resq, err := v.c.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resq.Body.Close()

	var r *response
	err = json.NewDecoder(resq.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r, nil

}
