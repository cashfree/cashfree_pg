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

// checks if the LinkMetaEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkMetaEntity{}

// LinkMetaEntity Payment link meta information object
type LinkMetaEntity struct {
	// Notification URL for server-server communication. It should be an https URL.
	NotifyUrl *string `json:"notify_url,omitempty"`
	// If \"true\", link will directly open UPI Intent flow on mobile, and normal link flow elsewhere
	UpiIntent *string `json:"upi_intent,omitempty"`
	// The URL to which user will be redirected to after the payment is done on the link. Maximum length: 250.
	ReturnUrl *string `json:"return_url,omitempty"`
	// Allowed payment modes for this link. Pass comma-separated values among following options - \"cc\", \"dc\", \"ccc\", \"ppc\", \"nb\", \"upi\", \"paypal\", \"app\". Leave it blank to show all available payment methods
	PaymentMethods *string `json:"payment_methods,omitempty"`
}

// NewLinkMetaEntity instantiates a new LinkMetaEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkMetaEntity() *LinkMetaEntity {
	this := LinkMetaEntity{}
	return &this
}

// NewLinkMetaEntityWithDefaults instantiates a new LinkMetaEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkMetaEntityWithDefaults() *LinkMetaEntity {
	this := LinkMetaEntity{}
	return &this
}

// GetNotifyUrl returns the NotifyUrl field value if set, zero value otherwise.
func (o *LinkMetaEntity) GetNotifyUrl() string {
	if o == nil || IsNil(o.NotifyUrl) {
		var ret string
		return ret
	}
	return *o.NotifyUrl
}

// GetNotifyUrlOk returns a tuple with the NotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkMetaEntity) GetNotifyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyUrl) {
		return nil, false
	}
	return o.NotifyUrl, true
}

// HasNotifyUrl returns a boolean if a field has been set.
func (o *LinkMetaEntity) HasNotifyUrl() bool {
	if o != nil && !IsNil(o.NotifyUrl) {
		return true
	}

	return false
}

// SetNotifyUrl gets a reference to the given string and assigns it to the NotifyUrl field.
func (o *LinkMetaEntity) SetNotifyUrl(v string) {
	o.NotifyUrl = &v
}

// GetUpiIntent returns the UpiIntent field value if set, zero value otherwise.
func (o *LinkMetaEntity) GetUpiIntent() string {
	if o == nil || IsNil(o.UpiIntent) {
		var ret string
		return ret
	}
	return *o.UpiIntent
}

// GetUpiIntentOk returns a tuple with the UpiIntent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkMetaEntity) GetUpiIntentOk() (*string, bool) {
	if o == nil || IsNil(o.UpiIntent) {
		return nil, false
	}
	return o.UpiIntent, true
}

// HasUpiIntent returns a boolean if a field has been set.
func (o *LinkMetaEntity) HasUpiIntent() bool {
	if o != nil && !IsNil(o.UpiIntent) {
		return true
	}

	return false
}

// SetUpiIntent gets a reference to the given string and assigns it to the UpiIntent field.
func (o *LinkMetaEntity) SetUpiIntent(v string) {
	o.UpiIntent = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *LinkMetaEntity) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkMetaEntity) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *LinkMetaEntity) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *LinkMetaEntity) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *LinkMetaEntity) GetPaymentMethods() string {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret string
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkMetaEntity) GetPaymentMethodsOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *LinkMetaEntity) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given string and assigns it to the PaymentMethods field.
func (o *LinkMetaEntity) SetPaymentMethods(v string) {
	o.PaymentMethods = &v
}

func (o LinkMetaEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkMetaEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotifyUrl) {
		toSerialize["notify_url"] = o.NotifyUrl
	}
	if !IsNil(o.UpiIntent) {
		toSerialize["upi_intent"] = o.UpiIntent
	}
	if !IsNil(o.ReturnUrl) {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	return toSerialize, nil
}

type NullableLinkMetaEntity struct {
	value *LinkMetaEntity
	isSet bool
}

func (v NullableLinkMetaEntity) Get() *LinkMetaEntity {
	return v.value
}

func (v *NullableLinkMetaEntity) Set(val *LinkMetaEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkMetaEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkMetaEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkMetaEntity(val *LinkMetaEntity) *NullableLinkMetaEntity {
	return &NullableLinkMetaEntity{value: val, isSet: true}
}

func (v NullableLinkMetaEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkMetaEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


