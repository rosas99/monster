package v1

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
	TemplateCount string `json:"templateCount" valid:"required,stringlength(1|255)"`
	MobileCount   string `json:"mobileCount" valid:"required,stringlength(1|255)"`
	TimeInterval  string `json:"timeInterval" valid:"required,stringlength(1|255)"`
	Region        string `json:"region" valid:"required,stringlength(1|255)"`
	Mobile        string `json:"mobile" valid:"required,stringlength(1|255)"`
	Code          string `json:"code" valid:"required,stringlength(1|255)"`
}

type CreateTemplateResponse struct {
	OrderID string `json:"orderID" valid:"required,stringlength(1|255)"`
}

type UpdateTemplateRequest struct {
	TemplateName  string `json:"phoneNumber" valid:"required,stringlength(1|255)"`
	Content       string `json:"content" valid:"required,stringlength(1|255)"`
	TemplateType  string `json:"templateType" valid:"required,stringlength(1|255)"`
	Brand         string `json:"brand" valid:"required,stringlength(1|255)"`
	Providers     string `json:"providers" valid:"required,stringlength(1|255)"`
	TokenId       string `json:"tokenId" valid:"required,stringlength(1|255)"`
	TemplateCode  string `json:"templateCode" valid:"required,stringlength(1|255)"`
	Sign          string `json:"sign" valid:"required,stringlength(1|255)"`
	UserId        string `json:"userId" valid:"required,stringlength(1|255)"`
	TemplateCount string `json:"templateCount" valid:"required,stringlength(1|255)"`
	MobileCount   string `json:"mobileCount" valid:"required,stringlength(1|255)"`
	TimeInterval  string `json:"timeInterval" valid:"required,stringlength(1|255)"`
	Region        string `json:"region" valid:"required,stringlength(1|255)"`
	Mobile        string `json:"mobile" valid:"required,stringlength(1|255)"`
	Code          string `json:"code" valid:"required,stringlength(1|255)"`
}

type TemplateReply struct {
	TemplateName  string `json:"phoneNumber" valid:"required,stringlength(1|255)"`
	Content       string `json:"content" valid:"required,stringlength(1|255)"`
	TemplateType  string `json:"templateType" valid:"required,stringlength(1|255)"`
	Brand         string `json:"brand" valid:"required,stringlength(1|255)"`
	Providers     string `json:"providers" valid:"required,stringlength(1|255)"`
	TokenId       string `json:"tokenId" valid:"required,stringlength(1|255)"`
	TemplateCode  string `json:"templateCode" valid:"required,stringlength(1|255)"`
	Sign          string `json:"sign" valid:"required,stringlength(1|255)"`
	UserId        string `json:"userId" valid:"required,stringlength(1|255)"`
	TemplateCount string `json:"templateCount" valid:"required,stringlength(1|255)"`
	MobileCount   string `json:"mobileCount" valid:"required,stringlength(1|255)"`
	TimeInterval  string `json:"timeInterval" valid:"required,stringlength(1|255)"`
	Region        string `json:"region" valid:"required,stringlength(1|255)"`
	Mobile        string `json:"mobile" valid:"required,stringlength(1|255)"`
	Code          string `json:"code" valid:"required,stringlength(1|255)"`
	CreatedAt     string `json:"createdAt" valid:"required,stringlength(1|255)"`
	UpdatedAt     string `json:"updatedAt" valid:"required,stringlength(1|255)"`
}

type ListTemplateRequest struct {
	Limit        int64  `json:"limit" valid:"required,stringlength(1|255)"`
	Offset       int64  `json:"offset" valid:"required,stringlength(1|255)"`
	TemplateCode string `json:"templateCode" valid:"required,stringlength(1|255)"`
	ExtCode      string `json:"extCode" valid:"required,stringlength(1|255)"`
}

type ListTemplateResponse struct {
}
type GetTemplateRequest struct {
	ID string `json:"id" valid:"required,stringlength(1|255)"`
}
type DeleteTemplateRequest struct {
	ID string `json:"id" valid:"required,stringlength(1|255)"`
}
