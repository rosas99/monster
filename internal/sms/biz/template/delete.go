package template

import (
	"context"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// Delete deletes a template from the database.
func (t *templateBiz) Delete(ctx context.Context, rq *v1.DeleteTemplateRequest) error {
	if err := t.ds.Templates().Delete(ctx, rq.Id); err != nil {
		return err
	}

	return nil
}
