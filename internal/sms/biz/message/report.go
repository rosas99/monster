package message

import (
	"context"
	"encoding/json"
	"github.com/rosas99/monster/internal/pkg/meta"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// AILIYUNReport receives AILI cloud message reports and link them to their sending history.
func (b *messageBiz) AILIYUNReport(ctx context.Context, rq *v1.AILIYUNReportListRequest) error {
	for _, item := range rq.AILIYUNReportList {
		filter := make(map[string]any)
		filter["message_id"] = item.BizId
		count, list, _ := b.ds.Histories().List(ctx, "", meta.WithFilter(filter))
		if count > 0 {
			history := list[0]
			marshal, _ := json.Marshal(history)
			history.Report = string(marshal)
		}
	}
	return nil
}
