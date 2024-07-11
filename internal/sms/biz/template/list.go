package template

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/pkg/meta"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync"
)

// List retrieves a list of all templates from the database.
func (t *templateBiz) List(ctx context.Context, rq *v1.ListTemplateRequest) (*v1.ListTemplateResponse, error) {

	// todo 查询到template 转换为resp
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
					CreatedAt: timestamppb.New(template.CreatedAt),
					UpdatedAt: timestamppb.New(template.UpdatedAt),
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
		// 从map加载组装后的数据
		template, _ := m.Load(item.ID)
		templates = append(templates, template.(*v1.TemplateReply))
	}

	log.C(ctx).Debugw("Get orders from backend storage", "count", len(templates))

	return &v1.ListTemplateResponse{TotalCount: count, Templates: templates}, nil
}
