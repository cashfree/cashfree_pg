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

// CFRefundRequest struct for CFRefundRequest
type CFRefundRequest struct {
	// Amount to be refunded. Should be lesser than or equal to the transaction amount. (Decimals allowed)
	RefundAmount float64 `json:"refund_amount"`
	// An unique ID to associate the refund with. Provie alphanumeric values
	RefundId string `json:"refund_id"`
	// A refund note for your reference.
	RefundNote *string `json:"refund_note,omitempty"`
	RefundSpeed *string `json:"refund_speed,omitempty"`
	RefundSplits []CFVendorSplit `json:"refund_splits,omitempty"`
}

// NewCFRefundRequest instantiates a new CFRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFRefundRequest(refundAmount float64, refundId string) *CFRefundRequest {
	this := CFRefundRequest{}
	this.RefundAmount = refundAmount
	this.RefundId = refundId
	return &this
}

// NewCFRefundRequestWithDefaults instantiates a new CFRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFRefundRequestWithDefaults() *CFRefundRequest {
	this := CFRefundRequest{}
	return &this
}

// GetRefundAmount returns the RefundAmount field value
func (o *CFRefundRequest) GetRefundAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value
// and a boolean to check if the value has been set.
func (o *CFRefundRequest) GetRefundAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RefundAmount, true
}

// SetRefundAmount sets field value
func (o *CFRefundRequest) SetRefundAmount(v float64) {
	o.RefundAmount = v
}

// GetRefundId returns the RefundId field value
func (o *CFRefundRequest) GetRefundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value
// and a boolean to check if the value has been set.
func (o *CFRefundRequest) GetRefundIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RefundId, true
}

// SetRefundId sets field value
func (o *CFRefundRequest) SetRefundId(v string) {
	o.RefundId = v
}

// GetRefundNote returns the RefundNote field value if set, zero value otherwise.
func (o *CFRefundRequest) GetRefundNote() string {
	if o == nil || o.RefundNote == nil {
		var ret string
		return ret
	}
	return *o.RefundNote
}

// GetRefundNoteOk returns a tuple with the RefundNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFRefundRequest) GetRefundNoteOk() (*string, bool) {
	if o == nil || o.RefundNote == nil {
		return nil, false
	}
	return o.RefundNote, true
}

// HasRefundNote returns a boolean if a field has been set.
func (o *CFRefundRequest) HasRefundNote() bool {
	if o != nil && o.RefundNote != nil {
		return true
	}

	return false
}

// SetRefundNote gets a reference to the given string and assigns it to the RefundNote field.
func (o *CFRefundRequest) SetRefundNote(v string) {
	o.RefundNote = &v
}

// GetRefundSplits returns the RefundSplits field value if set, zero value otherwise.
func (o *CFRefundRequest) GetRefundSplits() []CFVendorSplit {
	if o == nil || o.RefundSplits == nil {
		var ret []CFVendorSplit
		return ret
	}
	return o.RefundSplits
}

// GetRefundSplitsOk returns a tuple with the RefundSplits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFRefundRequest) GetRefundSplitsOk() ([]CFVendorSplit, bool) {
	if o == nil || o.RefundSplits == nil {
		return nil, false
	}
	return o.RefundSplits, true
}

// HasRefundSplits returns a boolean if a field has been set.
func (o *CFRefundRequest) HasRefundSplits() bool {
	if o != nil && o.RefundSplits != nil {
		return true
	}

	return false
}

// SetRefundSplits gets a reference to the given []CFVendorSplit and assigns it to the RefundSplits field.
func (o *CFRefundRequest) SetRefundSplits(v []CFVendorSplit) {
	o.RefundSplits = v
}

func (o CFRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["refund_amount"] = o.RefundAmount
	}
	if true {
		toSerialize["refund_id"] = o.RefundId
	}
	if o.RefundNote != nil {
		toSerialize["refund_note"] = o.RefundNote
	}
	if o.RefundSplits != nil {
		toSerialize["refund_splits"] = o.RefundSplits
	}
	return json.Marshal(toSerialize)
}

type NullableCFRefundRequest struct {
	value *CFRefundRequest
	isSet bool
}

func (v NullableCFRefundRequest) Get() *CFRefundRequest {
	return v.value
}

func (v *NullableCFRefundRequest) Set(val *CFRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCFRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCFRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFRefundRequest(val *CFRefundRequest) *NullableCFRefundRequest {
	return &NullableCFRefundRequest{value: val, isSet: true}
}

func (v NullableCFRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


