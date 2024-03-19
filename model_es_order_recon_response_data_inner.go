/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2023-08-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the ESOrderReconResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ESOrderReconResponseDataInner{}

// ESOrderReconResponseDataInner struct for ESOrderReconResponseDataInner
type ESOrderReconResponseDataInner struct {
	Amount *float32 `json:"amount,omitempty"`
	SettlementEligibilityTime *string `json:"settlement_eligibility_time,omitempty"`
	MerchantOrderId *string `json:"merchant_order_id,omitempty"`
	TxTime *string `json:"tx_time,omitempty"`
	Settled *string `json:"settled,omitempty"`
	EntityId *string `json:"entity_id,omitempty"`
	MerchantSettlementUtr *string `json:"merchant_settlement_utr,omitempty"`
	Currency *string `json:"currency,omitempty"`
	SaleType *string `json:"sale_type,omitempty"`
	CustomerName *string `json:"customer_name,omitempty"`
	CustomerEmail *string `json:"customer_email,omitempty"`
	CustomerPhone *string `json:"customer_phone,omitempty"`
	MerchantVendorCommission *string `json:"merchant_vendor_commission,omitempty"`
	SplitServiceCharge *string `json:"split_service_charge,omitempty"`
	SplitServiceTax *string `json:"split_service_tax,omitempty"`
	PgServiceTax *string `json:"pg_service_tax,omitempty"`
	PgServiceCharge *string `json:"pg_service_charge,omitempty"`
	PgChargePostpaid *string `json:"pg_charge_postpaid,omitempty"`
	MerchantSettlementId *string `json:"merchant_settlement_id,omitempty"`
	AddedOn *string `json:"added_on,omitempty"`
	Tags *string `json:"tags,omitempty"`
	EntityType *string `json:"entity_type,omitempty"`
	SettlementInitiatedOn *string `json:"settlement_initiated_on,omitempty"`
	SettlementTime *string `json:"settlement_time,omitempty"`
	OrderSplits []ESOrderReconResponseDataInnerOrderSplitsInner `json:"order_splits,omitempty"`
	EligibleSplitBalance *string `json:"eligible_split_balance,omitempty"`
}


func (o ESOrderReconResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ESOrderReconResponseDataInner) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.SettlementEligibilityTime) {
		toSerialize["settlement_eligibility_time"] = o.SettlementEligibilityTime
	}
	if !IsNil(o.MerchantOrderId) {
		toSerialize["merchant_order_id"] = o.MerchantOrderId
	}
	if !IsNil(o.TxTime) {
		toSerialize["tx_time"] = o.TxTime
	}
	if !IsNil(o.Settled) {
		toSerialize["settled"] = o.Settled
	}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if !IsNil(o.MerchantSettlementUtr) {
		toSerialize["merchant_settlement_utr"] = o.MerchantSettlementUtr
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.SaleType) {
		toSerialize["sale_type"] = o.SaleType
	}
	if !IsNil(o.CustomerName) {
		toSerialize["customer_name"] = o.CustomerName
	}
	if !IsNil(o.CustomerEmail) {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if !IsNil(o.CustomerPhone) {
		toSerialize["customer_phone"] = o.CustomerPhone
	}
	if !IsNil(o.MerchantVendorCommission) {
		toSerialize["merchant_vendor_commission"] = o.MerchantVendorCommission
	}
	if !IsNil(o.SplitServiceCharge) {
		toSerialize["split_service_charge"] = o.SplitServiceCharge
	}
	if !IsNil(o.SplitServiceTax) {
		toSerialize["split_service_tax"] = o.SplitServiceTax
	}
	if !IsNil(o.PgServiceTax) {
		toSerialize["pg_service_tax"] = o.PgServiceTax
	}
	if !IsNil(o.PgServiceCharge) {
		toSerialize["pg_service_charge"] = o.PgServiceCharge
	}
	if !IsNil(o.PgChargePostpaid) {
		toSerialize["pg_charge_postpaid"] = o.PgChargePostpaid
	}
	if !IsNil(o.MerchantSettlementId) {
		toSerialize["merchant_settlement_id"] = o.MerchantSettlementId
	}
	if !IsNil(o.AddedOn) {
		toSerialize["added_on"] = o.AddedOn
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.EntityType) {
		toSerialize["entity_type"] = o.EntityType
	}
	if !IsNil(o.SettlementInitiatedOn) {
		toSerialize["settlement_initiated_on"] = o.SettlementInitiatedOn
	}
	if !IsNil(o.SettlementTime) {
		toSerialize["settlement_time"] = o.SettlementTime
	}
	if !IsNil(o.OrderSplits) {
		toSerialize["order_splits"] = o.OrderSplits
	}
	if !IsNil(o.EligibleSplitBalance) {
		toSerialize["eligible_split_balance"] = o.EligibleSplitBalance
	}
	return toSerialize, nil
}



