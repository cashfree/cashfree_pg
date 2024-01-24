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
	"fmt"
)

// PayOrderRequestPaymentMethod - struct for PayOrderRequestPaymentMethod
type PayOrderRequestPaymentMethod struct {
	AppPaymentMethod *AppPaymentMethod
	CardEMIPaymentMethod *CardEMIPaymentMethod
	CardPaymentMethod *CardPaymentMethod
	CardlessEMIPaymentMethod *CardlessEMIPaymentMethod
	NetBankingPaymentMethod *NetBankingPaymentMethod
	PaylaterPaymentMethod *PaylaterPaymentMethod
	UPIPaymentMethod *UPIPaymentMethod
}

// AppPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns AppPaymentMethod wrapped in PayOrderRequestPaymentMethod
func AppPaymentMethodAsPayOrderRequestPaymentMethod(v *AppPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		AppPaymentMethod: v,
	}
}

// CardEMIPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns CardEMIPaymentMethod wrapped in PayOrderRequestPaymentMethod
func CardEMIPaymentMethodAsPayOrderRequestPaymentMethod(v *CardEMIPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		CardEMIPaymentMethod: v,
	}
}

// CardPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns CardPaymentMethod wrapped in PayOrderRequestPaymentMethod
func CardPaymentMethodAsPayOrderRequestPaymentMethod(v *CardPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		CardPaymentMethod: v,
	}
}

// CardlessEMIPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns CardlessEMIPaymentMethod wrapped in PayOrderRequestPaymentMethod
func CardlessEMIPaymentMethodAsPayOrderRequestPaymentMethod(v *CardlessEMIPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		CardlessEMIPaymentMethod: v,
	}
}

// NetBankingPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns NetBankingPaymentMethod wrapped in PayOrderRequestPaymentMethod
func NetBankingPaymentMethodAsPayOrderRequestPaymentMethod(v *NetBankingPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		NetBankingPaymentMethod: v,
	}
}

// PaylaterPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns PaylaterPaymentMethod wrapped in PayOrderRequestPaymentMethod
func PaylaterPaymentMethodAsPayOrderRequestPaymentMethod(v *PaylaterPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		PaylaterPaymentMethod: v,
	}
}

// UPIPaymentMethodAsPayOrderRequestPaymentMethod is a convenience function that returns UPIPaymentMethod wrapped in PayOrderRequestPaymentMethod
func UPIPaymentMethodAsPayOrderRequestPaymentMethod(v *UPIPaymentMethod) PayOrderRequestPaymentMethod {
	return PayOrderRequestPaymentMethod{
		UPIPaymentMethod: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PayOrderRequestPaymentMethod) UnmarshalJSON(data []byte) error {
	fmt.Println("data")
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PayOrderRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	if src.AppPaymentMethod != nil {
		return json.Marshal(&src.AppPaymentMethod)
	}

	if src.CardEMIPaymentMethod != nil {
		return json.Marshal(&src.CardEMIPaymentMethod)
	}

	if src.CardPaymentMethod != nil {
		return json.Marshal(&src.CardPaymentMethod)
	}

	if src.CardlessEMIPaymentMethod != nil {
		return json.Marshal(&src.CardlessEMIPaymentMethod)
	}

	if src.NetBankingPaymentMethod != nil {
		return json.Marshal(&src.NetBankingPaymentMethod)
	}

	if src.PaylaterPaymentMethod != nil {
		return json.Marshal(&src.PaylaterPaymentMethod)
	}

	if src.UPIPaymentMethod != nil {
		return json.Marshal(&src.UPIPaymentMethod)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PayOrderRequestPaymentMethod) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppPaymentMethod != nil {
		return obj.AppPaymentMethod
	}

	if obj.CardEMIPaymentMethod != nil {
		return obj.CardEMIPaymentMethod
	}

	if obj.CardPaymentMethod != nil {
		return obj.CardPaymentMethod
	}

	if obj.CardlessEMIPaymentMethod != nil {
		return obj.CardlessEMIPaymentMethod
	}

	if obj.NetBankingPaymentMethod != nil {
		return obj.NetBankingPaymentMethod
	}

	if obj.PaylaterPaymentMethod != nil {
		return obj.PaylaterPaymentMethod
	}

	if obj.UPIPaymentMethod != nil {
		return obj.UPIPaymentMethod
	}

	// all schemas are nil
	return nil
}



