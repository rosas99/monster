package template

import (
	"context"
	"github.com/rosas99/monster/internal/sms/model"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// MessageConfigurationEnum defines an enumerated type for different message configuration.
type MessageConfigurationEnum = string

// defines a group of constants for message configuration.

const (
	TimeIntervalForMobile         MessageConfigurationEnum = "TIME_INTERVAL_FOR_MOBILE"
	MessageCountForMobilePerDay   MessageConfigurationEnum = "MESSAGE_COUNT_FOR_MOBILE_PER_DAY"
	MessageCountForTemplatePerDay MessageConfigurationEnum = "MESSAGE_COUNT_FOR_TEMPLATE_PER_DAY"
)

// Update updates a template's information in the database.
func (t *templateBiz) Update(ctx context.Context, rq *v1.UpdateTemplateRequest) error {
	orderM, err := t.ds.Templates().Get(ctx, rq.TemplateCode)
	if err != nil {
		return err
	}

	//不为空才更新，类型Mybatis的动态sql
	//if rq.Customer != nil {
	//	orderM.Customer = *rq.Customer
	//}
	//
	//if rq.Product != nil {
	//	orderM.Product = *rq.Product
	//}
	//
	//if rq.Quantity != nil {
	//	orderM.Quantity = *rq.Quantity
	//}

	err = t.ds.Templates().Update(ctx, orderM)

	configurationsM := []*model.ConfigurationM{
		{
			ConfigKey:    MessageCountForMobilePerDay,
			ConfigValue:  rq.GetMobileCount(),
			TemplateCode: rq.GetTemplateCode(),
		},
		{
			ConfigKey:    MessageCountForTemplatePerDay,
			ConfigValue:  rq.GetTemplateCount(),
			TemplateCode: rq.GetTemplateCode(),
		},
		{
			ConfigKey:    TimeIntervalForMobile,
			ConfigValue:  rq.GetTimeInterval(),
			TemplateCode: rq.GetTemplateCode(),
		}}
	if err := t.ds.Configurations().CreateBatch(ctx, configurationsM); err != nil {
		// todo 错误码定义
		return v1.ErrorOrderCreateFailed("create order failed: %v", err)
	}
	return err
}
