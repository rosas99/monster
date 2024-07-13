package checker

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/store"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
	"time"
)

type MessageCountForMobileRule struct {
	DS  store.IStore
	RDS *redis.Client
}

func NewMessageCountForMobileRule(DS store.IStore, RDS *redis.Client) *MessageCountForMobileRule {
	return &MessageCountForMobileRule{DS: DS, RDS: RDS}
}

var _ Rule = (*MessageCountForMobileRule)(nil)

func (m *MessageCountForMobileRule) isValid(ctx context.Context, rq *types.Request) bool {
	start := time.Now().Unix()
	key := factory.WrapperMobileCount(rq.Mobile, rq.TemplateCode)
	rds := m.RDS

	sentCount, err := rds.Incr(ctx, key).Result()
	if err != nil {
		log.Errorf("Failed to increment count for key: %s, error: %v", key, err)
		return false
	}

	if sentCount == 1 {
		filter := map[string]any{
			"mobile": rq.Mobile,
			"status": "OK",
		}
		count, _, dbErr := m.DS.Histories().List(ctx, rq.TemplateCode, meta.WithFilter(filter))
		if dbErr != nil {
			log.Errorf("Failed to list histories from database for mobile: %s, error: %v", rq.Mobile, dbErr)
			return false
		}
		if count == 0 {
			if err := rds.SetNX(ctx, key, 1, types.LimitLeftTime).Err(); err != nil {
				log.Errorf("Failed to set key with expiration for key: %s, error: %v", key, err)
				return false
			}
		}
	}

	log.Infof("TemplateAndMobileChecker-----checker time效验号码模板总时间----: %d", time.Now().Unix()-start)

	isValid := sentCount <= rq.LimitValue
	if !isValid {
		log.Infow(":warning:", "key", key, "sentCount", sentCount, "isValid", isValid)
	}
	return isValid
}

func (m *MessageCountForMobileRule) getFailReason() error {
	return errors.New("exceed_limit_for_this_phone")
}
