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
	"github.com/rosas99/monster/pkg/auth"
	"github.com/rosas99/monster/pkg/log"
	"github.com/rosas99/monster/pkg/token"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"regexp"
	"sync"
)

type UserBiz interface {
	ChangePassword(ctx context.Context, r *v1.ChangePasswordRequest) error
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
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

// ChangePassword 是 UserBiz 接口中 `ChangePassword` 方法的实现.
func (b *userBiz) ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error {
	userM, err := b.ds.Users().Get(ctx, username)
	if err != nil {
		return err
	}

	if err := auth.Compare(userM.Password, r.OldPassword); err != nil {
		log.C(ctx).Errorw(err, "Failed to list orders from storage")

		return errors.New("old password is invalid")
	}

	userM.Password, _ = auth.Encrypt(r.NewPassword)
	if err := b.ds.Users().Update(ctx, userM); err != nil {
		return err
	}

	return nil
}

// Login 是 UserBiz 接口中 `Login` 方法的实现.
func (b *userBiz) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	// 获取登录用户的所有信息
	user, err := b.ds.Users().Get(ctx, r.Username)
	if err != nil {
		return nil, errors.New("old password is invalid")

	}

	// 对比传入的明文密码和数据库中已加密过的密码是否匹配
	if err := auth.Compare(user.Password, r.Password); err != nil {
		return nil, errors.New("old password is invalid")
	}

	// 如果匹配成功，说明登录成功，签发 token 并返回
	t, err := token.Sign(r.Username)
	if err != nil {
		return nil, errors.New("old password is invalid")
	}

	return &v1.LoginResponse{Token: t}, nil
}

func (t *userBiz) Create(ctx context.Context, rq *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {

	var userM model.UserM
	_ = copier.Copy(&userM, rq)
	if err := t.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return nil, errors.New("old password is invalid")
		}
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

	var resp v1.GetUserResponse
	_ = copier.Copy(&resp, userM)
	resp.UserInfo.CreatedAt = timestamppb.New(userM.CreatedAt)
	resp.UserInfo.UpdatedAt = timestamppb.New(userM.UpdatedAt)

	return &resp, nil
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
	userM, err := t.ds.Users().Get(ctx, rq.GetUsername())
	if err != nil {
		return err
	}

	if rq.Email != nil {
		userM.Email = *rq.Email
	}

	if rq.Nickname != nil {
		userM.Nickname = *rq.Nickname
	}

	if rq.Phone != nil {
		userM.Phone = *rq.Phone
	}

	err = t.ds.Users().Update(ctx, userM)

	return err
}

// Delete 是 OrderBiz 接口中 `Delete` 方法的实现.
func (t *userBiz) Delete(ctx context.Context, rq *v1.DeleteUserRequest) error {
	if err := t.ds.Users().Delete(ctx, rq.GetUserID()); err != nil {
		return err
	}

	return nil
}
