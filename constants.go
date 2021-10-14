package yookassa

const (
	APIEndpoint = "https://api.yookassa.ru/v3/"
	PaymentsEndpoint = "payments/"
	CaptureEndpoint = "capture"
)

const (
	ErrorType = "error"
)

type Status string
const (
	Pending           Status = "pending"
	WaitingForCapture Status = "waiting_for_capture"
	Succeeded         Status = "succeeded"
	Canceled          Status = "canceled"
)