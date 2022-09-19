package main

import "time"

type ProductAttrs struct {
	UnitPrice   int
	ReducedRate bool
}

type ProductMap map[string]*ProductAttrs

func (m ProductMap) Get(product string) *ProductAttrs {
	return nil
}

func (m ProductMap) Calculate(req Request) *Response {
	return nil
}

type Request struct {
	ClientName string
	Items      []*RequestItem
}

type RequestItem struct {
	ProductName string
	Quantity    int
}

type Response struct {
	ClientName  string
	EstimatedAt time.Time
	SubTotal    int
	Tax         int
	Total       int
	Items       []*ResponseItem
}

func (m *Response) Calculate() {
}

type ResponseItem struct {
	ProductName string
	Quantity    int
	SubTotal    int
	TaxRate     int
	Tax         int
}

func (m *ResponseItem) Calculate(attrs *ProductAttrs, quantity int) {
}
