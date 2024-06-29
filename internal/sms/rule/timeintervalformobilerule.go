package rule

type TimeIntervalForMobileRule struct{}

func (m TimeIntervalForMobileRule) IsValid(request *CheckerRequest) bool {
	//TODO implement me
	panic("implement me")
}

func (m TimeIntervalForMobileRule) GetFailReason() *CheckerResponse {
	//TODO implement me
	panic("implement me")
}

var _ Rule = (*TimeIntervalForMobileRule)(nil)
