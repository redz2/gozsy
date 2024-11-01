package main

import (
	"context"
	"net/http"

	"github.com/redz2/gozsy/pkg/datasource/victoriametrics"
)

func main() {
	client := &http.Client{}
	vm := victoriametrics.NewVMStorage(client, "http://vmselect.ops.lixiangoa.com/select/0/prometheus/api/v1/query?query")
	// vm := vm.NewVMStorage(client, "http://vmselect.ops.lixiangoa.com/select/0/prometheus/api/v1/query?query")
	ctx := context.Context(context.Background())
	res, err := vm.Query(ctx, "up")
	if err != nil {
		return
	}
	print(res)

}
