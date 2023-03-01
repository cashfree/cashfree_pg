/*
New Payment Gateway APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2022-01-01
Contact: nextgenapi@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg_sdk_go

import (
	"encoding/json"
)

// CFCardlessEMIPayment struct for CFCardlessEMIPayment
type CFCardlessEMIPayment struct {
	CardlessEmi CFCardlessEMI `json:"cardless_emi"`
}

// NewCFCardlessEMIPayment instantiates a new CFCardlessEMIPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFCardlessEMIPayment(cardlessEmi CFCardlessEMI) *CFCardlessEMIPayment {
	this := CFCardlessEMIPayment{}
	this.CardlessEmi = cardlessEmi
	return &this
}

// NewCFCardlessEMIPaymentWithDefaults instantiates a new CFCardlessEMIPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFCardlessEMIPaymentWithDefaults() *CFCardlessEMIPayment {
	this := CFCardlessEMIPayment{}
	return &this
}

// GetCardlessEmi returns the CardlessEmi field value
func (o *CFCardlessEMIPayment) GetCardlessEmi() CFCardlessEMI {
	if o == nil {
		var ret CFCardlessEMI
		return ret
	}

	return o.CardlessEmi
}

// GetCardlessEmiOk returns a tuple with the CardlessEmi field value
// and a boolean to check if the value has been set.
func (o *CFCardlessEMIPayment) GetCardlessEmiOk() (*CFCardlessEMI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardlessEmi, true
}

// SetCardlessEmi sets field value
func (o *CFCardlessEMIPayment) SetCardlessEmi(v CFCardlessEMI) {
	o.CardlessEmi = v
}

func (o CFCardlessEMIPayment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cardless_emi"] = o.CardlessEmi
	}
	return json.Marshal(toSerialize)
}

type NullableCFCardlessEMIPayment struct {
	value *CFCardlessEMIPayment
	isSet bool
}

func (v NullableCFCardlessEMIPayment) Get() *CFCardlessEMIPayment {
	return v.value
}

func (v *NullableCFCardlessEMIPayment) Set(val *CFCardlessEMIPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableCFCardlessEMIPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableCFCardlessEMIPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFCardlessEMIPayment(val *CFCardlessEMIPayment) *NullableCFCardlessEMIPayment {
	return &NullableCFCardlessEMIPayment{value: val, isSet: true}
}

func (v NullableCFCardlessEMIPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFCardlessEMIPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


