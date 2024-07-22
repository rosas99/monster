package meta

const (
	// AppName is the default argument to specify on a context when you want to list or filter resources across all scopes.
	AppName = "default_app"
	KpiName = "default_kpi"
)

type KpiOption func(*KpiOptions)

type KpiOptions struct {
	Kpi map[string]any
}

func NewKpiOptions(opts ...KpiOption) KpiOptions {
	los := KpiOptions{
		Kpi: map[string]any{
			"appName": AppName,
			"kpiName": KpiName,
			"code":    200,
			"message": "success",
		},
	}

	for _, opt := range opts {
		opt(&los)
	}

	return los
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

func WithStatus(status string) KpiOption {
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
