package checker

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/sms/store"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
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
	start := time.Now().Unix()
	key := factory.WrapperTimeInterval(rq.Mobile, rq.TemplateCode)
	rds := m.RDS

	// 使用INCR命令递增计数
	count, err := rds.Incr(ctx, key).Result()
	if err != nil {
		log.Errorf("Failed to increment count for key: %s, error: %v", key, err)
		return false
	}

	// 如果INCR返回1，说明是新创建的键，需要设置过期时间
	if count == 1 {
		// 这里我们设置一个过期时间，这个过期时间应该基于您的业务逻辑来确定
		// 例如，如果LimitLeftTime是允许再次发送消息的最小时间间隔（秒）
		// 那么这里可以设置为LimitLeftTime的一半，以确保有足够的时间间隔
		halfInterval := types.LimitLeftTime / 2
		if err := rds.SetNX(ctx, key, time.Now().Unix(), halfInterval*time.Second).Err(); err != nil {
			log.Errorf("Failed to set key with expiration for key: %s, error: %v", key, err)
			return false
		}
	}

	// 记录操作完成的时间
	log.Infof("TemplateAndMobileChecker-----checker time效验号码模板总时间----: %d", time.Now().Unix()-start)

	// 判断是否超过了时间间隔
	// 这里我们检查当前时间是否超过了上次时间戳加上时间间隔
	isValid := time.Now().Unix() > (count-1)*int64(types.LimitLeftTime)+time.Now().Unix()
	if !isValid {
		log.Infow(":warning:", "key", key, "count", count, "isValid", isValid)
	}
	return isValid
}

func (m *TimeIntervalForMobileRule) getFailReason() error {
	return errors.New("sent_message_too_frequently")
}
