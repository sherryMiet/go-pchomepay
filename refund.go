package go_pchomepay_sdk

import "encoding/json"

const (
	RefundUrl     = "https://api.pchomepay.com.tw/v2/refund"
	TestRefundUrl = "https://sandbox-api.pchomepay.com.tw/v2/refund"
)

type RefundRequestCall struct {
	Token             string
	RefundRequestData *RefundRequestData
}

type RefundRequestData struct {
	OrderId     string `json:"order_id"`
	RefundId    string `json:"refund_id"`
	TradeAmount int    `json:"trade_amount"`
}
type RefundResponseData struct {
	OrderId     string `json:"order_id"`
	RefundId    string `json:"refund_id"`
	PayType     string `json:"pay_type"`
	TradeAmt    int    `json:"trade_amt"`
	Fee         int    `json:"fee"`
	TransferFee int    `json:"transfer_fee"`
	Status      string `json:"status"`
}

func NewRefund() *RefundRequestData {
	return &RefundRequestData{}
}

func (p *RefundRequestData) CreateRefund(OrderId, RefundId string, TradeAmt int) *RefundRequestData {
	p.RefundId = RefundId
	p.OrderId = OrderId
	p.TradeAmount = TradeAmt
	return p
}

func (c *Client) Refund(Data *RefundRequestData) *RefundRequestCall {
	token, err := c.GetTokenTest()
	if err != nil {
		return nil
	}
	return &RefundRequestCall{
		Token:             token.Token,
		RefundRequestData: Data,
	}
}
func (c *Client) RefundTest(Data *RefundRequestData) *RefundRequestCall {
	token, err := c.GetToken()
	if err != nil {
		return nil
	}
	return &RefundRequestCall{
		Token:             token.Token,
		RefundRequestData: Data,
	}
}

func (p RefundRequestCall) Do() (res *RefundResponseData, err error) {
	marshal, err := json.Marshal(p.RefundRequestData)
	if err != nil {
		return
	}
	print(string(marshal))
	request, err := SendPCHOMEPayRequest(marshal, p.Token, RefundUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(request, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p RefundRequestCall) DoTest() (res *RefundResponseData, err error) {
	marshal, err := json.Marshal(p.RefundRequestData)
	if err != nil {
		return
	}
	request, err := SendPCHOMEPayRequest(marshal, p.Token, TestRefundUrl)
	if err != nil {
		return
	}
	print(string(request))
	err = json.Unmarshal(request, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
