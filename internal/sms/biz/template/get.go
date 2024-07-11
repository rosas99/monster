package template

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// Get retrieves a single template from the database.
func (t *templateBiz) Get(ctx context.Context, rq *v1.GetTemplateRequest) (*v1.TemplateReply, error) {
	templateM, err := t.ds.Templates().Get(ctx, rq.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrorOrderNotFound(err.Error())
		}

		return nil, err
	}

	var template v1.TemplateReply
	_ = copier.Copy(&template, templateM)
	template.CreatedAt = timestamppb.New(templateM.CreatedAt)
	template.UpdatedAt = timestamppb.New(templateM.UpdatedAt)

	_, cfgList, err := t.ds.Configurations().List(ctx, templateM.TemplateCode)
	for _, m := range cfgList {
		switch m.ConfigKey {
		case TimeIntervalForMobile:
			template.MobileCount = m.ConfigValue
			fallthrough
		case MessageCountForMobilePerDay:
			template.TemplateCount = m.ConfigValue
			fallthrough
		case MessageCountForTemplatePerDay:
			template.TimeInterval = m.ConfigValue
		}
	}
	return &template, nil
}
