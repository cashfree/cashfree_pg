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

// checks if the VendorSplit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VendorSplit{}

// VendorSplit Use to split order when cashfree's Easy Split is enabled for your account.
type VendorSplit struct {
	// Vendor id created in Cashfree system
	VendorId *string `json:"vendor_id,omitempty"`
	// Amount which will be associated with this vendor
	Amount *float32 `json:"amount,omitempty"`
	// Percentage of order amount which shall get added to vendor account
	Percentage *float32 `json:"percentage,omitempty"`
}


func (o VendorSplit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VendorSplit) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}



