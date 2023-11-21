// Copyright 2019 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package v1_test provides examples making requests to Prometheus using the
// Golang client.
package victoriametrics

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type QueryResult struct {
	Name  string
	Value [][]string
}

func VectorConvertSlice(v model.Vector) {

}

func VmQuery(VictoriaUrl string, PromQL string) (v model.Vector, err error) {
	client, err := api.NewClient(api.Config{
		Address: VictoriaUrl,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, PromQL, time.Now(), v1.WithTimeout(5*time.Second))
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		return
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}

	v = result.(model.Vector)
	return
}

func VmQueryRange(VictoriaUrl string, PromQL string) (v model.Value, err error) {
	client, err := api.NewClient(api.Config{
		Address: VictoriaUrl,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := v1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Minute,
	}
	result, warnings, err := v1api.QueryRange(ctx, PromQL, r, v1.WithTimeout(5*time.Second))
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		return
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}

	v = result

	// fmt.Printf("Result:\n%T\n", result)
	return
}

// {"status":"success","isPartial":false,"data":{"resultType":"vector","result":[{"metric":{"verb":"DELETE"},"value":[1695476758,"0"]},{"metric":{"verb":"GET"},"value":[1695476758,"32.93333333333333"]},{"metric":{"verb":"LIST"},"value":[1695476758,"0.26666666666666666"]},{"metric":{"verb":"PATCH"},"value":[1695476758,"0.23333333333333334"]},{"metric":{"verb":"POST"},"value":[1695476758,"0.43333333333333335"]},{"metric":{"verb":"PUT"},"value":[1695476758,"9.433333333333334"]},{"metric":{"verb":"WATCH"},"value":[1695476758,"3.6666666666666665"]}]}}
