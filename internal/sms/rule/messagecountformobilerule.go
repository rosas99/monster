package rule

type MessageCountForMobileRule struct{}

func (m MessageCountForMobileRule) IsValid(request *CheckerRequest) bool {
	//TODO implement me
	panic("implement me")
}

func (m MessageCountForMobileRule) GetFailReason() *CheckerResponse {
	//TODO implement me
	panic("implement me")
}

var _ Rule = (*MessageCountForMobileRule)(nil)
