package meta

type KpiOption func(*KpiOptions)

type KpiOptions struct {
	Kpi map[string]any
}

func NewKpiOptions(opts ...KpiOption) KpiOptions {
	kos := KpiOptions{
		Kpi: map[string]any{
			"code":    200, // default value
			"message": "success",
		},
	}

	for _, opt := range opts {
		opt(&kos)
	}

	return kos
}

func WithAppName(appName string) KpiOption {
	return func(o *KpiOptions) {
		o.Kpi["appName"] = appName
	}
}

func WithKpiName(kpiName string) KpiOption {
	return func(o *KpiOptions) {
		o.Kpi["kpiName"] = kpiName
	}
}

func WithCode(code string) KpiOption {
	return func(o *KpiOptions) {
		o.Kpi["code"] = code
	}
}

func WithMessage(message string) KpiOption {
	return func(o *KpiOptions) {
		o.Kpi["message"] = message
	}
}

func WithStatus(status bool) KpiOption {
	return func(o *KpiOptions) {
		o.Kpi["status"] = status
	}
}

func WithTraceId(traceId string) KpiOption {
	return func(o *KpiOptions) {
		o.Kpi["traceId"] = traceId
	}
}

func WithCostTime(costTime int64) KpiOption {
	return func(o *KpiOptions) {
		o.Kpi["costTime"] = costTime
	}
}

func WithExtra(extra map[string]any) KpiOption {
	return func(o *KpiOptions) {
		for key, value := range extra {
			o.Kpi[key] = value
		}
	}
}
