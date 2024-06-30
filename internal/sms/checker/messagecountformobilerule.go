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
	"strconv"
	"time"
)

type MessageCountForMobileRule struct {
	DS  store.IStore
	RDS *redis.Client
}

var _ Rule = (*MessageCountForMobileRule)(nil)

func (m *MessageCountForMobileRule) IsValid(rq *types.Request) bool {

	start := time.Now().Unix()
	key := factory.WrapperMobileCount(rq.mobile, rq.templateCode)
	ctx := context.Background()
	// 查询redis
	rds := m.RDS
	sentCount, err := rds.Get(ctx, key).Result()
	if err != nil {
		filter := map[string]any{
			"mobile": rq.mobile,
			"status": "OK",
		}
		count, _, err := m.DS.Histories().List(ctx, rq.templateCode, meta.WithFilter(filter))
		if err != nil {
			sentCount = ""
		} else if strconv.FormatInt(count, 10) == "0" {
			sentCount = ""
		}
	}

	if sentCount == "" {
		rds.SetNX(ctx, key, 1, types.LimitLeftTime)
		log.Infof("TemplateAndMobileChecker-----checker time效验号码模板总时间----: %d", time.Now().Unix()-start)
		return true
	} else {
		sentCount, _ := strconv.ParseInt(sentCount, 10, 64)
		sentCount += 1
		ttl, _ := rds.TTL(ctx, key).Result()
		rds.Expire(ctx, key, ttl)
		log.Infof("TemplateAndMobileChecker-----checker time效验号码模板总时间----: %d", time.Now().Unix()-start)
		isValid := rq.limitValue > sentCount

		if !isValid {
			log.Infow(":warning:", "key", key, "sentCount", sentCount, "isValid", isValid)
		}
		return isValid
	}
}

func (m *MessageCountForMobileRule) GetFailReason() error {
	return errors.New("exceed_limit_for_this_phone")
}
