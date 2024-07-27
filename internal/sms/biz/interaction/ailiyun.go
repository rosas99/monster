package interaction

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/sms/types"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// AILIYUNUplink processes uplink messages received from AILIYUN and logs them.
func (b *interactionBiz) AILIYUNUplink(ctx context.Context, rq *v1.AILIYUNUplinkListRequest) error {
	for _, item := range rq.AILIYUNCallbackList {
		var msgRequest types.UplinkMsgRequest
		_ = copier.Copy(msgRequest, item)
		b.logger.WriteUplinkMessage(ctx, &msgRequest)
	}

	return nil
}
