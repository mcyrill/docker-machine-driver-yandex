// Code generated by protoc-gen-goext. DO NOT EDIT.

package billing

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Sku) SetId(v string) {
	m.Id = v
}

func (m *Sku) SetName(v string) {
	m.Name = v
}

func (m *Sku) SetDescription(v string) {
	m.Description = v
}

func (m *Sku) SetServiceId(v string) {
	m.ServiceId = v
}

func (m *Sku) SetPricingUnit(v string) {
	m.PricingUnit = v
}

func (m *Sku) SetPricingVersions(v []*PricingVersion) {
	m.PricingVersions = v
}

func (m *PricingVersion) SetType(v PricingVersionType) {
	m.Type = v
}

func (m *PricingVersion) SetEffectiveTime(v *timestamppb.Timestamp) {
	m.EffectiveTime = v
}

func (m *PricingVersion) SetPricingExpressions(v []*PricingExpression) {
	m.PricingExpressions = v
}

func (m *PricingExpression) SetRates(v []*Rate) {
	m.Rates = v
}

func (m *Rate) SetStartPricingQuantity(v string) {
	m.StartPricingQuantity = v
}

func (m *Rate) SetUnitPrice(v string) {
	m.UnitPrice = v
}

func (m *Rate) SetCurrency(v string) {
	m.Currency = v
}
