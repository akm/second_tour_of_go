package main

import (
	"fmt"
	"time"
)

// 商品の属性を表す型
// 商品JSONファイルをUnmarshalして生成されるので、されるので、コンストラクタは不要
type ProductAttrs struct {
	UnitPrice   int
	ReducedRate bool
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
	ClientName string
	Items      []*RequestItem
}

// 見積もりの明細を表す型
// 見積もりRequestJSONファイルをUnmarshalして生成されるので、されるので、コンストラクタは不要
type RequestItem struct {
	ProductName string
	Quantity    int
}

// 見積もり結果を表す型
type Response struct {
	ClientName  string
	EstimatedAt time.Time
	SubTotal    int
	Tax         int
	Total       int
	Items       []*ResponseItem
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
	ProductName string
	Quantity    int
	SubTotal    int
	TaxRate     int
	Tax         int
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
