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

type Rule interface {
	isValid(ctx context.Context, rq *types.Request) error
}

type RuleFactory struct {
	rules map[string]Rule
}

// NewRuleFactory 构造函数，初始化 RuleFactory 实例
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
			return err

		}

		var c types.Request
		err = checker.isValid(ctx, &c)
		if err != nil {
			return err
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
