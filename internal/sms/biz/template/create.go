package template

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/sms/model"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// Create creates a new template and stores it in the database.
func (t *templateBiz) Create(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error) {

	var templateM model.TemplateM
	_ = copier.Copy(&templateM, rq)
	if err := t.ds.Templates().Create(ctx, &templateM); err != nil {
		return nil, v1.ErrorOrderCreateFailed("create order failed: %v", err)
	}

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
		return nil, v1.ErrorOrderCreateFailed("create order failed: %v", err)
	}

	return &v1.CreateTemplateResponse{OrderID: templateM.ID}, nil
}
