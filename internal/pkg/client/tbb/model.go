package tbb

type AddTransactionRequest struct {
	From    string
	To      string
	Value   int
	FromPWD string
}

type AddTransactionResponse struct{}
