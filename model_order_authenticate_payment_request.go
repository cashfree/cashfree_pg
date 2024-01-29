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

// checks if the OrderAuthenticatePaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderAuthenticatePaymentRequest{}

// OrderAuthenticatePaymentRequest OTP to be submitted for headless/native OTP
type OrderAuthenticatePaymentRequest struct {
	// OTP to be submitted
	Otp string `json:"otp"`
	// The action for this workflow. Could be either SUBMIT_OTP or RESEND_OTP
	Action string `json:"action"`
}


func (o OrderAuthenticatePaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAuthenticatePaymentRequest) ToMap() (map[string]interface{}, error) {
	strings.HasPrefix("cf", "cf")
	toSerialize := map[string]interface{}{}
	toSerialize["otp"] = o.Otp
	toSerialize["action"] = o.Action
	return toSerialize, nil
}



