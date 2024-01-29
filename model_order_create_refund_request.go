/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2022-09-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the OrderCreateRefundRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRefundRequest{}

// OrderCreateRefundRequest create refund request object
type OrderCreateRefundRequest struct {
	// Amount to be refunded. Should be lesser than or equal to the transaction amount. (Decimals allowed)
	RefundAmount float64 `json:"refund_amount"`
	// An unique ID to associate the refund with. Provie alphanumeric values
	RefundId string `json:"refund_id"`
	// A refund note for your reference.
	RefundNote *string `json:"refund_note,omitempty"`
	// Speed at which the refund is processed. It's an optional field with default being STANDARD
	RefundSpeed *string `json:"refund_speed,omitempty"`
	RefundSplits []VendorSplit `json:"refund_splits,omitempty"`
}


func (o OrderCreateRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRefundRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["refund_amount"] = o.RefundAmount
	toSerialize["refund_id"] = o.RefundId
	if !IsNil(o.RefundNote) {
		toSerialize["refund_note"] = o.RefundNote
	}
	if !IsNil(o.RefundSpeed) {
		toSerialize["refund_speed"] = o.RefundSpeed
	}
	if !IsNil(o.RefundSplits) {
		toSerialize["refund_splits"] = o.RefundSplits
	}
	return toSerialize, nil
}



