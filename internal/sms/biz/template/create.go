package template

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/types"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// Create creates a new template and stores it in the database.
func (t *templateBiz) Create(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error) {

	var templateM model.TemplateM
	_ = copier.Copy(&templateM, rq)
	if err := t.ds.Templates().Create(ctx, &templateM); err != nil {
		return nil, err
	}

	configurationsM := []*model.ConfigurationM{
		{
			ConfigKey:    types.MessageCountForMobilePerDay,
			ConfigValue:  rq.MobileCount,
			TemplateCode: rq.TemplateCode,
			Order:        1,
		},
		{
			ConfigKey:    types.TimeIntervalForMobilePerDay,
			ConfigValue:  rq.TimeInterval,
			TemplateCode: rq.TemplateCode,
			Order:        2,
		},
		{
			ConfigKey:    types.MessageCountForTemplatePerDay,
			ConfigValue:  rq.TemplateCount,
			TemplateCode: rq.TemplateCode,
			Order:        3,
		}}

	if err := t.ds.Configurations().CreateBatch(ctx, configurationsM); err != nil {
		return nil, err
	}

	return &v1.CreateTemplateResponse{OrderID: templateM.TemplateCode}, nil
}
