package v1

type SendMessageRequest struct {
	TemplateCode string `json:"templateCode" valid:"alphanum,required,stringlength(1|255)"`
	Mobile       string `json:"mobile" valid:"required,stringlength(1|255)"`
	Brand        string `json:"brand" valid:"required,stringlength(1|255)"`
	Code         string `json:"code" valid:"required,stringlength(1|255)"`
}

//type CommonResponse struct {
//	Code int64  `json:"code" valid:"required,stringlength(1|255)"`
//	Msg  string `json:"msg" valid:"required,stringlength(1|255)"`
//}

type VerifyCodeRequest struct {
	TemplateCode string `json:"templateCode" valid:"alphanum,required,stringlength(1|255)"`
	Mobile       string `json:"mobile" valid:"required,stringlength(1|255)"`
	Brand        string `json:"brand" valid:"required,stringlength(1|255)"`
	Code         string `json:"code" valid:"required,stringlength(1|255)"`
}

type AILIYUNReportListRequest struct {
	AILIYUNReportList []AILIYUNReport `json:"msg" valid:"required,stringlength(1|255)"`
}

type AILIYUNCallbackListRequest struct {
	AILIYUNCallbackList []AILIYUNCallbackList `json:"msg" valid:"required,stringlength(1|255)"`
}

type AILIYUNCallbackList struct {
	PhoneNumber string `json:"phoneNumber" valid:"required,stringlength(1|255)"`
	SendTime    string `json:"sendTime" valid:"required,stringlength(1|255)"`
	Content     string `json:"content" valid:"required,stringlength(1|255)"`
	SignName    string `json:"signName" valid:"required,stringlength(1|255)"`
	DestCode    string `json:"destCode" valid:"required,stringlength(1|255)"`
	SequenceId  string `json:"sequenceId" valid:"required,stringlength(1|255)"`
	RequestId   string `json:"requestId" valid:"required,stringlength(1|255)"`
}

type AILIYUNReport struct {
	PhoneNumber string `json:"phoneNumber" valid:"alphanum,required,stringlength(1|255)"`
	SendTime    string `json:"sendTime-time" valid:"required,stringlength(1|255)"`
	ReportTime  string `json:"reportTime" valid:"required,stringlength(1|255)"`
	Success     bool   `json:"success" valid:"required,stringlength(1|255)"`
	ErrCode     string `json:"errCode" valid:"required,stringlength(1|255)"`
	ErrMsg      string `json:"errMsg" valid:"required,stringlength(1|255)"`
	SmsSize     string `json:"smsSizebrand" valid:"required,stringlength(1|255)"`
	BizId       string `json:"bizIdbrand" valid:"required,stringlength(1|255)"`
	OutId       string `json:"outIdbrand" valid:"required,stringlength(1|255)"`
}
