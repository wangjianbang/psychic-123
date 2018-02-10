package models


type OrderStatus struct { //好像不一定要全部齐全？后面删除一个确认下就行
	Id              float64 `json:"id"`
	Symbol          string  `json:"symbol"`            //交易对
	AccountId       float64 `json:"account-id"`        //订单 ID
	Amount          string  `json:"amount"`            //订单数量
	Price           string  `json:"price"`             //订单价格
	CreatedAt       float64 `json:"created-at"`        //订单创建时间
	Type            string  `json:"type"`              //订单类型：buy-market：市价买,sell-limit：限价卖
	FieldAmount     string  `json:"field-amount"`      //已成交数量
	FieldCashAmount string  `json:"field-cash-amount"` //已成交总金额
	FieldFees       string  `json:"field-fees"`        //已成交手续费（买入为币，卖出为钱）
	FinishedAt      float64 `json:"finished-at"`       //最后成交时间
	UserId          float64 `json:"user-id"`           //账户 ID
	Source          string  `json:"source"`            //订单来源
	State           string  `json:"state"`             //订单状态  submitted 已提交, partial-filled 部分成交, partial-canceled 部分成交撤销, filled 完全成交, canceled 已撤销
	CanceledAt      float64 `json:"canceled-at"`       //订单撤销时间
	Exchange        string  `json:"exchange"`          //huobi
	Batch           string  `json:"batch"`             //？？
}

type OrderStatusReturn struct {
	Status  string        `json:"status"`
	Data    OrderStatus `json:"data"`
	ErrCode string        `json:"err-code"`
	ErrMsg  string        `json:"err-msg"`
}
