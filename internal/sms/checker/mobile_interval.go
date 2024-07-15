package checker

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/sms/store"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
	"strconv"
	"time"
)

type TimeIntervalForMobileRule struct {
	DS  store.IStore
	RDS *redis.Client
}

func NewTimeIntervalForMobileRule(DS store.IStore, RDS *redis.Client) *TimeIntervalForMobileRule {
	return &TimeIntervalForMobileRule{DS: DS, RDS: RDS}
}

var _ Rule = (*TimeIntervalForMobileRule)(nil)

func (m *TimeIntervalForMobileRule) isValid(ctx context.Context, rq *types.Request) bool {
	start := time.Now().UnixMilli()
	key := factory.WrapperTimeInterval(rq.Mobile, rq.TemplateCode)

	timeStampStr, err := m.RDS.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Errorf("Failed to get timestamp from Redis for key: %s, error: %v", key, err)
		return false
	}

	if timeStampStr == "" {
		_, err2 := m.RDS.Set(ctx, key, time.Now().UnixMilli(), 24*time.Hour).Result()
		if err2 != nil {
			return false
		}
		return true
	}

	remainingTTL, err2 := m.RDS.TTL(ctx, key).Result()
	if err2 != nil {
		return false
	}
	_, err2 = m.RDS.Set(ctx, key, time.Now().UnixMilli(), remainingTTL).Result()
	if err2 != nil {
		return false
	}

	log.Infof("timeInterval --- checker time -----效验时间戳的总时间: %d", time.Now().UnixMilli()-start)

	timeStampInt, err := strconv.ParseInt(timeStampStr, 10, 64)
	if err != nil {
		return false
	}

	interval2 := time.Now().UnixMilli() - timeStampInt
	isValid := interval2 >= rq.LimitValue
	if !isValid {
		log.Warnf("%s request too frequently!", rq.Mobile)
	}
	return isValid
}

func (m *TimeIntervalForMobileRule) getFailReason() error {
	return errors.New("sent_message_too_frequently")
}
