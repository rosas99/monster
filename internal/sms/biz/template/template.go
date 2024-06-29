package template

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/store"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"sync"
)

type TemplateBiz interface {
	Create(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error)
	Get(ctx context.Context, rq *v1.GetTemplateRequest) (*v1.TemplateReply, error)
	List(ctx context.Context, rq *v1.ListTemplateRequest) (*v1.ListTemplateResponse, error)
	Update(ctx context.Context, rq *v1.UpdateTemplateRequest) error
	Delete(ctx context.Context, rq *v1.DeleteTemplateRequest) error
}
type templateBiz struct {
	ds  store.IStore
	rds *redis.Client
}

func New(ds store.IStore, rds *redis.Client) *templateBiz {
	return &templateBiz{ds: ds, rds: rds}
}

type MessageConfigurationEnum = string

const (
	TimeIntervalForMobile         MessageConfigurationEnum = "TIME_INTERVAL_FOR_MOBILE"           // Status used for disabling a secret.
	MessageCountForMobilePerDay   MessageConfigurationEnum = "MESSAGE_COUNT_FOR_MOBILE_PER_DAY"   // Status used for enabling a secret.
	MessageCountForTemplatePerDay MessageConfigurationEnum = "MESSAGE_COUNT_FOR_TEMPLATE_PER_DAY" // Status used for enabling a secret.
)

// todo 使用事务钩子
func (t *templateBiz) Create(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error) {
	// todo 增加场景码 进行台账统计

	// new 空对象
	var templateM model.TemplateM
	// 深拷贝请求参数
	_ = copier.Copy(&templateM, rq)
	if err := t.ds.Templates().Create(ctx, &templateM); err != nil {
		// todo 错误码定义
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
		// todo 错误码定义
		return nil, v1.ErrorOrderCreateFailed("create order failed: %v", err)
	}

	//todo 开始写消息
	//b.writer.WriteMessages()

	return &v1.CreateTemplateResponse{OrderID: templateM.ID}, nil
}

func (t *templateBiz) Get(ctx context.Context, rq *v1.GetTemplateRequest) (*v1.TemplateReply, error) {
	// todo 查询config 设置到template
	templateM, err := t.ds.Templates().Get(ctx, rq.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrorOrderNotFound(err.Error())
		}

		return nil, err
	}

	// model复制到resp 都是指针
	var template v1.TemplateReply
	_ = copier.Copy(&template, templateM)
	// 这里要把time转换成timestamp ？
	template.CreatedAt = timestamppb.New(templateM.CreateAt)
	template.UpdatedAt = timestamppb.New(templateM.UpdateAt)

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
					// 除了时间其他要手动set吗
					CreatedAt: timestamppb.New(template.CreateAt),
					UpdatedAt: timestamppb.New(template.UpdateAt),
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

func (t *templateBiz) Update(ctx context.Context, rq *v1.UpdateTemplateRequest) error {
	orderM, err := t.ds.Templates().Get(ctx, rq.TemplateCode)
	if err != nil {
		return err
	}

	// 不为空才更新，类型Mybatis的动态sql
	//if rq.Customer != nil {
	//	orderM.Customer = *rq.Customer
	//}
	//
	//if rq.Product != nil {
	//	orderM.Product = *rq.Product
	//}
	//
	//if rq.Quantity != nil {
	//	orderM.Quantity = *rq.Quantity
	//}

	err = t.ds.Templates().Update(ctx, orderM)

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
		// todo 错误码定义
		return v1.ErrorOrderCreateFailed("create order failed: %v", err)
	}
	return err
}

// Delete 是 OrderBiz 接口中 `Delete` 方法的实现.
func (t *templateBiz) Delete(ctx context.Context, rq *v1.DeleteTemplateRequest) error {
	if err := t.ds.Templates().Delete(ctx, rq.Id); err != nil {
		return err
	}

	return nil
}
