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

// checks if the EMIOffer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EMIOffer{}

// EMIOffer struct for EMIOffer
type EMIOffer struct {
	// Type of emi offer. Possible values are `credit_card_emi`, `debit_card_emi`, `cardless_emi`
	Type string `json:"type"`
	// Bank Name
	Issuer string `json:"issuer"`
	Tenures []int32 `json:"tenures"`
}

// NewEMIOffer instantiates a new EMIOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEMIOffer(type_ string, issuer string, tenures []int32) *EMIOffer {
	this := EMIOffer{}
	this.Type = type_
	this.Issuer = issuer
	this.Tenures = tenures
	return &this
}

// NewEMIOfferWithDefaults instantiates a new EMIOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEMIOfferWithDefaults() *EMIOffer {
	this := EMIOffer{}
	return &this
}

// GetType returns the Type field value
func (o *EMIOffer) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EMIOffer) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EMIOffer) SetType(v string) {
	o.Type = v
}

// GetIssuer returns the Issuer field value
func (o *EMIOffer) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *EMIOffer) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *EMIOffer) SetIssuer(v string) {
	o.Issuer = v
}

// GetTenures returns the Tenures field value
func (o *EMIOffer) GetTenures() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Tenures
}

// GetTenuresOk returns a tuple with the Tenures field value
// and a boolean to check if the value has been set.
func (o *EMIOffer) GetTenuresOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenures, true
}

// SetTenures sets field value
func (o *EMIOffer) SetTenures(v []int32) {
	o.Tenures = v
}

func (o EMIOffer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EMIOffer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["issuer"] = o.Issuer
	toSerialize["tenures"] = o.Tenures
	return toSerialize, nil
}

type NullableEMIOffer struct {
	value *EMIOffer
	isSet bool
}

func (v NullableEMIOffer) Get() *EMIOffer {
	return v.value
}

func (v *NullableEMIOffer) Set(val *EMIOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableEMIOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableEMIOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEMIOffer(val *EMIOffer) *NullableEMIOffer {
	return &NullableEMIOffer{value: val, isSet: true}
}

func (v NullableEMIOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEMIOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

