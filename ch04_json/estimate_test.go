package main

import (
	"testing"
)

func newTestProductMap() ProductMap {
	return ProductMap{
		"Apple":      {UnitPrice: 200, ReducedRate: false},
		"Orange":     {UnitPrice: 120, ReducedRate: true},
		"Banana":     {UnitPrice: 250, ReducedRate: true},
		"Kiwi Fruit": {UnitPrice: 100, ReducedRate: true},
		"Lemon":      {UnitPrice: 150, ReducedRate: false},
	}
}

func TestProductMapGet(t *testing.T) {
	m := newTestProductMap()
	if r := m.Get("Apple"); r == nil || r.UnitPrice != 200 || r.ReducedRate != false {
		t.Errorf("m.Get(\"Apple\") = %v, want {UnitPrice: 200, ReducedRate: false}", r)
	}
	if r := m.Get("Banana"); r == nil || r.UnitPrice != 250 || r.ReducedRate != true {
		t.Errorf("m.Get(\"Banana\") = %v, want {UnitPrice: 250, ReducedRate: true}", r)
	}
	if r := m.Get("Grape"); r != nil {
		t.Errorf("m.Get(\"Grape\") = %v, want nil", r)
	}
}

func TestProductMapCalculate(t *testing.T) {
	t.Run("basic pattern", func(t *testing.T) {
		m := newTestProductMap()
		res, err := m.Calculate(Request{
			ClientName: "John Smith",
			Items: []*RequestItem{
				{ProductName: "Apple", Quantity: 2},
				{ProductName: "Orange", Quantity: 3},
				{ProductName: "Banana", Quantity: 4},
			},
		})
		if err != nil {
			t.Errorf("m.Calculate() = %v, want nil", err)
		}
		if res == nil {
			t.Fatalf("m.Calculate() = nil, want not nil")
		}
		if res.ClientName != "John Smith" {
			t.Errorf("res.ClientName = %v, want \"John Smith\"", res.ClientName)
		}
		if res.EstimatedAt.IsZero() {
			t.Errorf("res.EstimatedAt = %v, want non-zero", res.EstimatedAt)
		}
		if res.SubTotal != 2*200+3*120+4*250 {
			t.Errorf("res.SubTotal = %v, want %v", res.SubTotal, 2*200+3*120+4*250)
		}
		if res.Tax != 2*200*10/100+3*120*8/100+4*250*8/100 {
			t.Errorf("res.Tax = %v, want %v", res.Tax, 2*200*10/100+3*120*8/100+4*250*8/100)
		}
		if res.Total != 2*200+3*120+4*250+2*200*10/100+3*120*8/100+4*250*8/100 {
			t.Errorf("res.Total = %v, want %v", res.Total, 2*200+3*120+4*250+2*200*10/100+3*120*8/100+4*250*8/100)
		}
		if len(res.Items) != 3 {
			t.Fatalf("len(res.Items) = %v, want 3", len(res.Items))
		}
		assertResponseItem(t, res.Items[0], "Apple", 2, 2*200, 10, 2*200*10/100)
		assertResponseItem(t, res.Items[1], "Orange", 3, 3*120, 8, 3*120*8/100)
		assertResponseItem(t, res.Items[2], "Banana", 4, 4*250, 8, 4*250*8/100)
	})

	t.Run("including unknown product", func(t *testing.T) {
		m := newTestProductMap()
		_, err := m.Calculate(Request{
			ClientName: "John Smith",
			Items: []*RequestItem{
				{ProductName: "Apple", Quantity: 2},
				{ProductName: "Grape", Quantity: 3},
				{ProductName: "Banana", Quantity: 4},
			},
		})
		if err == nil {
			t.Errorf("m.Calculate() should be return an error")
		}
	})
}

func assertResponseItem(t *testing.T, actual *ResponseItem, productName string, qty, subTotal, taxRate, tax int) {
	if actual == nil {
		t.Fatalf("actual = nil, want not nil")
	}
	if actual.ProductName != productName {
		t.Errorf("actual.ProductName = %v, want %v", actual.ProductName, productName)
	}
	if actual.Quantity != qty {
		t.Errorf("actual.Quantity = %v, want %v", actual.Quantity, qty)
	}
	if actual.SubTotal != subTotal {
		t.Errorf("actual.SubTotal = %v, want %v", actual.SubTotal, subTotal)
	}
	if actual.TaxRate != taxRate {
		t.Errorf("actual.TaxRate = %v, want %v", actual.TaxRate, taxRate)
	}
	if actual.Tax != tax {
		t.Errorf("actual.Tax = %v, want %v", actual.Tax, tax)
	}
}

func TestNewResponseItem(t *testing.T) {
	m := newTestProductMap()
	assertResponseItem(t,
		NewResponseItem("Apple", m.Get("Apple"), 2),
		"Apple", 2, 2*200, 10, 2*200*10/100,
	)
	assertResponseItem(t,
		NewResponseItem("Orange", m.Get("Orange"), 3),
		"Orange", 3, 3*120, 8, 3*120*8/100,
	)
}
