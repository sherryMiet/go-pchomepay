package go_pchomepay_sdk

import (
	"encoding/json"
)

const (
	PaymentURL     = "https://api.pchomepay.com.tw/v1/payment"
	TestPaymentURL = "https://sandbox-api.pchomepay.com.tw/v1/payment"
)

type PaymentRequestCall struct {
	Token              string
	PaymentRequestData *PaymentRequestData
}

type PaymentRequestData struct {
	OrderId         string                        `json:"order_id"`
	PayType         []string                      `json:"pay_type"`
	Amount          int                           `json:"amount"`
	ReturnURL       string                        `json:"return_url,omitempty"`
	FailReturnURL   string                        `json:"fail_return_url,omitempty"`
	NotifyURL       string                        `json:"notify_url,omitempty"`
	Items           []PaymentRequestDataItems     `json:"items"`
	ATMInfo         *PaymentRequestDataATMInfo    `json:"atm_info,omitempty"`
	CardInfo        []*PaymentRequestDataCardInfo `json:"card_info,omitempty"`
	CardInstallment string                        `json:"card_installment,omitempty"`
	ReturnTimer     string                        `json:"return_timer"`
	MemberKey       string                        `json:"member_key"`
}

type PaymentResponseData struct {
	OrderId    string `json:"order_id"`
	PaymentURL string `json:"payment_url"`
}

type PaymentRequestDataItems struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PaymentRequestDataATMInfo struct {
	ExpireDays int `json:"expire_days"`
}

type PaymentRequestDataCardInfo struct {
	Installment int `json:"installment"`
}

func NewItem() *PaymentRequestDataItems {
	return &PaymentRequestDataItems{}
}

func NewItems() *[]PaymentRequestDataItems {
	return &[]PaymentRequestDataItems{}
}

func NewPayment() *PaymentRequestData {
	return &PaymentRequestData{}
}

func (p *PaymentRequestData) SetItems(Items []PaymentRequestDataItems) *PaymentRequestData {
	p.Items = Items
	return p

}
func (p *PaymentRequestData) SetReturnURL(ReturnURL, FailReturnURL string) *PaymentRequestData {
	p.ReturnURL = ReturnURL
	p.FailReturnURL = FailReturnURL
	return p
}

func (p *PaymentRequestData) SetNotifyURL(NotifyURL string) *PaymentRequestData {
	p.NotifyURL = NotifyURL
	return p
}

func (p *PaymentRequestData) CreateOrder(OrderId string, Amount int) *PaymentRequestData {
	p.OrderId = OrderId
	p.Amount = Amount
	return p
}
func (p *PaymentRequestData) SetATM(ExpireDays int) *PaymentRequestData {
	p.ATMInfo = &PaymentRequestDataATMInfo{
		ExpireDays: ExpireDays,
	}
	p.PayType = []string{"ATM"}
	return p
}

func (p *PaymentRequestData) SetCard(CardInstallment string) *PaymentRequestData {
	//Installment := &PaymentRequestDataCardInfo{
	//	Installment: CardInstallment,
	//}
	//p.CardInfo = append(p.CardInfo, Installment)
	p.PayType = []string{"CARD"}
	p.CardInstallment = CardInstallment
	return p
}

func (c *Client) Payment(Data *PaymentRequestData) *PaymentRequestCall {
	token, err := c.GetToken()
	if err != nil {
		return nil
	}
	return &PaymentRequestCall{
		Token:              token.Token,
		PaymentRequestData: Data,
	}
}
func (c *Client) PaymentTest(Data *PaymentRequestData) *PaymentRequestCall {
	token, err := c.GetTokenTest()
	if err != nil {
		return nil
	}
	return &PaymentRequestCall{
		Token:              token.Token,
		PaymentRequestData: Data,
	}
}
func (p PaymentRequestCall) Do() (res *PaymentResponseData, err error) {
	marshal, err := json.Marshal(p.PaymentRequestData)
	if err != nil {
		return
	}
	print(string(marshal))
	request, err := SendPCHOMEPayRequest(marshal, p.Token, PaymentURL)
	if err != nil {
		return
	}
	err = json.Unmarshal(request, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p PaymentRequestCall) DoTest() (res *PaymentResponseData, err error) {
	marshal, err := json.Marshal(p.PaymentRequestData)
	if err != nil {
		return
	}
	print(string(marshal))
	request, err := SendPCHOMEPayRequest(marshal, p.Token, TestPaymentURL)
	if err != nil {
		return
	}
	err = json.Unmarshal(request, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
