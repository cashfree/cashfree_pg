/*
Cashfree Payment Gateway APIs

Cashfree's Payment Gateway APIs provide developers with a streamlined pathway to integrate advanced payment processing capabilities into their applications, platforms and websites.

API version: 2023-08-01
Contact: developers@cashfree.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cashfree_pg

import (
	"encoding/json"
	"strings"
)

// checks if the CartItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartItem{}

// CartItem Each item in the cart.
type CartItem struct {
	// Unique identifier of the item
	ItemId *string `json:"item_id,omitempty"`
	// Name of the item
	ItemName *string `json:"item_name,omitempty"`
	// Description of the item
	ItemDescription *string `json:"item_description,omitempty"`
	// Tags attached to that item
	ItemTags []string `json:"item_tags,omitempty"`
	// Item details url
	ItemDetailsUrl *string `json:"item_details_url,omitempty"`
	// Item image url
	ItemImageUrl *string `json:"item_image_url,omitempty"`
	// Original price
	ItemOriginalUnitPrice *float64 `json:"item_original_unit_price,omitempty"`
	// Discounted Price
	ItemDiscountedUnitPrice *float64 `json:"item_discounted_unit_price,omitempty"`
	// Currency of the item.
	ItemCurrency *string `json:"item_currency,omitempty"`
	// Quantity if that item
	ItemQuantity *float32 `json:"item_quantity,omitempty"`
}


func (o CartItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartItem) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["item_id"] = o.ItemId
	}
	if !IsNil(o.ItemName) {
		toSerialize["item_name"] = o.ItemName
	}
	if !IsNil(o.ItemDescription) {
		toSerialize["item_description"] = o.ItemDescription
	}
	if !IsNil(o.ItemTags) {
		toSerialize["item_tags"] = o.ItemTags
	}
	if !IsNil(o.ItemDetailsUrl) {
		toSerialize["item_details_url"] = o.ItemDetailsUrl
	}
	if !IsNil(o.ItemImageUrl) {
		toSerialize["item_image_url"] = o.ItemImageUrl
	}
	if !IsNil(o.ItemOriginalUnitPrice) {
		toSerialize["item_original_unit_price"] = o.ItemOriginalUnitPrice
	}
	if !IsNil(o.ItemDiscountedUnitPrice) {
		toSerialize["item_discounted_unit_price"] = o.ItemDiscountedUnitPrice
	}
	if !IsNil(o.ItemCurrency) {
		toSerialize["item_currency"] = o.ItemCurrency
	}
	if !IsNil(o.ItemQuantity) {
		toSerialize["item_quantity"] = o.ItemQuantity
	}
	return toSerialize, nil
}


