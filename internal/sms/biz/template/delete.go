package template

import (
	"context"
)

// Delete deletes a template from the database.
func (t *templateBiz) Delete(ctx context.Context, id int64) error {
	filters := map[string]any{"id": id}
	if err := t.ds.Templates().Delete(ctx, filters); err != nil {
		return err
	}

	return nil
}
