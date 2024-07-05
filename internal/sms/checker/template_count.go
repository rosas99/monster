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

type MessageCountForTemplateRule struct {
	DS  store.IStore
	RDS *redis.Client
}

func NewMessageCountForTemplateRule(DS store.IStore, RDS *redis.Client) *MessageCountForTemplateRule {
	return &MessageCountForTemplateRule{DS: DS, RDS: RDS}
}

var _ Rule = (*MessageCountForTemplateRule)(nil)

func (m *MessageCountForTemplateRule) isValid(rq *types.Request) bool {

	start := time.Now().Unix()
	key := factory.WrapperTemplateCount(rq.Mobile, rq.TemplateCode)
	ctx := context.Background()
	// 查询redis
	rds := m.RDS
	sentCount, err := rds.Get(ctx, key).Result()
	if err != nil {
		sentCount = ""
	}

	if sentCount == "" {
		rds.Set(ctx, key, time.Now().Unix(), types.LimitLeftTime)
		log.Infof("TemplateAndMobileChecker-----checker time效验号码模板总时间----: %d", time.Now().Unix()-start)
		return true
	} else {
		sentCount, _ := strconv.ParseInt(sentCount, 10, 64)
		sentCount += 1
		ttl, _ := rds.TTL(ctx, key).Result()
		rds.Expire(ctx, key, ttl)
		log.Infof("TemplateAndMobileChecker-----checker time效验号码模板总时间----: %d", time.Now().Unix()-start)
		isValid := rq.LimitValue > sentCount
		if !isValid {
			log.Infow(":warning:", "key", key, "sentCount", sentCount, "isValid", isValid)
		}
		return isValid
	}
}

func (m *MessageCountForTemplateRule) getFailReason() error {
	return errors.New("exceed_limit_for_this_template")
}
