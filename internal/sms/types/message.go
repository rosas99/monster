package types

// TemplateMsgRequest defines a template message request for kafka queue.
type TemplateMsgRequest struct {
	PhoneNumber  string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	SendTime     string `protobuf:"bytes,2,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	Content      string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	SignName     string `protobuf:"bytes,4,opt,name=sign_name,json=signName,proto3" json:"sign_name,omitempty"`
	DestCode     string `protobuf:"bytes,5,opt,name=dest_code,json=destCode,proto3" json:"dest_code,omitempty"`
	SequenceId   int64  `protobuf:"varint,6,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	RequestId    string `protobuf:"bytes,6,opt,name=requestId,proto3" json:"requestId,omitempty"`
	TemplateCode string `gorm:"column:template_code;not null" json:"template_code"`
}

// UplinkMsgRequest defines an uplink message request for kafka queue.
type UplinkMsgRequest struct {
	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	SendTime    string `protobuf:"bytes,2,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	SignName    string `protobuf:"bytes,4,opt,name=sign_name,json=signName,proto3" json:"sign_name,omitempty"`
	DestCode    string `protobuf:"bytes,5,opt,name=dest_code,json=destCode,proto3" json:"dest_code,omitempty"`
	SequenceId  int64  `protobuf:"varint,6,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	RequestId   string `protobuf:"bytes,6,opt,name=requestId,proto3" json:"requestId,omitempty"`
}
