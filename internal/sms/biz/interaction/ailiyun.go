package interaction

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/sms/types"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// AILIYUNCallback receives an uplink message and writes to kafka queue.
func (b *interactionBiz) AILIYUNCallback(ctx context.Context, rq *v1.AILIYUNCallbackListRequest) (*v1.CommonResponse, error) {
	for _, item := range rq.AILIYUNCallbackList {
		var msgRequest types.UplinkMsgRequest
		err := copier.Copy(msgRequest, item)
		if err != nil {
			return nil, err
		}
		b.logger.WriteUplinkMessage(ctx, &msgRequest)

	}
	// log kpi

	return &v1.CommonResponse{Code: 0, Msg: "success"}, nil
}
