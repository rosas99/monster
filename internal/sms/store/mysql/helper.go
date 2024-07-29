package mysql

import "github.com/rosas99/monster/internal/sms/model"

type ByOrder []*model.ConfigurationM

func (o ByOrder) Len() int           { return len(o) }
func (o ByOrder) Less(i, j int) bool { return o[i].Order < o[j].Order }
func (o ByOrder) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }

// ModelToReply converts a model.UserM to a v1.UserReply. It copies the data from
// userM to user and sets the CreatedAt and UpdatedAt fields to their respective timestamps.
//func ModelToReply(userM *model.UserM) *v1.UserReply {
//	var user v1.UserReply
//	_ = copier.Copy(&user, userM)
//	user.CreatedAt = timestamppb.New(userM.CreatedAt)
//	user.UpdatedAt = timestamppb.New(userM.UpdatedAt)
//	return &user
//}
