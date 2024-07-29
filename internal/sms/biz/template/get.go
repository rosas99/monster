package template

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/pkg/errno"
	"github.com/rosas99/monster/internal/sms/types"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"gorm.io/gorm"
)

// Get retrieves a single template from the database.
func (t *templateBiz) Get(ctx context.Context, id int64) (*v1.TemplateReply, error) {
	filters := map[string]any{"id": id}
	templateM, err := t.ds.Templates().Fetch(ctx, filters)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrPageNotFound
		}

		return nil, err
	}

	var template v1.TemplateReply
	_ = copier.Copy(&template, templateM)
	template.CreatedAt = templateM.CreatedAt.Format("2006-01-02 15:04:05")
	template.UpdatedAt = templateM.UpdatedAt.Format("2006-01-02 15:04:05")

	_, cfgList, err := t.ds.Configurations().List(ctx, templateM.TemplateCode)
	for _, m := range cfgList {
		switch m.ConfigKey {
		case types.TimeIntervalForMobilePerDay:
			template.MobileCount = m.ConfigValue
			fallthrough
		case types.MessageCountForMobilePerDay:
			template.TemplateCount = m.ConfigValue
			fallthrough
		case types.MessageCountForTemplatePerDay:
			template.TimeInterval = m.ConfigValue
		}
	}
	return &template, nil
}
