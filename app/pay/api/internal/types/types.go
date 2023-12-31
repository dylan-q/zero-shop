// Code generated by goctl. DO NOT EDIT.
package types

type NotifyReq struct {
	TradeType      string  `form:"trade_type"`       // 支付类型
	TradeState     string  `form:"trade_state"`      // 交易状态
	PayTotal       float64 `form:"pay_total"`        // 支付金额
	TransactionId  string  `form:"transaction_id"`   // 微信的订单号
	TradeStateDesc string  `form:"trade_state_desc"` // 订单支付描述
}

type CreateThirdReq struct {
	GoodsId     int64   `json:"goods_id"`
	MarketPrice float64 `json:"market_price"`
	SalePrice   float64 `json:"sale_price"`
	OrderStatus int64   `json:"order_status"`
}

type CreateThirdResp struct {
	OrderSn string `json:"order_sn"`
}

type NotifyResp struct {
}
