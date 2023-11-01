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
)

// checks if the OfferValidations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferValidations{}

// OfferValidations Offer validation object
type OfferValidations struct {
	// Minimum Amount for Offer to be Applicable
	MinAmount *float32 `json:"min_amount,omitempty"`
	PaymentMethod OfferValidationsPaymentMethod `json:"payment_method"`
	// Maximum amount of Offer that can be availed.
	MaxAllowed float32 `json:"max_allowed"`
}

// NewOfferValidations instantiates a new OfferValidations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferValidations(paymentMethod OfferValidationsPaymentMethod, maxAllowed float32) *OfferValidations {
	this := OfferValidations{}
	this.PaymentMethod = paymentMethod
	this.MaxAllowed = maxAllowed
	return &this
}

// NewOfferValidationsWithDefaults instantiates a new OfferValidations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferValidationsWithDefaults() *OfferValidations {
	this := OfferValidations{}
	return &this
}

// GetMinAmount returns the MinAmount field value if set, zero value otherwise.
func (o *OfferValidations) GetMinAmount() float32 {
	if o == nil || IsNil(o.MinAmount) {
		var ret float32
		return ret
	}
	return *o.MinAmount
}

// GetMinAmountOk returns a tuple with the MinAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferValidations) GetMinAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.MinAmount) {
		return nil, false
	}
	return o.MinAmount, true
}

// HasMinAmount returns a boolean if a field has been set.
func (o *OfferValidations) HasMinAmount() bool {
	if o != nil && !IsNil(o.MinAmount) {
		return true
	}

	return false
}

// SetMinAmount gets a reference to the given float32 and assigns it to the MinAmount field.
func (o *OfferValidations) SetMinAmount(v float32) {
	o.MinAmount = &v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *OfferValidations) GetPaymentMethod() OfferValidationsPaymentMethod {
	if o == nil {
		var ret OfferValidationsPaymentMethod
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *OfferValidations) GetPaymentMethodOk() (*OfferValidationsPaymentMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *OfferValidations) SetPaymentMethod(v OfferValidationsPaymentMethod) {
	o.PaymentMethod = v
}

// GetMaxAllowed returns the MaxAllowed field value
func (o *OfferValidations) GetMaxAllowed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxAllowed
}

// GetMaxAllowedOk returns a tuple with the MaxAllowed field value
// and a boolean to check if the value has been set.
func (o *OfferValidations) GetMaxAllowedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAllowed, true
}

// SetMaxAllowed sets field value
func (o *OfferValidations) SetMaxAllowed(v float32) {
	o.MaxAllowed = v
}

func (o OfferValidations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferValidations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinAmount) {
		toSerialize["min_amount"] = o.MinAmount
	}
	toSerialize["payment_method"] = o.PaymentMethod
	toSerialize["max_allowed"] = o.MaxAllowed
	return toSerialize, nil
}

type NullableOfferValidations struct {
	value *OfferValidations
	isSet bool
}

func (v NullableOfferValidations) Get() *OfferValidations {
	return v.value
}

func (v *NullableOfferValidations) Set(val *OfferValidations) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferValidations) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferValidations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferValidations(val *OfferValidations) *NullableOfferValidations {
	return &NullableOfferValidations{value: val, isSet: true}
}

func (v NullableOfferValidations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferValidations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

