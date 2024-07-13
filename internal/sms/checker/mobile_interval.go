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
	rds := m.RDS
	intervalTime := rq.LimitValue

	timeStampStr, err := rds.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Errorf("Failed to get timestamp from Redis for key: %s, error: %v", key, err)
		return false
	}
	if timeStampStr == "" {
		if _, err := rds.Set(ctx, key, start, types.LimitLeftTime).Result(); err != nil {
			log.Errorf("Failed to set key with expiration for key: %s, error: %v", key, err)
			return false
		}
		log.Infof("timeInterval --- checker time -----效验时间戳的总时间: %d", time.Now().UnixMilli()-start)
		return true
	} else {
		remainingTTL, _ := rds.TTL(ctx, key).Result()
		if _, err := rds.Set(ctx, key, start, remainingTTL*time.Second).Result(); err != nil {
			log.Errorf("Failed to update key with new timestamp for key: %s, error: %v", key, err)
			return false
		}
		log.Infof("timeInterval --- checker time -----效验时间戳的总时间: %d", time.Now().UnixMilli()-start)
	}

	timeStampInt, _ := strconv.ParseInt(timeStampStr, 10, 64)
	interval := time.Now().UnixMilli() - timeStampInt
	isValid := interval >= intervalTime*1000
	if !isValid {
		log.Warnf("%s request too frequently!", rq.Mobile)
	}
	return isValid
}

func (m *TimeIntervalForMobileRule) getFailReason() error {
	return errors.New("sent_message_too_frequently")
}
