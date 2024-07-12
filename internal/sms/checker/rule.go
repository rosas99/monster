package checker

import (
	"context"
	"errors"
	"fmt"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
	"sort"
)

// 策略模式
// todo 参考iam注册方式
// Rule 接口定义了验证规则需要实现的方法
type Rule interface {
	isValid(ctx context.Context, rq *types.Request) bool
	getFailReason() error
}

type RuleFactory struct {
	rules map[string]Rule
}

// NewRuleFactory 构造函数，初始化 RuleFactory 实例
// todo 注册到biz 中
func NewRuleFactory() *RuleFactory {
	return &RuleFactory{
		rules: make(map[string]Rule),
	}
}

// RegisterRule 注册 Rule 实现
func (rf *RuleFactory) RegisterRule(key string, rule Rule) {
	rf.rules[key] = rule
}

func (rf *RuleFactory) CheckRules(ctx context.Context, cfgList []*model.ConfigurationM) error {
	if len(cfgList) == 0 {
		return errors.New("no configuration")
	}

	// 升序排序
	sort.SliceStable(cfgList, func(i, j int) bool {
		return cfgList[i].Order < cfgList[j].Order
	})

	for _, cfg := range cfgList {
		checker, err := rf.CreateChecker(cfg)
		if err != nil {
			// todo  log
			log.C(ctx).Errorw(err, "Failed to list orders from storage")

		}
		var c types.Request
		if !checker.isValid(ctx, &c) {
			return checker.getFailReason()
		}
	}

	return nil

}

// CreateChecker 根据 CheckerRequest 创建对应的 Rule
func (rf *RuleFactory) CreateChecker(cfg *model.ConfigurationM) (Rule, error) {
	checkType := cfg.ConfigKey
	rule, exists := rf.rules[checkType]
	if !exists {
		return nil, fmt.Errorf("invalid check type: %s", checkType)
	}

	return rule, nil
}
