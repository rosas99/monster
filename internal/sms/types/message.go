package types

type TemplateMsgRequest struct {
	Matcher   string     `protobuf:"bytes,1,opt,name=matcher,proto3" json:"matcher,omitempty"`
	Request   []any      `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Result    bool       `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	Explains  [][]string `protobuf:"bytes,4,opt,name=explains,proto3" json:"explains,omitempty"`
	Timestamp int64      `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	RequestId string     `protobuf:"bytes,6,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

type UplinkMsgRequest struct {
	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	SendTime    string `protobuf:"bytes,2,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	SignName    string `protobuf:"bytes,4,opt,name=sign_name,json=signName,proto3" json:"sign_name,omitempty"`
	DestCode    string `protobuf:"bytes,5,opt,name=dest_code,json=destCode,proto3" json:"dest_code,omitempty"`
	SequenceId  int64  `protobuf:"varint,6,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	RequestId   string `protobuf:"bytes,6,opt,name=requestId,proto3" json:"requestId,omitempty"`
}
