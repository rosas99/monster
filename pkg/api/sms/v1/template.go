package v1

import (
	"regexp"
)

type CreateTemplateRequest struct {
	TemplateName  string `json:"phoneNumber" valid:"required,stringlength(1|255)"`
	Content       string `json:"content" valid:"required,stringlength(1|255)"`
	TemplateType  string `json:"templateType" valid:"required,stringlength(1|255)"`
	Brand         string `json:"brand" valid:"required,stringlength(1|255)"`
	Providers     string `json:"providers" valid:"required,stringlength(1|255)"`
	TokenId       string `json:"tokenId" valid:"required,stringlength(1|255)"`
	TemplateCode  string `json:"templateCode" valid:"required,stringlength(1|255)"`
	Sign          string `json:"sign" valid:"required,stringlength(1|255)"`
	UserId        string `json:"userId" valid:"required,stringlength(1|255)"`
	TemplateCount string `json:"templateCount" valid:"required,gt=0,stringlength(1|255)"`
	MobileCount   string `json:"mobileCount" valid:"required,gt=0,stringlength(1|255)"`
	TimeInterval  string `json:"timeInterval" valid:"required,gt=0,stringlength(1|255)"`
	Region        string `json:"region" valid:"required,stringlength(1|255)"`
	Mobile        string `json:"mobile" valid:"required,isMobile"`
	Code          string `json:"code" valid:"required,stringlength(1|255)"`
}

var mobileRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

type CreateTemplateResponse struct {
	OrderID string `json:"orderID"`
}

type UpdateTemplateRequest struct {
	TemplateName  *string `json:"phoneNumber" valid:"required,stringlength(1|255)"`
	Content       *string `json:"content" valid:"required,stringlength(1|255)"`
	Type          *string `json:"templateType" valid:"required,stringlength(1|255)"`
	Brand         *string `json:"brand" valid:"required,stringlength(1|255)"`
	Providers     *string `json:"providers" valid:"required,stringlength(1|255)"`
	TokenId       *string `json:"tokenId" valid:"required,stringlength(1|255)"`
	TemplateCode  *string `json:"templateCode" valid:"required,stringlength(1|255)"`
	Sign          *string `json:"sign" valid:"required,stringlength(1|255)"`
	UserId        *string `json:"userId" valid:"required,stringlength(1|255)"`
	TemplateCount *string `json:"templateCount" valid:"required,gt=0,stringlength(1|255)"`
	MobileCount   *string `json:"mobileCount" valid:"required,gt=0,stringlength(1|255)"`
	TimeInterval  *string `json:"timeInterval" valid:"required,gt=0,stringlength(1|255)"`
	Region        *string `json:"region" valid:"required,stringlength(1|255)"`
	Mobile        *string `json:"mobile" valid:"required,isMobile"`
	Code          *string `json:"code" valid:"required,stringlength(1|255)"`
}

type TemplateReply struct {
	TemplateName  string `json:"phoneNumber"`
	Content       string `json:"content"`
	TemplateType  string `json:"templateType"`
	Brand         string `json:"brand"`
	Providers     string `json:"providers"`
	TokenId       string `json:"tokenId"`
	TemplateCode  string `json:"templateCode"`
	Sign          string `json:"sign"`
	UserId        string `json:"userId"`
	TemplateCount string `json:"templateCount"`
	MobileCount   string `json:"mobileCount"`
	TimeInterval  string `json:"timeInterval"`
	Region        string `json:"region"`
	Mobile        string `json:"mobile"`
	Code          string `json:"code"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

type ListTemplateRequest struct {
	Limit        int64  `json:"limit" valid:"required,stringlength(1|255)"`
	Offset       int64  `json:"offset" valid:"required,stringlength(1|255)"`
	TemplateCode string `json:"templateCode" valid:"required,stringlength(1|255)"`
	ExtCode      string `json:"extCode" valid:"required,stringlength(1|255)"`
}

type ListTemplateResponse struct {
	TotalCount int64            `json:"totalCount"`
	Templates  []*TemplateReply `json:"templates"`
}
type GetTemplateRequest struct {
	ID string `json:"id" valid:"required,stringlength(1|255)"`
}
type DeleteTemplateRequest struct {
	ID int64 `json:"id" valid:"required,stringlength(1|255)"`
}
