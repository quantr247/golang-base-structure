package model

type MGPSRequestData struct {
	RequestID     string
	TransactionID string
	Description   string
	Data          string
}

type MGPSResponseData struct {
	RequestID     string
	TransactionID string
	Result        string
	StatusCode    string
	StatusMessage string
}
