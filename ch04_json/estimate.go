package main

import (
	"fmt"
	"time"
)

// 商品の属性を表す型
// 商品JSONファイルをUnmarshalして生成されるので、されるので、コンストラクタは不要
type ProductAttrs struct {
	UnitPrice   int  `json:"unit_price"`
	ReducedRate bool `json:"reduced_rate"`
}

func (m *ProductAttrs) TaxRate() int {
	if m.ReducedRate {
		return 8
	} else {
		return 10
	}
}

// 商品名とProductAttrsを関連付けるmapを拡張した型
// 商品JSONファイルをUnmarshalして生成されるので、されるので、コンストラクタは不要
type ProductMap map[string]*ProductAttrs

func (m ProductMap) Get(product string) *ProductAttrs {
	return m[product]
}

func (m ProductMap) Calculate(req Request) (*Response, error) {
	res := NewResponse(req.ClientName)
	for _, item := range req.Items {
		attrs := m.Get(item.ProductName)
		if attrs == nil {
			return nil, fmt.Errorf("unknown product: %v", item.ProductName)
		}
		res.Items = append(res.Items, NewResponseItem(item.ProductName, attrs, item.Quantity))
	}
	res.Calculate()
	return res, nil
}

// 見積もりRequestを表す型
// 見積もりRequestJSONファイルをUnmarshalして生成されるので、されるので、コンストラクタは不要
type Request struct {
	ClientName string         `json:"client_name"`
	Items      []*RequestItem `json:"items"`
}

// 見積もりの明細を表す型
// 見積もりRequestJSONファイルをUnmarshalして生成されるので、されるので、コンストラクタは不要
type RequestItem struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}

// 見積もり結果を表す型
type Response struct {
	ClientName  string          `json:"client_name"`
	EstimatedAt time.Time       `json:"estimated_at"`
	SubTotal    int             `json:"sub_total"`
	Tax         int             `json:"tax"`
	Total       int             `json:"total"`
	Items       []*ResponseItem `json:"items"`
}

func NewResponse(clientName string) *Response {
	return &Response{ClientName: clientName, EstimatedAt: time.Now()}
}

func (m *Response) Calculate() {
	m.SubTotal = 0
	m.Tax = 0
	for _, item := range m.Items {
		m.SubTotal += item.SubTotal
		m.Tax += item.Tax
	}
	m.Total = m.SubTotal + m.Tax
}

// 見積もり結果の明細を表す型
type ResponseItem struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	SubTotal    int    `json:"sub_total"`
	TaxRate     int    `json:"tax_rate"`
	Tax         int    `json:"tax"`
}

func NewResponseItem(productName string, attrs *ProductAttrs, quantity int) *ResponseItem {
	subTotal := attrs.UnitPrice * quantity
	taxRate := attrs.TaxRate()
	return &ResponseItem{
		ProductName: productName,
		Quantity:    quantity,
		SubTotal:    subTotal,
		TaxRate:     taxRate,
		Tax:         subTotal * taxRate / 100,
	}
}
