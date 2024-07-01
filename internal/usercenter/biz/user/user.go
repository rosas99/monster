package user

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/usercenter/model"
	"github.com/rosas99/monster/internal/usercenter/store"
	v1 "github.com/rosas99/monster/pkg/api/usercenter/v1"
	"github.com/rosas99/monster/pkg/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"sync"
)

type UserBiz interface {
	Create(ctx context.Context, rq *v1.CreateUserRequest) (*v1.CreateUserResponse, error)
	Get(ctx context.Context, rq *v1.GetUserRequest) (*v1.GetUserResponse, error)
	List(ctx context.Context, rq *v1.ListUserRequest) (*v1.ListUserResponse, error)
	Update(ctx context.Context, rq *v1.UpdateUserRequest) error
	Delete(ctx context.Context, rq *v1.DeleteUserRequest) error
}
type userBiz struct {
	ds  store.IStore
	rds *redis.Client
}

func New(ds store.IStore, rds *redis.Client) *userBiz {
	return &userBiz{ds: ds, rds: rds}
}

type MessageConfigurationEnum = string

func (t *userBiz) Create(ctx context.Context, rq *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {

	var userM model.UserM
	_ = copier.Copy(&userM, rq)
	if err := t.ds.Users().Create(ctx, &userM); err != nil {
		return nil, nil
	}

	return &v1.CreateUserResponse{UserID: userM.ID}, nil
}

func (t *userBiz) Get(ctx context.Context, rq *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	userM, err := t.ds.Users().Get(ctx, rq.GetUserName())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	var ret v1.GetUserResponse
	_ = copier.Copy(&ret, userM)
	ret.UserInfo.CreatedAt = timestamppb.New(userM.CreatedAt)
	ret.UserInfo.UpdatedAt = timestamppb.New(userM.UpdatedAt)

	return &ret, nil
}

func (t *userBiz) List(ctx context.Context, rq *v1.ListUserRequest) (*v1.ListUserResponse, error) {

	// todo 查询到template 转换为resp
	count, list, err := t.ds.Users().List(ctx, meta.WithOffset(rq.Offset), meta.WithLimit(rq.Limit))
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
				var t v1.UserInfo
				_ = copier.Copy(&t, template)
				m.Store(template.ID, &v1.UserInfo{
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
	users := make([]*v1.UserInfo, 0, len(list))
	for _, item := range list {
		// 从map加载组装后的数据
		template, _ := m.Load(item.ID)
		users = append(users, template.(*v1.UserInfo))
	}

	log.C(ctx).Debugw("Get orders from backend storage", "count", len(users))

	return &v1.ListUserResponse{TotalCount: count, Users: users}, nil
}

func (t *userBiz) Update(ctx context.Context, rq *v1.UpdateUserRequest) error {
	orderM, err := t.ds.Users().Get(ctx, rq.GetUserame())
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

	err = t.ds.Users().Update(ctx, orderM)

	return err
}

// Delete 是 OrderBiz 接口中 `Delete` 方法的实现.
func (t *userBiz) Delete(ctx context.Context, rq *v1.DeleteUserRequest) error {
	if err := t.ds.Users().Delete(ctx, rq.GetUserID()); err != nil {
		return err
	}

	return nil
}
