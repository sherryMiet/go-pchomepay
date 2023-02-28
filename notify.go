package go_pchomepay_sdk

type PaymentNotify struct {
	NotifyType    string `json:"notify_type"`
	NotifyMessage string `json:"notify_message"`
}

type PaymentNotifyMessage struct {
	OrderId        string                    `json:"order_id"`
	PayType        string                    `json:"pay_type"`
	Amount         string                    `json:"amount"`
	TradeAmount    int                       `json:"trade_amount"`
	PlatformAmount int                       `json:"platform_amount"`
	PPFee          int                       `json:"pp_fee"`
	CreateDate     string                    `json:"create_date"`
	PayDate        string                    `json:"pay_date"`
	ActualPayDate  string                    `json:"actual_pay_date"`
	FailDate       string                    `json:"fail_date"`
	ConfirmDate    string                    `json:"confirm_date"`
	Status         string                    `json:"status"`
	StatusCode     string                    `json:"status_code"`
	PaymentInfo    NotifyPaymentInfo         `json:"payment_info"`
	AvailableDate  string                    `json:"available_date"`
	Items          []PaymentRequestDataItems `json:"items"`
}

type NotifyPaymentInfo struct {
	VirtualAccount string  `json:"virtual_account"`
	BankCode       string  `json:"bank_code"`
	ExpireDate     string  `json:"expire_date"`
	Installment    string  `json:"installment"`
	Rate           float32 `json:"rate"`
	PPRate         float32 `json:"pp_rate"`
	CardLastNumber string  `json:"card_last_number"`
	PPFee          int32   `json:"pp_fee"`
}

func NewNotify() *PaymentNotify {
	return &PaymentNotify{}
}
func NewPaymentNotifyMessage() *PaymentNotifyMessage {
	return &PaymentNotifyMessage{}
}
