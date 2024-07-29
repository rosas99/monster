package message

import (
	"context"
	"encoding/json"
	"github.com/rosas99/monster/internal/pkg/meta"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
)

// AILIYUNReport updates the message processing history reports for Aliyun.
func (b *messageBiz) AILIYUNReport(ctx context.Context, rq *v1.AILIYUNReportListRequest) error {
	for _, item := range rq.AILIYUNReportList {
		filter := map[string]any{"message_id": item.BizId}
		count, list, _ := b.ds.Histories().List(ctx, meta.WithFilter(filter))
		if count > 0 {
			history := list[0]
			marshal, err := json.Marshal(history)
			if err != nil {

				log.C(ctx).Warnf("Failed to marshal history record: %v", err)
				return err
			}
			history.Report = string(marshal)
			err = b.ds.Histories().Update(ctx, history)
			if err != nil {
				log.C(ctx).Warnf("Failed to update history report: %v", err)
			}
		}
	}
	return nil
}
