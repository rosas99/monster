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

var _ Rule = (*TimeIntervalForMobileRule)(nil)

func (m *TimeIntervalForMobileRule) isValid(rq *types.Request) bool {
	start := time.Now().Unix()
	// todo 修改为store
	key := factory.WrapperTimeInterval(rq.Mobile, rq.TemplateCode)
	ctx := context.Background()
	rds := m.RDS
	timeStampStr, err := rds.Get(ctx, key).Result()
	if err != nil {
		timeStampStr = ""
	}

	if timeStampStr == "" {
		rds.SetNX(ctx, key, 1, types.LimitLeftTime)
		log.Infof("TemplateAndMobileChecker-----checker time效验号码模板总时间----: %d", time.Now().Unix()-start)
		return true
	} else {
		ttl, _ := rds.TTL(ctx, key).Result()
		rds.Expire(ctx, key, ttl)
		log.Infof("TemplateAndMobileChecker-----checker time效验号码模板总时间----: %d", time.Now().Unix()-start)
		timeStampInt, _ := strconv.ParseInt(timeStampStr, 10, 64)
		isValid := time.Now().Unix() > timeStampInt
		if !isValid {
			log.Infow(":warning:", "key", key, "timeStampInt", timeStampInt, "isValid", isValid)
		}
		return isValid
	}
}

func (m *TimeIntervalForMobileRule) getFailReason() error {
	return errors.New("sent_message_too_frequently")
}
