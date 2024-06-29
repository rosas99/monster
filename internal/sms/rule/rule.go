package rule

import (
	"errors"
	"fmt"
	"github.com/rosas99/monster/internal/sms/model"
)

// CheckerRequest 模拟验证请求
type CheckerRequest struct {
	// 根据需要添加字段
}

// CheckerResponse 模拟验证失败的原因
type CheckerResponse struct {
	Reason string
}

// 策略模式
// todo 参考iam注册方式
// Rule 接口定义了验证规则需要实现的方法
type Rule interface {
	IsValid(*CheckerRequest) bool
	GetFailReason() *CheckerResponse
}

type RuleFactory struct {
	rules map[string]Rule
	// todo  redis

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

func init() {
	factory := NewRuleFactory()
	factory.RegisterRule("MESSAGE_COUNT_FOR_TEMPLATE_PER_DAY", &MessageCountForTemplateRule{})
}

func (rf *RuleFactory) CheckRules(template *model.TemplateM, mobile string, cfgList []*model.ConfigurationM) error {
	if len(cfgList) == 0 {
		return errors.New("no configuration")
	}

	// todo 排序 枚举设置了值
	// 或者cfg初始化时定义序号
	for _, cfg := range cfgList {
		// 创建checker
		checker, err := rf.CreateChecker(cfg)
		if err != nil {

		}
		var c CheckerRequest
		if !checker.IsValid(&c) {
			// 这里fail reason可以使用生成的错误码替换
			checker.GetFailReason()
			return errors.New("no configuration")

		}

	}
	return nil

}

// CreateChecker 根据 CheckerRequest 创建对应的 Rule
func (rf *RuleFactory) CreateChecker(cfg *model.ConfigurationM) (Rule, error) {
	// 假设 CheckerRequest 中有一个字段用于确定使用哪个 Rule
	// 这里只是一个示例，您需要根据实际情况实现具体的逻辑
	checkType := cfg.ConfigKey
	// todo 对应实现也要引入redis
	rule, exists := rf.rules[checkType]
	if !exists {
		return nil, fmt.Errorf("invalid check type: %s", checkType)
	}

	return rule, nil
}
