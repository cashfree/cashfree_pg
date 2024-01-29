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

// checks if the OfferDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferDetails{}

// OfferDetails Offer details and type
type OfferDetails struct {
	// Offer Type for the Offer.
	OfferType string `json:"offer_type"`
	DiscountDetails *DiscountDetails `json:"discount_details,omitempty"`
	CashbackDetails *CashbackDetails `json:"cashback_details,omitempty"`
}


func (o OfferDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["offer_type"] = o.OfferType
	if !IsNil(o.DiscountDetails) {
		toSerialize["discount_details"] = o.DiscountDetails
	}
	if !IsNil(o.CashbackDetails) {
		toSerialize["cashback_details"] = o.CashbackDetails
	}
	return toSerialize, nil
}



