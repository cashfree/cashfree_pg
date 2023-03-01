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

// CFPaymentsEntityUPIPayment struct for CFPaymentsEntityUPIPayment
type CFPaymentsEntityUPIPayment struct {
	Channel string `json:"channel"`
	UpiId *string `json:"upi_id,omitempty"`
}

// NewCFPaymentsEntityUPIPayment instantiates a new CFPaymentsEntityUPIPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFPaymentsEntityUPIPayment(channel string) *CFPaymentsEntityUPIPayment {
	this := CFPaymentsEntityUPIPayment{}
	this.Channel = channel
	return &this
}

// NewCFPaymentsEntityUPIPaymentWithDefaults instantiates a new CFPaymentsEntityUPIPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFPaymentsEntityUPIPaymentWithDefaults() *CFPaymentsEntityUPIPayment {
	this := CFPaymentsEntityUPIPayment{}
	return &this
}

// GetChannel returns the Channel field value
func (o *CFPaymentsEntityUPIPayment) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *CFPaymentsEntityUPIPayment) GetChannelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *CFPaymentsEntityUPIPayment) SetChannel(v string) {
	o.Channel = v
}

// GetUpiId returns the UpiId field value if set, zero value otherwise.
func (o *CFPaymentsEntityUPIPayment) GetUpiId() string {
	if o == nil || o.UpiId == nil {
		var ret string
		return ret
	}
	return *o.UpiId
}

// GetUpiIdOk returns a tuple with the UpiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFPaymentsEntityUPIPayment) GetUpiIdOk() (*string, bool) {
	if o == nil || o.UpiId == nil {
		return nil, false
	}
	return o.UpiId, true
}

// HasUpiId returns a boolean if a field has been set.
func (o *CFPaymentsEntityUPIPayment) HasUpiId() bool {
	if o != nil && o.UpiId != nil {
		return true
	}

	return false
}

// SetUpiId gets a reference to the given string and assigns it to the UpiId field.
func (o *CFPaymentsEntityUPIPayment) SetUpiId(v string) {
	o.UpiId = &v
}

func (o CFPaymentsEntityUPIPayment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["channel"] = o.Channel
	}
	if o.UpiId != nil {
		toSerialize["upi_id"] = o.UpiId
	}
	return json.Marshal(toSerialize)
}

type NullableCFPaymentsEntityUPIPayment struct {
	value *CFPaymentsEntityUPIPayment
	isSet bool
}

func (v NullableCFPaymentsEntityUPIPayment) Get() *CFPaymentsEntityUPIPayment {
	return v.value
}

func (v *NullableCFPaymentsEntityUPIPayment) Set(val *CFPaymentsEntityUPIPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableCFPaymentsEntityUPIPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableCFPaymentsEntityUPIPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFPaymentsEntityUPIPayment(val *CFPaymentsEntityUPIPayment) *NullableCFPaymentsEntityUPIPayment {
	return &NullableCFPaymentsEntityUPIPayment{value: val, isSet: true}
}

func (v NullableCFPaymentsEntityUPIPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFPaymentsEntityUPIPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


