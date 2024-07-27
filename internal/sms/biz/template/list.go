package template

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/pkg/meta"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
	"golang.org/x/sync/errgroup"
	"sync"
)

// List retrieves a list of all templates from the database.
func (t *templateBiz) List(ctx context.Context, rq *v1.ListTemplateRequest) (*v1.ListTemplateResponse, error) {

	count, list, err := t.ds.Templates().List(ctx, rq.TemplateCode, meta.WithOffset(rq.Offset), meta.WithLimit(rq.Limit))
	if err != nil {
		log.C(ctx).Errorw(err, "Failed to list orders from storage")
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)
	for _, template := range list {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				var t v1.TemplateReply
				_ = copier.Copy(&t, template)
				m.Store(template.ID, &v1.TemplateReply{
					CreatedAt: template.CreatedAt.Format("2006-01-02 15:04:05"),
					UpdatedAt: template.UpdatedAt.Format("2006-01-02 15:04:05"),
				})
				return nil
			}

		})
	}

	if err := eg.Wait(); err != nil {
		log.C(ctx).Errorw(err, "Failed to wait all function calls returned")
		return nil, err
	}

	// The following code block is used to maintain the consistency of query order.
	templates := make([]*v1.TemplateReply, 0, len(list))
	for _, item := range list {
		template, _ := m.Load(item.ID)
		templates = append(templates, template.(*v1.TemplateReply))
	}

	log.C(ctx).Debugw("Get orders from backend storage", "count", len(templates))

	return &v1.ListTemplateResponse{TotalCount: count, Templates: templates}, nil
}

func (t *templateBiz) ListWithBadPerformance(ctx context.Context, rq *v1.ListTemplateRequest) (*v1.ListTemplateResponse, error) {

	count, list, err := t.ds.Templates().List(ctx, rq.TemplateCode, meta.WithOffset(rq.Offset), meta.WithLimit(rq.Limit))
	if err != nil {
		log.C(ctx).Errorw(err, "Failed to list orders from storage")
		return nil, err
	}

	templates := make([]*v1.TemplateReply, 0, len(list))

	for _, item := range list {
		var t v1.TemplateReply
		_ = copier.Copy(&t, item)
		templates = append(templates, &v1.TemplateReply{
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: item.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	log.C(ctx).Debugw("Get orders from backend storage", "count", len(templates))

	return &v1.ListTemplateResponse{TotalCount: count, Templates: templates}, nil
}
