// Package historyclean is a watcher implement used to delete expired record from the database.
package historyclean

import (
	"context"
	"github.com/rosas99/monster/internal/nightwatch/watcher"
	"github.com/rosas99/monster/internal/pkg/client/store"
	"github.com/rosas99/monster/pkg/log"
	//"sigs.k8s.io/cluster-api/util/secret"
	"time"
)

var _ watcher.Watcher = (*historiesCleanWatcher)(nil)

// watcher implement.
type historiesCleanWatcher struct {
	store store.Interface
}

// Run runs the watcher.
func (w *historiesCleanWatcher) Run() {
	_, histories, err := w.store.Sms().Histories().List(context.Background(), "")
	if err != nil {
		log.Errorw(err, "Failed to list secrets")
		return
	}

	for _, history := range histories {
		// deletes all records that are older than one year.
		if history.CreatedAt.Unix() < time.Now().AddDate(-1, 0, 0).Unix() {
			err := w.store.Sms().Histories().Delete(context.TODO(), history.ID)
			if err != nil {
				log.Warnw("Failed to delete secret from database", "userID", history.ID, "name", "secret.Name")
				continue
			}
			log.Infow("Successfully deleted secret from database", "userID", history.ID, "name", "secret.Name")
		}
	}
}

// Init initializes the watcher for later execution.
func (w *historiesCleanWatcher) Init(ctx context.Context, config *watcher.Config) error {
	w.store = config.Store
	return nil
}

func init() {
	watcher.Register("historyclean", &historiesCleanWatcher{})
}
