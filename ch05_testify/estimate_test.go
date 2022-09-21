package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	if r := m.Get("Apple"); assert.NotNil(t, r) {
		assert.Equal(t, 200, r.UnitPrice)
		assert.False(t, r.ReducedRate)
	}
	if r := m.Get("Banana"); assert.NotNil(t, r) {
		assert.Equal(t, 250, r.UnitPrice)
		assert.True(t, r.ReducedRate)
	}
	require.Nil(t, m.Get("Grape"))
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
		require.NoError(t, err)
		require.NotNil(t, res)
		assert.Equal(t, "John Smith", res.ClientName)
		assert.False(t, res.EstimatedAt.IsZero())
		assert.Equal(t, 2*200+3*120+4*250, res.SubTotal)
		assert.Equal(t, 2*200*10/100+3*120*8/100+4*250*8/100, res.Tax)
		assert.Equal(t, 2*200+3*120+4*250+2*200*10/100+3*120*8/100+4*250*8/100, res.Total)
		if assert.Len(t, res.Items, 3) {
			assertResponseItem(t, res.Items[0], "Apple", 2, 2*200, 10, 2*200*10/100)
			assertResponseItem(t, res.Items[1], "Orange", 3, 3*120, 8, 3*120*8/100)
			assertResponseItem(t, res.Items[2], "Banana", 4, 4*250, 8, 4*250*8/100)
		}
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
		require.Error(t, err)
	})
}

func assertResponseItem(t *testing.T, actual *ResponseItem, productName string, qty, subTotal, taxRate, tax int) {
	require.NotNil(t, actual)
	assert.Equal(t, productName, actual.ProductName)
	assert.Equal(t, qty, actual.Quantity)
	assert.Equal(t, subTotal, actual.SubTotal)
	assert.Equal(t, taxRate, actual.TaxRate)
	assert.Equal(t, tax, actual.Tax)
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
