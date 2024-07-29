package checker

import (
	"context"
	"errors"
	"fmt"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/pkg/log"
	"strconv"
)

//todo 补充注释

// Request  模拟验证请求
type Request struct {
	Mobile       string
	Id           int64
	TemplateCode string
	LimitValue   int64
}

type Rule interface {
	isValid(ctx context.Context, rq *Request) error
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

	for _, cfg := range cfgList {
		checker, err := rf.CreateChecker(cfg)
		if err != nil {
			log.C(ctx).Errorw(err, "Failed to list orders from storage")
			return err

		}

		val, err := strconv.ParseInt(cfg.ConfigValue, 10, 64)
		request := &Request{
			Mobile:       cfg.TemplateCode,
			TemplateCode: cfg.TemplateCode,
			LimitValue:   val,
		}
		err = checker.isValid(ctx, request)
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
