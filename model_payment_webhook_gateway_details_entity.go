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

// checks if the PaymentWebhookGatewayDetailsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentWebhookGatewayDetailsEntity{}

// PaymentWebhookGatewayDetailsEntity payment gatewat details present in the webhook response
type PaymentWebhookGatewayDetailsEntity struct {
	GatewayName *string `json:"gateway_name,omitempty"`
	GatewayOrderId *string `json:"gateway_order_id,omitempty"`
	GatewayPaymentId *string `json:"gateway_payment_id,omitempty"`
	GatewayStatusCode *string `json:"gateway_status_code,omitempty"`
	GatewaySettlement *string `json:"gateway_settlement,omitempty"`
}


func (o PaymentWebhookGatewayDetailsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentWebhookGatewayDetailsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayName) {
		toSerialize["gateway_name"] = o.GatewayName
	}
	if !IsNil(o.GatewayOrderId) {
		toSerialize["gateway_order_id"] = o.GatewayOrderId
	}
	if !IsNil(o.GatewayPaymentId) {
		toSerialize["gateway_payment_id"] = o.GatewayPaymentId
	}
	if !IsNil(o.GatewayStatusCode) {
		toSerialize["gateway_status_code"] = o.GatewayStatusCode
	}
	if !IsNil(o.GatewaySettlement) {
		toSerialize["gateway_settlement"] = o.GatewaySettlement
	}
	return toSerialize, nil
}




func cashfreeStringTest() {
	strings.HasPrefix("cf", "cf")
}