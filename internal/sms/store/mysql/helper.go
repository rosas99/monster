package mysql

import "github.com/rosas99/monster/internal/sms/model"

type ByOrder []*model.ConfigurationM

func (o ByOrder) Len() int           { return len(o) }
func (o ByOrder) Less(i, j int) bool { return o[i].Order < o[j].Order }
func (o ByOrder) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
