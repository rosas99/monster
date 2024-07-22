package monitor

import (
	"context"
	"encoding/json"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
)

const (
	// AppName is the appName name when starting monster-sms server.
	AppName = "monster-sms"
)

// LogKpi writes a log message for the api request.
func (i *impl) LogKpi(kpiName, traceId, status, templateCode string, costTime int64) {
	extra := map[string]any{"template_code": templateCode}

	kpi := meta.NewKpiOptions(meta.WithAppName(AppName), meta.WithKpiName(kpiName), meta.WithTraceId(traceId),
		meta.WithStatus(status), meta.WithCostTime(costTime), meta.WithExtra(extra)).Kpi

	out, _ := json.Marshal(kpi)
	if err := i.writer.WriteMessages(context.Background(), kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	}
}
