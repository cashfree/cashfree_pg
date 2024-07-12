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

// checks if the CreateSubscriptionRequestSubscriptionMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionRequestSubscriptionMeta{}

// CreateSubscriptionRequestSubscriptionMeta struct for CreateSubscriptionRequestSubscriptionMeta
type CreateSubscriptionRequestSubscriptionMeta struct {
	// The url to redirect after checkout.
	ReturnUrl *string `json:"return_url,omitempty"`
	// Notification channel for the subscription. SMS, EMAIL are possible values.
	NotificationChannel []string `json:"notification_channel,omitempty"`
}


func (o CreateSubscriptionRequestSubscriptionMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionRequestSubscriptionMeta) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnUrl) {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if !IsNil(o.NotificationChannel) {
		toSerialize["notification_channel"] = o.NotificationChannel
	}
	return toSerialize, nil
}


