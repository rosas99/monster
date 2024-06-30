package types

type TemplateMsgRequest struct {
	Matcher   string     `protobuf:"bytes,1,opt,name=matcher,proto3" json:"matcher,omitempty"`
	Request   []any      `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Result    bool       `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	Explains  [][]string `protobuf:"bytes,4,opt,name=explains,proto3" json:"explains,omitempty"`
	Timestamp int64      `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	RequestId string     `protobuf:"bytes,6,opt,name=requestId,proto3" json:"requestId,omitempty"`
}
