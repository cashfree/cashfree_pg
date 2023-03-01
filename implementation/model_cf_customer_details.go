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

// CFCustomerDetails struct for CFCustomerDetails
type CFCustomerDetails struct {
	// A unique identifier for the customer. Use alphanumeric values only.
	CustomerId string `json:"customer_id"`
	// Customer email address.
	CustomerEmail string `json:"customer_email"`
	// Customer phone number.
	CustomerPhone string `json:"customer_phone"`
	// Customer bank account. Required if you want to do a bank account check (TPV)
	CustomerBankAccountNumber string `json:"customer_bank_account_number"`
	// Customer bank IFSC. Required if you want to do a bank account check (TPV)
	CustomerBankIfsc string `json:"customer_bank_ifsc"`
	// Customer bank code. Required for net banking payments, if you want to do a bank account check (TPV)
	CustomerBankCode int32 `json:"customer_bank_code"`
}

// NewCFCustomerDetails instantiates a new CFCustomerDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFCustomerDetails(customerId string, customerEmail string, customerPhone string, customerBankAccountNumber string, customerBankIfsc string, customerBankCode int32) *CFCustomerDetails {
	this := CFCustomerDetails{}
	this.CustomerId = customerId
	this.CustomerEmail = customerEmail
	this.CustomerPhone = customerPhone
	this.CustomerBankAccountNumber = customerBankAccountNumber
	this.CustomerBankIfsc = customerBankIfsc
	this.CustomerBankCode = customerBankCode
	return &this
}

// NewCFCustomerDetailsWithDefaults instantiates a new CFCustomerDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFCustomerDetailsWithDefaults() *CFCustomerDetails {
	this := CFCustomerDetails{}
	return &this
}

// GetCustomerId returns the CustomerId field value
func (o *CFCustomerDetails) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CFCustomerDetails) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CFCustomerDetails) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetCustomerEmail returns the CustomerEmail field value
func (o *CFCustomerDetails) GetCustomerEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerEmail
}

// GetCustomerEmailOk returns a tuple with the CustomerEmail field value
// and a boolean to check if the value has been set.
func (o *CFCustomerDetails) GetCustomerEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerEmail, true
}

// SetCustomerEmail sets field value
func (o *CFCustomerDetails) SetCustomerEmail(v string) {
	o.CustomerEmail = v
}

// GetCustomerPhone returns the CustomerPhone field value
func (o *CFCustomerDetails) GetCustomerPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerPhone
}

// GetCustomerPhoneOk returns a tuple with the CustomerPhone field value
// and a boolean to check if the value has been set.
func (o *CFCustomerDetails) GetCustomerPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerPhone, true
}

// SetCustomerPhone sets field value
func (o *CFCustomerDetails) SetCustomerPhone(v string) {
	o.CustomerPhone = v
}

// GetCustomerBankAccountNumber returns the CustomerBankAccountNumber field value
func (o *CFCustomerDetails) GetCustomerBankAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerBankAccountNumber
}

// GetCustomerBankAccountNumberOk returns a tuple with the CustomerBankAccountNumber field value
// and a boolean to check if the value has been set.
func (o *CFCustomerDetails) GetCustomerBankAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerBankAccountNumber, true
}

// SetCustomerBankAccountNumber sets field value
func (o *CFCustomerDetails) SetCustomerBankAccountNumber(v string) {
	o.CustomerBankAccountNumber = v
}

// GetCustomerBankIfsc returns the CustomerBankIfsc field value
func (o *CFCustomerDetails) GetCustomerBankIfsc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerBankIfsc
}

// GetCustomerBankIfscOk returns a tuple with the CustomerBankIfsc field value
// and a boolean to check if the value has been set.
func (o *CFCustomerDetails) GetCustomerBankIfscOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerBankIfsc, true
}

// SetCustomerBankIfsc sets field value
func (o *CFCustomerDetails) SetCustomerBankIfsc(v string) {
	o.CustomerBankIfsc = v
}

// GetCustomerBankCode returns the CustomerBankCode field value
func (o *CFCustomerDetails) GetCustomerBankCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CustomerBankCode
}

// GetCustomerBankCodeOk returns a tuple with the CustomerBankCode field value
// and a boolean to check if the value has been set.
func (o *CFCustomerDetails) GetCustomerBankCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerBankCode, true
}

// SetCustomerBankCode sets field value
func (o *CFCustomerDetails) SetCustomerBankCode(v int32) {
	o.CustomerBankCode = v
}

func (o CFCustomerDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["customer_email"] = o.CustomerEmail
	}
	if true {
		toSerialize["customer_phone"] = o.CustomerPhone
	}
	if true {
		toSerialize["customer_bank_account_number"] = o.CustomerBankAccountNumber
	}
	if true {
		toSerialize["customer_bank_ifsc"] = o.CustomerBankIfsc
	}
	if true {
		toSerialize["customer_bank_code"] = o.CustomerBankCode
	}
	return json.Marshal(toSerialize)
}

type NullableCFCustomerDetails struct {
	value *CFCustomerDetails
	isSet bool
}

func (v NullableCFCustomerDetails) Get() *CFCustomerDetails {
	return v.value
}

func (v *NullableCFCustomerDetails) Set(val *CFCustomerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCFCustomerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCFCustomerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFCustomerDetails(val *CFCustomerDetails) *NullableCFCustomerDetails {
	return &NullableCFCustomerDetails{value: val, isSet: true}
}

func (v NullableCFCustomerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFCustomerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


