package main

import "time"

// 商品の属性を表す型
// 商品JSONファイルをUnmarshalして生成されるので、されるので、コンストラクタは不要
type ProductAttrs struct {
	UnitPrice   int
	ReducedRate bool
}

// 商品名とProductAttrsを関連付けるmapを拡張した型
// 商品JSONファイルをUnmarshalして生成されるので、されるので、コンストラクタは不要
type ProductMap map[string]*ProductAttrs

func (m ProductMap) Get(product string) *ProductAttrs {
	return m[product]
}

func (m ProductMap) Calculate(req Request) *Response {
	return nil
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
	return nil
}

func (m *Response) Calculate() {
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
	return nil
}
