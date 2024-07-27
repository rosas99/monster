package template

import (
	"context"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/types"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// Update updates a template's information in the database.
func (t *templateBiz) Update(ctx context.Context, id int64, rq *v1.UpdateTemplateRequest) error {
	var err error
	orderM, err := t.ds.Templates().Get(ctx, id)
	if err != nil {
		return err
	}

	//if rq.Brand != nil {
	//	orderM.Brand = *rq.Brand
	//}

	if rq.Sign != nil {
		orderM.Sign = *rq.Sign
	}

	if rq.Content != nil {
		orderM.Content = *rq.Content
	}

	if rq.Type != nil {
		orderM.Type = *rq.Type
	}

	if rq.TemplateName != nil {
		orderM.TemplateName = *rq.TemplateName
	}

	if rq.Region != nil {
		orderM.Region = *rq.Region
	}

	if rq.Providers != nil {
		orderM.Providers = *rq.Providers
	}

	if err = t.ds.Templates().Update(ctx, orderM); err != nil {
		return err
	}

	configurationsM := []*model.ConfigurationM{
		{
			ConfigKey:    types.MessageCountForMobilePerDay,
			ConfigValue:  *rq.MobileCount,
			TemplateCode: *rq.TemplateCode,
		},
		{
			ConfigKey:    types.MessageCountForTemplatePerDay,
			ConfigValue:  *rq.TemplateCount,
			TemplateCode: *rq.TemplateCode,
		},
		{
			ConfigKey:    types.TimeIntervalForMobilePerDay,
			ConfigValue:  *rq.TimeInterval,
			TemplateCode: *rq.TemplateCode,
		}}

	err = t.ds.Configurations().CreateBatch(ctx, configurationsM)
	if err != nil {
		return err
	}

	return err
}
