// Package watcher provides functions used by all watchers.
package watcher

import (
	"github.com/rosas99/monster/internal/pkg/client/store"
	//clientset "github.com/rosas99/monster/pkg/generated/clientset/versioned"
)

// Config aggregates the configurations of all watchers and serves as a configuration aggregator.
type Config struct {
	// The purpose of nightwatch is to handle asynchronous tasks on the onex platform
	// in a unified manner, so a store aggregation type is needed here.
	Store store.Interface

	// Client is the client for onex-apiserver.
	//Client clientset.Interface

	// Then maximum concurrency event of user watcher.
	UserWatcherMaxWorkers int64
}
