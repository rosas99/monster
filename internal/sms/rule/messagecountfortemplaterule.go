package rule

type MessageCountForTemplateRule struct{}

func (m MessageCountForTemplateRule) IsValid(request *CheckerRequest) bool {
	//TODO implement me
	panic("implement me")
}

func (m MessageCountForTemplateRule) GetFailReason() *CheckerResponse {
	//TODO implement me
	panic("implement me")
}

var _ Rule = (*MessageCountForTemplateRule)(nil)
