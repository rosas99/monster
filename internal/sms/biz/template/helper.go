package template

import (
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/sms/model"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	//"google.golang.org/protobuf/types/known/timestamppb"
)

// ModelToReply converts a model.UserM to a v1.UserReply. It copies the data from
// userM to user and sets the CreatedAt and UpdatedAt fields to their respective timestamps.
func ModelToReply(userM *model.TemplateM) *v1.CreateTemplateResponse {
	var user v1.CreateTemplateResponse
	_ = copier.Copy(&user, userM)
	//user.CreatedAt = timestamppb.New(userM.CreatedAt)
	//user.UpdatedAt = timestamppb.New(userM.UpdatedAt)
	return &user
}
