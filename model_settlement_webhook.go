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

// checks if the SettlementWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettlementWebhook{}

// SettlementWebhook Settlement webhook object
type SettlementWebhook struct {
	Data *SettlementWebhookDataEntity `json:"data,omitempty"`
	EventTime *string `json:"event_time,omitempty"`
	Type *string `json:"type,omitempty"`
}


func (o SettlementWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettlementWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.EventTime) {
		toSerialize["event_time"] = o.EventTime
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}




func cashfreeStringTest() {
	strings.HasPrefix("cf", "cf")
}