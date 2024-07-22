package main

import (
	"encoding/json"
	"fmt"
	"github.com/rosas99/monster/internal/pkg/meta"
)

func main() {
	extra := map[string]any{"template_code": "1asd"}

	kpi := meta.NewKpiOptions(meta.WithAppName("AppName"), meta.WithKpiName("kpiName"), meta.WithTraceId("traceId"),
		meta.WithStatus("status"), meta.WithCostTime(123), meta.WithExtra(extra)).Kpi

	marshal, _ := json.Marshal(kpi)

	//fmt.Print(kpi)
	fmt.Print(string(marshal))
}
