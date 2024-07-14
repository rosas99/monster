package template

import (
	"context"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/types"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
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
			ConfigKey:    types.MessageCountForMobilePerDay,
			ConfigValue:  rq.MobileCount,
			TemplateCode: rq.TemplateCode,
		},
		{
			ConfigKey:    types.MessageCountForTemplatePerDay,
			ConfigValue:  rq.TemplateCount,
			TemplateCode: rq.TemplateCode,
		},
		{
			ConfigKey:    types.TimeIntervalForMobilePerDay,
			ConfigValue:  rq.TimeInterval,
			TemplateCode: rq.TemplateCode,
		}}
	if err := t.ds.Configurations().CreateBatch(ctx, configurationsM); err != nil {
		// todo 错误码定义
		return err
	}
	return err
}
