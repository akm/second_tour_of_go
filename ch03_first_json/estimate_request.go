package main

import (
	"fmt"
	"strings"
	"time"
)

type EstimateRequest struct {
	ClientID    *int64               `json:"client_id,omitempty"`
	ClientName  string               `json:"client_name"`
	Items       EstimateRequestItems `json:"items"`
	RequestedAt time.Time            `json:"requested_at"`
}

func (m *EstimateRequest) ClientDisplayName() string {
	if m.ClientID != nil {
		return fmt.Sprintf("%s(%d)", m.ClientName, *m.ClientID)
	}
	return m.ClientName
}

func (m *EstimateRequest) Text() string {
	return fmt.Sprintf("Request from %s at %s\n", m.ClientDisplayName(), m.RequestedAt.Format(time.RFC3339)) +
		m.Items.Text()
}

type EstimateRequestItems []*EstimateRequestItem

func (s EstimateRequestItems) Text() string {
	lines := make([]string, len(s))
	for i, item := range s {
		lines[i] = item.String()
	}
	return strings.Join(lines, "\n")
}

type EstimateRequestItem struct {
	ProductID   *int64 `json:"product_id,omitempty"`
	ProductName string `json:"product_name"`
	Amount      int    `json:"amount"`
}

func (m *EstimateRequestItem) ProductDisplayName() string {
	if m.ProductID != nil {
		return fmt.Sprintf("%s(%d)", m.ProductName, *m.ProductID)
	}
	return m.ProductName
}

func (m *EstimateRequestItem) String() string {
	return fmt.Sprintf("%s: %d", m.ProductDisplayName(), m.Amount)
}
