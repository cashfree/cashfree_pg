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

// checks if the CardlessEMIQueries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardlessEMIQueries{}

// CardlessEMIQueries cardless EMI query object
type CardlessEMIQueries struct {
	// OrderId of the order. Either of `order_id` or `amount` is mandatory.
	OrderId *string `json:"order_id,omitempty"`
	// Amount of the order. OrderId of the order. Either of `order_id` or `amount` is mandatory.
	Amount *float32 `json:"amount,omitempty"`
	CustomerDetails *CustomerDetailsCardlessEMI `json:"customer_details,omitempty"`
}


func (o CardlessEMIQueries) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardlessEMIQueries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CustomerDetails) {
		toSerialize["customer_details"] = o.CustomerDetails
	}
	return toSerialize, nil
}



