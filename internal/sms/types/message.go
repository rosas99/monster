package types

// TemplateMsgRequest defines a template message request for kafka queue.
type TemplateMsgRequest struct {
	PhoneNumber  string   `json:"phoneNumber"`
	SendTime     string   `json:"sendTime"`
	Content      string   `json:"content"`
	SignName     string   `json:"signName"`
	DestCode     string   `json:"destCode"`
	SequenceId   int64    `json:"sequenceId"`
	RequestId    string   `json:"requestId"`
	TemplateCode string   `json:"templateCode"`
	Providers    []string `json:"providers"`
}

// UplinkMsgRequest defines an uplink message request for kafka queue.
type UplinkMsgRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	SendTime    string `json:"sendTime"`
	Content     string `json:"content"`
	SignName    string `json:"signName"`
	DestCode    string `json:"destCode"`
	SequenceId  int64  `json:"sequenceId"`
	RequestId   string `json:"requestId"`
}
