// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package validate

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rosas99/monster/internal/pkg/known"
	"github.com/rosas99/monster/internal/sms/monitor"
	"github.com/rosas99/monster/internal/sms/store"
	"github.com/rosas99/monster/pkg/log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// Validation make sure users have the right resource permission and operation.
func Validation(ds store.IStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.FullPath() {
		case "/v1/template":
			if c.Request.Method == http.MethodGet {
				start := time.Now().UnixMilli()
				id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
				filters := map[string]any{"id": id}
				_, err := ds.Templates().Fetch(context.Background(), filters)
				if err != nil {
					log.C(c).Errorf("Failed to get template by ID: %d. Error: %v", id, err)
					monitor.GetMonitor().LogTemplateKpi("template", c.Request.Header.Get(known.TraceIDKey),
						false, time.Now().UnixMilli()-start)
					c.Abort()
					return
				}
				log.C(c).Infof("Successfully validated template ID: %d", id)
			}
		default:
			log.C(c).Infof("No validation required for path: %s", c.FullPath())
		}

		c.Next()
	}
}

func isMobileNo(mobiles string) bool {
	pattern := `^[0-9]{6}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(mobiles)
}
