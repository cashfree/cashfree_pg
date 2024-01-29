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

// checks if the PaymentMethodAppInPaymentsEntityApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodAppInPaymentsEntityApp{}

// PaymentMethodAppInPaymentsEntityApp struct for PaymentMethodAppInPaymentsEntityApp
type PaymentMethodAppInPaymentsEntityApp struct {
	Channel *string `json:"channel,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Phone *string `json:"phone,omitempty"`
}


func (o PaymentMethodAppInPaymentsEntityApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodAppInPaymentsEntityApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	return toSerialize, nil
}



