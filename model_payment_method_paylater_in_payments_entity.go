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

// checks if the PaymentMethodPaylaterInPaymentsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodPaylaterInPaymentsEntity{}

// PaymentMethodPaylaterInPaymentsEntity paylater payment method object for pay api
type PaymentMethodPaylaterInPaymentsEntity struct {
	Paylater *PaymentMethodAppInPaymentsEntityApp `json:"paylater,omitempty"`
}


func (o PaymentMethodPaylaterInPaymentsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodPaylaterInPaymentsEntity) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paylater) {
		toSerialize["paylater"] = o.Paylater
	}
	return toSerialize, nil
}



