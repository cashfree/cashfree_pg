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

// checks if the CreateSubscriptionRequestAuthorizationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionRequestAuthorizationDetails{}

// CreateSubscriptionRequestAuthorizationDetails struct for CreateSubscriptionRequestAuthorizationDetails
type CreateSubscriptionRequestAuthorizationDetails struct {
	// Authorization amount for the auth payment.
	AuthorizationAmount *float32 `json:"authorization_amount,omitempty"`
	// Indicates whether the authorization amount should be refunded to the customer automatically. Merchants can use this field to specify if the authorized funds should be returned to the customer after authorization of the subscription.
	AuthorizationAmountRefund *bool `json:"authorization_amount_refund,omitempty"`
}


func (o CreateSubscriptionRequestAuthorizationDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionRequestAuthorizationDetails) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationAmount) {
		toSerialize["authorization_amount"] = o.AuthorizationAmount
	}
	if !IsNil(o.AuthorizationAmountRefund) {
		toSerialize["authorization_amount_refund"] = o.AuthorizationAmountRefund
	}
	return toSerialize, nil
}


