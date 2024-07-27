package checker

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/errno"
	"github.com/rosas99/monster/internal/sms/store"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
	"time"
)

type MessageCountForTemplateRule struct {
	DS  store.IStore
	RDS *redis.Client
}

func NewMessageCountForTemplateRule(DS store.IStore, RDS *redis.Client) *MessageCountForTemplateRule {
	return &MessageCountForTemplateRule{DS: DS, RDS: RDS}
}

var _ Rule = (*MessageCountForTemplateRule)(nil)

// isValid verifies if the message count for a given template code is within the allowed limit.
func (m *MessageCountForTemplateRule) isValid(ctx context.Context, rq *Request) error {

	start := time.Now().Unix()
	key := factory.WrapperTemplateCount(rq.TemplateCode)
	sentCount, err := m.RDS.Incr(ctx, key).Result()
	if err != nil {
		log.Errorf("Failed to increment count for key: %s, error: %v", key, err)
		return err
	}

	if sentCount == 1 {
		err = m.RDS.Expire(ctx, key, types.LimitLeftTime).Err()
		if err != nil {
			// 处理错误
			log.Fatalf("Error setting expiration for key: %v", err)
		}
	}
	log.Infof("TemplateAndMobileChecker-----checker time效验号码模板总时间----: %d", time.Now().Unix()-start)
	isValid := sentCount <= rq.LimitValue
	if !isValid {
		log.Errorf(":warning:", "key", key, "sentCount", sentCount, "isValid", isValid)
		return errno.ErrTemplateCount
	}
	return nil
}
