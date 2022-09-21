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
		assert.Equal(t, []*ResponseItem{
			{ProductName: "Apple", Quantity: 2, SubTotal: 2 * 200, TaxRate: 10, Tax: 2 * 200 * 10 / 100},
			{ProductName: "Orange", Quantity: 3, SubTotal: 3 * 120, TaxRate: 8, Tax: 3 * 120 * 8 / 100},
			{ProductName: "Banana", Quantity: 4, SubTotal: 4 * 250, TaxRate: 8, Tax: 4 * 250 * 8 / 100},
		}, res.Items)
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

func TestNewResponseItem(t *testing.T) {
	m := newTestProductMap()
	assert.Equal(t,
		NewResponseItem("Apple", m.Get("Apple"), 2),
		&ResponseItem{ProductName: "Apple", Quantity: 2, SubTotal: 2 * 200, TaxRate: 10, Tax: 2 * 200 * 10 / 100},
	)
	assert.Equal(t,
		NewResponseItem("Orange", m.Get("Orange"), 3),
		&ResponseItem{ProductName: "Orange", Quantity: 3, SubTotal: 3 * 120, TaxRate: 8, Tax: 3 * 120 * 8 / 100},
	)
}
