package template

import (
	"context"
)

// Delete deletes a template from the database.
func (t *templateBiz) Delete(ctx context.Context, id int64) error {
	if err := t.ds.Templates().Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
