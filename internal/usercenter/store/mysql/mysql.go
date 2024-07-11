package mysql

import (
	"sync"

	"gorm.io/gorm"

	"github.com/rosas99/monster/internal/usercenter/store"
)

var (
	once sync.Once
	// 全局变量，保存已被初始化的 *Datastore 实例.
	S *Datastore
)

// Datastore 是 IStore 的一个具体实现.
type Datastore struct {
	db *gorm.DB
}

// 确保 Datastore 实现了 store.IStore 接口.
var _ store.IStore = (*Datastore)(nil)

// NewStore 创建一个 store.IStore 类型的实例.
func NewStore(db *gorm.DB) *Datastore {
	// 确保 s 只被初始化一次
	once.Do(func() {
		S = &Datastore{db}
	})

	return S
}

// DB 返回存储在 datastore 中的 *gorm.DB.
func (ds *Datastore) DB() *gorm.DB {
	return ds.db
}

func (ds *Datastore) Users() store.UserStore {
	return newUsers(ds.db)
}

// todo history
